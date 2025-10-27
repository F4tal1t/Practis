import os
from PIL import Image


def resize_images(input_folder, output_folder, target_width, target_height):
    """
    Resize all PNG images in the input folder to the specified dimensions
    and save them to the output folder.

    Args:
        input_folder (str): Path to folder containing PNG images
        output_folder (str): Path to folder where resized images will be saved
        target_width (int): Desired width in pixels
        target_height (int): Desired height in pixels
    """
    # Create output folder if it doesn't exist
    if not os.path.exists(output_folder):
        os.makedirs(output_folder)
        print(f"Created output directory: {output_folder}")

    # Get all PNG files in the input folder
    png_files = [f for f in os.listdir(input_folder) if f.lower().endswith('.png')]

    if not png_files:
        print("No PNG files found in the input folder.")
        return

    print(f"Found {len(png_files)} PNG files to resize.")

    # Process each PNG file
    for file_name in png_files:
        input_path = os.path.join(input_folder, file_name)
        output_path = os.path.join(output_folder, file_name)

        try:
            # Open the image
            img = Image.open(input_path)

            # Resize the image
            resized_img = img.resize((target_width, target_height), Image.Resampling.LANCZOS)

            # Save the resized image
            resized_img.save(output_path)

            print(f"Resized {file_name} to {target_width}x{target_height}")

        except Exception as e:
            print(f"Error processing {file_name}: {e}")


if __name__ == "__main__":
    # Example usage
    input_folder = input("Enter the input folder path: ")
    output_folder = input("Enter the output folder path (leave blank to create 'resized' subfolder): ")

    if not output_folder:
        output_folder = os.path.join(input_folder, "resized")

    try:
        target_width = int(input("Enter the target width in pixels:"))
        target_height = int(input("Enter the target height in pixels: "))

        resize_images(input_folder, output_folder, target_width, target_height)
        print("Resizing complete!")

    except ValueError:
        print("Width and height must be integers.")