import os
import shutil
import subprocess

def clean_build_directory(build_dir):
    """Remove the build directory if it exists."""
    if os.path.exists(build_dir):
        print(f"Cleaning the build directory: {build_dir}")
        shutil.rmtree(build_dir)
    else:
        print(f"Build directory does not exist, no need to clean: {build_dir}")

def run_cmake_and_make(build_dir):
    """Run cmake and make to build the project."""
    try:
        os.makedirs(build_dir, exist_ok=True)
        print("Running cmake...")
        subprocess.check_call(["cmake", ".."], cwd=build_dir)
        print("Running make...")
        subprocess.check_call(["make"], cwd=build_dir)
        print("Build completed successfully.")
    except subprocess.CalledProcessError as e:
        print(f"An error occurred during the build process: {e}")
    except Exception as e:
        print(f"An unexpected error occurred: {e}")

def main():
    script_dir = os.path.dirname(os.path.abspath(__file__))
    project_dir = script_dir  # Use the script's directory as the project root
    build_dir = os.path.join(project_dir, "build")

    print("Starting build process...")
    clean_build_directory(build_dir)
    run_cmake_and_make(build_dir)
    print("Build process finished.")

if __name__ == "__main__":
    main()
