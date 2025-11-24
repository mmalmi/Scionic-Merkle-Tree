package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/HORNET-Storage/Scionic-Merkle-Tree/v2/dag"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: test_helper_chunk <command> [args...]")
		fmt.Println("Commands:")
		fmt.Println("  create <input_path> <output_cbor> <chunk_size>  - Create DAG with specific chunk size")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "create":
		if len(os.Args) < 5 {
			fmt.Println("Usage: test_helper_chunk create <input_path> <output_cbor> <chunk_size>")
			os.Exit(1)
		}
		inputPath := os.Args[2]
		outputCbor := os.Args[3]
		chunkSize, err := strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Printf("Invalid chunk size: %v\n", err)
			os.Exit(1)
		}

		// Set chunk size
		if chunkSize <= 0 {
			dag.DisableChunking()
		} else {
			dag.SetChunkSize(chunkSize)
		}

		// Create DAG
		d, err := dag.CreateDag(inputPath, false)
		if err != nil {
			fmt.Printf("Error creating DAG: %v\n", err)
			os.Exit(1)
		}

		// Verify
		err = d.Verify()
		if err != nil {
			fmt.Printf("Error verifying DAG: %v\n", err)
			os.Exit(1)
		}

		// Save to CBOR
		cbor, err := d.ToCBOR()
		if err != nil {
			fmt.Printf("Error serializing to CBOR: %v\n", err)
			os.Exit(1)
		}

		err = os.WriteFile(outputCbor, cbor, 0644)
		if err != nil {
			fmt.Printf("Error writing CBOR file: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Success! Root: %s, Leaves: %d, ChunkSize: %d\n", d.Root, len(d.Leafs), chunkSize)

		// Reset to default
		dag.SetDefaultChunkSize()

	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
