# Libp2p ED25519 Key Management Tool

This tool facilitates the management of ED25519 cryptographic keys used in libp2p applications. It provides functionality to generate a PeerID from a private key and to encode a private key in base64 format, making it easier to handle and store.

## Prerequisites

Before you begin, ensure you have installed Go (version 1.21 or later is recommended). You will also need to have the `libp2p` Go modules available in your project.

## Installation

To use this tool, first clone the repository to your local machine. If the code is packaged separately, ensure it is placed in a directory where you can run Go commands. Dependencies are managed using Go modules, so they should be automatically resolved when you build or run the tool.

## Usage

The tool supports two commands:

- `peerid`: Generate a PeerID from the provided ED25519 private key file.
- `privatekey`: Encode the ED25519 private key file into a base64 string.

To run the tool, navigate to the directory containing the code and use the following syntax:

```bash
go run main.go <command> <path_to_keypair_file>
