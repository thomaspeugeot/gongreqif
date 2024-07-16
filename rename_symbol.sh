#!/bin/bash

# Define the file and the symbols
FILE_PATH="go/models/models.go"
OLD_SYMBOL="TYPE"
NEW_SYMBOL="REQTYPE"

# Define the line and character position of the symbol
LINE=66
CHAR=23

# Get the URI of the file
URI=$(realpath "$FILE_PATH")
URI="file://$URI"

# Start gopls server
gopls serve &

# Give the server some time to start
sleep 2

# Create the rename request JSON payload
REQUEST=$(jq -n --arg URI "$URI" --arg NEW_SYMBOL "$NEW_SYMBOL" --argjson LINE $LINE --argjson CHAR $CHAR '
{
    "method": "textDocument/rename",
    "params": {
        "textDocument": { "uri": $URI },
        "position": { "line": $LINE, "character": $CHAR },
        "newName": $NEW_SYMBOL
    }
}')

# Send the rename request to the gopls server
RESPONSE=$(echo "$REQUEST" | curl -s -X POST --data-binary @- http://localhost:3737)

# Stop the gopls server
kill %1

echo "Rename response: $RESPONSE"
