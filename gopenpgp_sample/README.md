# GoPenPGP File Encryption Sample

This Go application encrypts files using the ProtonMail GoPenPGP v3 library.

## Prerequisites

- Go 1.22.5 or later
- gpg command for key generation and decryption verification

  ```bash
  brew install gpg
  ```

## Usage

### 1. Generate PGP Key Pair using GPG

Using the provided script:

```bash
cd gopenpgp_sample
./generate_keys.sh
```

### 2. Encrypt a File

```bash
go run main.go encrypt test.txt ./keys/public_key.asc
```

This will create `test.txt.pgp` encrypted file.

### 3. Decrypt a File

Using Go application:

```bash
go run main.go decrypt test.txt.pgp keys/private_key.asc ""

# With passphrase
go run main.go decrypt test.txt.pgp keys/private_key.asc "your-passphrase"
```

Using gpg command:

```bash
gpg --decrypt test.txt.pgp
gpg --decrypt test.txt.pgp > test.txt.pgp.decrypted
```

## Clean up

```bash
./clean.sh
```
