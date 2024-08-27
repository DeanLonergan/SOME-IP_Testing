import os
import subprocess

# Initializes the Go module if go.mod doesn't exist.
def setup_go_module():
    if not os.path.exists('go.mod'):
        print("Initializing Go module...")
        subprocess.run(['go', 'mod', 'init', 'someip_testing'], check=True)
    else:
        print("Go module already initialized (go.mod found).")

# Installs the Godog testing framework using go get (this will update go.sum and ensure all dependencies are present).
def install_godog():
    subprocess.run(['go', 'get', 'github.com/cucumber/godog/cmd/godog@latest'], check=True)

# Runs the build.py script to compile the C++ project.
def run_build_script():
    subprocess.run(['python3', 'build.py'], check=True)

# Changes directory to the test folder and runs the Go tests using go test.
def run_go_tests():
    os.chdir('test')
    subprocess.run(['go', 'test'], check=True)

# Main function that orchestrates the setup and execution of tests.
def main():
    try:
        setup_go_module()
        install_godog()
        run_build_script()
        run_go_tests()
    except subprocess.CalledProcessError as e:
        print(f"An error occurred: {e}")
        exit(1)

if __name__ == '__main__':
    main()
