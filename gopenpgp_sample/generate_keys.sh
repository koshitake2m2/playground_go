#!/bin/bash

# Generate GPG key pair for testing
echo "Generating GPG key pair..."

# Clean up existing keys first
echo "Cleaning up existing keys..."
./clean.sh > /dev/null 2>&1

# Create a batch file for GPG key generation
mkdir -p ./tmp
mkdir -p ./keys

cat > ./tmp/gpg_batch.txt << EOF
%no-protection
Key-Type: RSA
Key-Length: 4096
Subkey-Type: RSA
Subkey-Length: 4096
Name-Real: Test User
Name-Email: test@example.com
Expire-Date: 0
EOF

# Generate the key
gpg --batch --generate-key ./tmp/gpg_batch.txt

# Export public key (should only be one now)
gpg --armor --export test@example.com > ./keys/public_key.asc

# Export private key
gpg --armor --export-secret-keys test@example.com > ./keys/private_key.asc

# Clean up
rm ./tmp/gpg_batch.txt

echo "Keys generated successfully!"
echo "Public key:  keys/public_key.asc"
echo "Private key: keys/private_key.asc"