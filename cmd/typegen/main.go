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
		useInterface     = flag.Bool("interface", true, "Generate as interface (false generates as type)")
		useReadonly      = flag.Bool("readonly", false, "Generate as readonly properties")
		namingConvention = flag.String("naming", "camelCase", "Naming convention (camelCase/snake_case)")
	)

	flag.Parse()

	// Create output directory
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Search all .go files in input directory
	var goFiles []string
	err := filepath.Walk(*inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			goFiles = append(goFiles, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to search files: %v", err)
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
