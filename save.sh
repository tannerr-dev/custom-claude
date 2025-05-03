#!/bin/bash

# Check if correct number of arguments provided
# if [ $# -ne 2 ]; then
#     echo "Usage: $0 source_file destination_file"
#     echo "  source_file: File to copy from and then clear"
#     echo "  destination_file: File to append contents to"
#     exit 1
# fi
#
# SOURCE_FILE="$1"
# DEST_FILE="$2"

#Check if correct number of arguments provided
if [ $# -ne 1 ]; then
    echo "$0 ERROR: specify destination..."
    exit 1
fi

DEST_FILE="./archive/$1"

SOURCE_FILE="msg.md"
# DEST_FILE="log.md"

# Check if source file exists
if [ ! -f "$SOURCE_FILE" ]; then
    echo "Error: Source file '$SOURCE_FILE' does not exist."
    exit 1
fi

# Check if source file is readable
if [ ! -r "$SOURCE_FILE" ]; then
    echo "Error: Cannot read from source file '$SOURCE_FILE'."
    exit 1
fi

# Check if destination file is writable (or can be created)
touch "$DEST_FILE.md" 2>/dev/null
if [ ! -w "$DEST_FILE.md" ]; then
    echo "Error: Cannot write to destination file '$DEST_FILE'."
    exit 1
fi

# Append source file contents to destination file
cat "$SOURCE_FILE" >> "$DEST_FILE.md"
if [ $? -ne 0 ]; then
    echo "Error: Failed to append contents to '$DEST_FILE.md'."
    exit 1
fi

# Clear the source file
# > "$SOURCE_FILE"
# if [ $? -ne 0 ]; then
#     echo "Error: Failed to clear contents of '$SOURCE_FILE'."
#     exit 1
# fi

# echo "Successfully appended contents from '$SOURCE_FILE' to '$DEST_FILE' and cleared '$SOURCE_FILE'."
echo "Successfully saved '$SOURCE_FILE' to '$DEST_FILE.md' ..."
exit 0
