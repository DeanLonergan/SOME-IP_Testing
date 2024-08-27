package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

// Global variables to store command references and their output buffers
var (
	clientCmd, serviceCmd, routingManagerCmd          *exec.Cmd
	clientOutput, serviceOutput, routingManagerOutput bytes.Buffer
)

// runCommand starts a command with the given arguments and captures its output in the provided buffer.
func runCommand(command string, output *bytes.Buffer, args ...string) (*exec.Cmd, error) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = output
	cmd.Stderr = output
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start %s: %w", command, err)
	}
	return cmd, nil
}

// theRoutingManagerIsRunning ensures that the routing manager is running.
func theRoutingManagerIsRunning() error {
	fmt.Println("Starting routing manager in the background...")
	cmd, err := runCommand("../build/routingmanagerd/routingmanagerd", &routingManagerOutput)
	if err != nil {
		return err
	}
	routingManagerCmd = cmd
	time.Sleep(2 * time.Second) // Allow time for initialization

	if !bytes.Contains(routingManagerOutput.Bytes(), []byte("Instantiating routing manager [Host]")) {
		return fmt.Errorf("routing manager failed to start as Host: %s", routingManagerOutput.String())
	}
	return nil
}

// theServiceIsRunning ensures that the service is running.
func theServiceIsRunning() error {
	fmt.Println("Starting service in the background...")
	cmd, err := runCommand("../build/service/service", &serviceOutput)
	if err != nil {
		return err
	}
	serviceCmd = cmd
	time.Sleep(2 * time.Second) // Allow time for initialization
	return nil
}

// theClientApplicationStarts starts the client application and captures its output.
func theClientApplicationStarts() error {
	fmt.Println("Starting client...")
	cmd, err := runCommand("../build/client/client", &clientOutput)
	if err != nil {
		return err
	}
	clientCmd = cmd
	time.Sleep(2 * time.Second) // Allow time for initialization
	return nil
}

// theClientShouldSuccessfullySubscribeToTheService checks that the client subscribed successfully.
func theClientShouldSuccessfullySubscribeToTheService() error {
	if !bytes.Contains(clientOutput.Bytes(), []byte("SUBSCRIBE ACK")) {
		return fmt.Errorf("subscription not successful: %s", clientOutput.String())
	}
	return nil
}

// theClientShouldReceiveAConfirmationMessage checks that the client received the expected message.
func theClientShouldReceiveAConfirmationMessage(expectedMessage string) error {
	if !bytes.Contains(clientOutput.Bytes(), []byte(expectedMessage)) {
		return fmt.Errorf("expected output \"%s\" not found in client output", expectedMessage)
	}
	return nil
}

// teardown stops all running processes.
func teardown() {
	fmt.Println("Running teardown...")

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
}

// InitializeScenario registers the step definitions with the Godog framework.
func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the routingmanagerd is running$`, theRoutingManagerIsRunning)
	ctx.Step(`^the service is running$`, theServiceIsRunning)
	ctx.Step(`^the client application starts$`, theClientApplicationStarts)
	ctx.Step(`^the client should successfully subscribe to the service$`, theClientShouldSuccessfullySubscribeToTheService)
	ctx.Step(`^the client should receive a confirmation message "([^"]*)"$`, theClientShouldReceiveAConfirmationMessage)
}

// TestFeatures runs the Godog test suite using Go's testing framework.
func TestFeatures(t *testing.T) {
	// Print the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	fmt.Println("Current working directory:", cwd)

	// Set Godog options
	opts := godog.Options{
		Format: "pretty",
		Paths:  []string{"./client.feature"},
	}

	// Run the Godog test suite
	status := godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	// Call teardown after tests
	teardown()

	// Fail the test if the test suite did not complete successfully
	if status != 0 {
		t.Fatalf("non-zero status returned, failed to run feature tests")
	}
}
