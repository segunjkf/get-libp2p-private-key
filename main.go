package main

import (
    "encoding/base64"
    "fmt"
    "io"
    "os"

    "github.com/libp2p/go-libp2p/core/crypto"
    "github.com/libp2p/go-libp2p/core/peer"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run main.go <command> <path_to_keypair_file>")
        os.Exit(1)
    }

    command := os.Args[1]
    keypairPath := os.Args[2]

	// Open the file
    file, err := os.Open(keypairPath)
    if err != nil {
        fmt.Printf("Error opening keypair file: %v\n", err)
        os.Exit(1)
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        fmt.Printf("Error reading keypair file: %v\n", err)
        os.Exit(1)
    }

    // Assuming `data` contains the raw ED25519 private key bytes directly.
    switch command {
    case "peerid":
        generatePeerID(data)
    case "privatekey":
        encodePrivateKey(data)
    default:
        fmt.Println("Invalid command. Available commands: peerid, privatekey")
        os.Exit(1)
    }
}

func generatePeerID(keyData []byte) {
    // Assuming keyData is the raw ED25519 private key bytes.
    privKey, err := crypto.UnmarshalEd25519PrivateKey(keyData)
    if err != nil {
        fmt.Printf("Error creating ED25519 private key: %v\n", err)
        return
    }

    pid, err := peer.IDFromPrivateKey(privKey)
    if err != nil {
        fmt.Printf("Error generating PeerID from private key: %v\n", err)
        return
    }

    fmt.Println("PeerID:", pid.String())
}

func encodePrivateKey(keyData []byte) {
    // Assuming keyData is the raw ED25519 private key bytes.
    privKey, err := crypto.UnmarshalEd25519PrivateKey(keyData)
    if err != nil {
        fmt.Printf("Error creating ED25519 private key: %v\n", err)
        return
    }

    // Serialize the private key to bytes
    privBytes, err := crypto.MarshalPrivateKey(privKey)
    if err != nil {
        fmt.Printf("Error marshalling private key to protobuf: %v\n", err)
        return
    }

    // Encode the serialized private key bytes to base64
    encoded := base64.StdEncoding.EncodeToString(privBytes)
    fmt.Println("Base64 Encoded Private Key:", encoded)
}
