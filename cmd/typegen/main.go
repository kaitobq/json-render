package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"kaitobq/json-render/pkg/typegen"
)

func main() {
	var (
		inputDir         = flag.String("input", ".", "Input directory (contains Go files)")
		outputDir        = flag.String("output", "./generated", "Output directory")
		useInterface     = flag.Bool("interface", false, "Generate as interface (false generates as type)")
		useReadonly      = flag.Bool("readonly", false, "Generate as readonly properties")
		namingConvention = flag.String("naming", "camelCase", "Naming convention (camelCase/snake_case)")
	)

	flag.Parse()

	// Create output directory
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Search for .go files in the input directory (only direct files, no subdirectories)
	entries, err := os.ReadDir(*inputDir)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	var goFiles []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") {
			goFiles = append(goFiles, filepath.Join(*inputDir, entry.Name()))
		}
	}

	if len(goFiles) == 0 {
		log.Fatalf("No Go files found in input directory: %s", *inputDir)
	}

	// Process each file
	for _, inputFile := range goFiles {
		// Generate output file name
		base := filepath.Base(inputFile)
		outputFile := strings.TrimSuffix(base, ".go") + ".ts"

		// Set generation options
		opts := &typegen.Options{
			UseInterface:     *useInterface,
			UseReadonly:      *useReadonly,
			OutputFileName:   outputFile,
			NamingConvention: typegen.NamingConvention(*namingConvention),
		}

		// Create generator and run
		generator := typegen.NewGenerator(inputFile, *outputDir, opts)
		if err := generator.Generate(); err != nil {
			log.Printf("Failed to generate file %s: %v", inputFile, err)
			continue
		}

		fmt.Printf("Generated TypeScript definitions: %s\n",
			filepath.Join(*outputDir, outputFile))
	}
}
