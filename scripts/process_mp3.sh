#!/bin/bash

# --- Configuration ---
# An array of input directories to search for MP3 files.
# The script will search these directories and all of their subdirectories.
INPUT_DIRS=("")

# The directory where the new, sped-up files will be saved.
OUTPUT_DIR=""

# The speed factor. 1.0 is normal speed, 1.5 is 50% faster, etc.
# We are using 1.15 as requested.
SPEED_FACTOR="1.15"

# --- Script Logic ---

# Exit immediately if a command exits with a non-zero status.
set -e

# Create the output directory if it doesn't already exist.
# The -p flag prevents an error if the directory already exists.
echo "Creating output directory (if it doesn't exist): $OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

# Loop through each of the specified input directories.
for DIR in "${INPUT_DIRS[@]}"; do
    # Check if the input directory actually exists before trying to search it.
    if [ -d "$DIR" ]; then
        echo "--------------------------------------------------"
        echo "Processing files in: $DIR"
        echo "--------------------------------------------------"

        # Use 'find' to locate all files ending with .mp3 in the current directory and its subdirectories.
        # The output of find is piped to a 'while' loop to process each file individually.
        # This method correctly handles filenames that may contain spaces.
        find "$DIR" -type f -name "*.mp3" | while read -r FILE; do
            # Get just the filename from the full path (e.g., "audio.mp3" from "/path/to/audio.mp3").
            FILENAME=$(basename "$FILE")

            # Print a message to the console to show which file is being processed.
            echo "Processing: $FILENAME"

            # Use ffmpeg to process the file.
            # -nostdin: **(THE FIX)** Prevents ffmpeg from reading from stdin, which fixes the loop issue.
            # -i: Specifies the input file.
            # -filter:a "atempo=...": Applies an audio filter to change the tempo (speed) without altering the pitch.
            # -vn: Specifies that there is no video, which can prevent potential errors with MP3 files.
            # -y: Overwrite output files without asking.
            # The final argument is the full path for the output file.
            ffmpeg -y -nostdin -i "$FILE" -filter:a "atempo=$SPEED_FACTOR" -vn "$OUTPUT_DIR/$FILENAME"
        done
    else
        echo "Warning: Input directory not found, skipping: $DIR"
    fi
done

echo "--------------------------------------------------"
echo "All done! Sped-up files are in: $OUTPUT_DIR"
echo "--------------------------------------------------"