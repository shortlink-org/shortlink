#!/bin/bash

# Define the directory to search
search_dir="../../../"

# Loop over all logo.svg files in the directory and its subdirectories
find "$search_dir" -path '*/docs/public/logo.svg' | while read file; do
    # Extract the service name from the file path
    nameService=$(dirname "$file" | awk -F/ '{print $(NF-2)}')

    # Define the new file name
    newFileName="${nameService}.svg"

    # Copy and rename the file to the current directory
    cp "$file" "./logos/$newFileName"
done
