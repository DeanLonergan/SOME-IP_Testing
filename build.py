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

def run_cmake_and_make(build_dir, source_dir):
    """Run cmake and make to build the project."""
    try:
        os.makedirs(build_dir, exist_ok=True)
        print(f"Running cmake for {source_dir}...")
        subprocess.check_call(["cmake", source_dir], cwd=build_dir)
        print("Running make...")
        subprocess.check_call(["make"], cwd=build_dir)
        print(f"Build completed successfully for {source_dir}.")
    except subprocess.CalledProcessError as e:
        print(f"An error occurred during the build process: {e}")
    except Exception as e:
        print(f"An unexpected error occurred: {e}")

def get_components(src_dir):
    """Retrieve a list of all component directories in the src directory."""
    components = []
    for item in os.listdir(src_dir):
        item_path = os.path.join(src_dir, item)
        if os.path.isdir(item_path) and os.path.exists(os.path.join(item_path, "CMakeLists.txt")):
            components.append(item)
    return components

def main():
    script_dir = os.path.dirname(os.path.abspath(__file__))
    build_root_dir = os.path.join(script_dir, "build")
    src_dir = os.path.join(script_dir, "src")
   
    print("Starting build process...")

    # Clean the main build directory
    clean_build_directory(build_root_dir)

    # Get the list of components dynamically
    components = get_components(src_dir)

    for component in components:
        source_dir = os.path.join(src_dir, component)
        component_build_dir = os.path.join(build_root_dir, component)

        print(f"\nBuilding component: {component}")
        run_cmake_and_make(component_build_dir, source_dir)

    print("Build process finished.")

if __name__ == "__main__":
    main()
