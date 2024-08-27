package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

var clientCmd, serviceCmd, routingManagerCmd *exec.Cmd
var clientStarted, serviceRunning, routingManagerRunning bool
var subscriptionSuccessful bool
var expectedOutput string
var clientOutputBuffer, serviceOutputBuffer, routingManagerOutputBuffer bytes.Buffer

// Function to run a command in the background and capture its output
func runCommandWithOutputCapture(command string, outputBuffer *bytes.Buffer, args ...string) (*exec.Cmd, error) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = outputBuffer
	cmd.Stderr = outputBuffer
	err := cmd.Start()
	return cmd, err
}

// Step definition to ensure the routingmanagerd is running in the background
func theRoutingManagerIsRunning() error {
	fmt.Println("Starting routing manager in the background...")
	var err error
	routingManagerCmd, err = runCommandWithOutputCapture("../build/routingmanagerd/routingmanagerd", &routingManagerOutputBuffer)
	if err != nil {
		return fmt.Errorf("could not start routing manager: %v", err)
	}
	routingManagerRunning = true

	// Give the routing manager a moment to initialize
	time.Sleep(2 * time.Second)

	return nil
}

// Step definition to ensure the service is running in the background
func theServiceIsRunning() error {
	if !routingManagerRunning {
		return fmt.Errorf("routing manager is not running")
	}

	fmt.Println("Starting service in the background...")
	var err error
	serviceCmd, err = runCommandWithOutputCapture("../build/service/service", &serviceOutputBuffer)
	if err != nil {
		return fmt.Errorf("could not start service: %v", err)
	}
	serviceRunning = true

	// Give the service a moment to initialize
	time.Sleep(2 * time.Second)

	return nil
}

// Step definition to start the client application and capture its output
func theClientApplicationStarts() error {
	if !serviceRunning {
		return fmt.Errorf("service is not running")
	}

	fmt.Println("Waiting for routing manager and service to fully start...")
	time.Sleep(5 * time.Second)

	fmt.Println("Starting client...")
	var err error
	clientCmd, err = runCommandWithOutputCapture("../build/client/client", &clientOutputBuffer)
	if err != nil {
		return fmt.Errorf("could not start client: %v", err)
	}
	clientStarted = true

	return nil
}

// Step definition to verify the client successfully subscribed
func theClientShouldSuccessfullySubscribeToTheService() error {
	if !clientStarted {
		return fmt.Errorf("client has not started")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	done := make(chan error)
	go func() {
		for {
			if strings.Contains(clientOutputBuffer.String(), expectedOutput) {
				subscriptionSuccessful = true
				done <- nil
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("timeout reached: client did not successfully subscribe")
	case err := <-done:
		if err != nil {
			return err
		}
	}

	return nil
}

// Step definition to set the expected confirmation message
func theClientShouldReceiveAConfirmationMessage(message string) error {
	expectedOutput = strings.TrimSpace(message)
	actualOutput := strings.TrimSpace(clientOutputBuffer.String())

	// Print the actual outputs for debugging
	fmt.Println("Routing Manager Output:\n", routingManagerOutputBuffer.String())
	fmt.Println("Service Output:\n", serviceOutputBuffer.String())
	fmt.Println("Client Output:\n", actualOutput)

	// Compare the expected message with the actual output
	if !strings.Contains(actualOutput, expectedOutput) {
		return fmt.Errorf("expected output %q not found in client output", expectedOutput)
	}

	return nil
}

// Teardown function to stop the client, service, and routing manager
func teardown() {
	if clientCmd != nil {
		fmt.Println("Stopping client...")
		clientCmd.Process.Kill()
		clientCmd.Wait()
		fmt.Println("Client stopped.")
	}

	if serviceCmd != nil {
		fmt.Println("Stopping service...")
		serviceCmd.Process.Kill()
		serviceCmd.Wait()
		fmt.Println("Service stopped.")
	}

	if routingManagerCmd != nil {
		fmt.Println("Stopping routing manager...")
		routingManagerCmd.Process.Kill()
		routingManagerCmd.Wait()
		fmt.Println("Routing manager stopped.")
	}

	// Manually clean up Unix domain sockets
	fmt.Println("Cleaning up Unix domain sockets...")
	os.Remove("/tmp/vsomeip-0")
	os.Remove("/tmp/vsomeip-*")
}

// InitializeScenario registers the step definitions with the Godog framework
func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the routingmanagerd is running$`, theRoutingManagerIsRunning)
	ctx.Step(`^the service is running$`, theServiceIsRunning)
	ctx.Step(`^the client application starts$`, theClientApplicationStarts)
	ctx.Step(`^the client should successfully subscribe to the service$`, theClientShouldSuccessfullySubscribeToTheService)
	ctx.Step(`^the client should receive a confirmation message "([^"]*)"$`, theClientShouldReceiveAConfirmationMessage)
}

// Test function that runs the Godog test suite using Go's testing framework
func TestFeatures(t *testing.T) {
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)

	opts := godog.Options{
		Format: "pretty",
		Paths:  []string{"./client.feature"}, // Path relative to the test directory
	}

	status := godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	teardown() // Manually call teardown to ensure all processes are stopped

	if status != 0 {
		t.Fatalf("non-zero status returned, failed to run feature tests")
	}
}
