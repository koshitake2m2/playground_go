#!/bin/bash

# Clean up existing GPG keys for test@example.com
echo "Cleaning up existing GPG keys for test@example.com..."

# Get all key fingerprints for test@example.com
FINGERPRINTS=$(gpg --with-colons --fingerprint test@example.com | grep "^fpr:" | cut -d':' -f10)

if [ -z "$FINGERPRINTS" ]; then
    echo "No existing keys found for test@example.com"
else
    echo "Found existing keys, deleting..."
    
    # Delete each key
    for fingerprint in $FINGERPRINTS; do
        echo "Deleting key: $fingerprint"
        
        # Delete secret key first (if exists)
        gpg --batch --yes --delete-secret-key "$fingerprint" 2>/dev/null
        
        # Delete public key
        gpg --batch --yes --delete-key "$fingerprint" 2>/dev/null
    done
    
    echo "All existing keys for test@example.com have been deleted"
fi

# Clean up key files
if [ -d "keys" ]; then
    echo "Cleaning up key files..."
    rm -rf ./keys
    echo "Key files cleaned"
fi

# Clean up temporary files
if [ -d "tmp" ]; then
    echo "Cleaning up temporary files..."
    rm -rf ./tmp
    echo "Temporary files cleaned"
fi

# Clean up encrypted test files
echo "Cleaning up test files..."
rm -f ./test.txt.pgp ./test.txt.pgp.decrypted

echo "Cleanup complete!"