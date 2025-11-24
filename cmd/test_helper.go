package main

import (
	"fmt"
	"os"

	"github.com/HORNET-Storage/Scionic-Merkle-Tree/v2/dag"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: test_helper <command> [args...]")
		fmt.Println("Commands:")
		fmt.Println("  create <input_path> <output_cbor>  - Create DAG and save to CBOR")
		fmt.Println("  verify <cbor_path>                  - Load and verify DAG from CBOR")
		fmt.Println("  info <input_path>                   - Print DAG info")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "create":
		if len(os.Args) < 4 {
			fmt.Println("Usage: test_helper create <input_path> <output_cbor>")
			os.Exit(1)
		}
		inputPath := os.Args[2]
		outputCbor := os.Args[3]

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

		fmt.Printf("Success! Root: %s, Leaves: %d\n", d.Root, len(d.Leafs))

	case "verify":
		if len(os.Args) < 3 {
			fmt.Println("Usage: test_helper verify <cbor_path>")
			os.Exit(1)
		}
		cborPath := os.Args[2]

		// Load CBOR
		data, err := os.ReadFile(cborPath)
		if err != nil {
			fmt.Printf("Error reading CBOR file: %v\n", err)
			os.Exit(1)
		}

		d, err := dag.FromCBOR(data)
		if err != nil {
			fmt.Printf("Error deserializing CBOR: %v\n", err)
			os.Exit(1)
		}

		// Verify
		err = d.Verify()
		if err != nil {
			fmt.Printf("Error verifying DAG: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Success! Root: %s, Leaves: %d\n", d.Root, len(d.Leafs))

	case "info":
		if len(os.Args) < 3 {
			fmt.Println("Usage: test_helper info <input_path>")
			os.Exit(1)
		}
		inputPath := os.Args[2]

		// Create DAG
		d, err := dag.CreateDag(inputPath, false)
		if err != nil {
			fmt.Printf("Error creating DAG: %v\n", err)
			os.Exit(1)
		}

		// Print info
		fmt.Printf("Root: %s\n", d.Root)
		fmt.Printf("Leaves: %d\n", len(d.Leafs))
		fmt.Printf("\nLeaf details:\n")
		for hash, leaf := range d.Leafs {
			fmt.Printf("  %s: type=%s name=%s links=%d\n",
				hash[:20], leaf.Type, leaf.ItemName, leaf.CurrentLinkCount)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
