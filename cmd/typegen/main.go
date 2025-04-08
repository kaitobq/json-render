package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"kaitobq/json-render/pkg/typegen"
)

func main() {
	var (
		inputDir         = flag.String("input", ".", "Input directory (directory containing Go files)")
		outputDir        = flag.String("output", "./generated", "Output directory")
		useInterface     = flag.Bool("interface", false, "Generate as interface (false generates type)")
		useReadonly      = flag.Bool("readonly", false, "Generate as readonly property")
		outputFileName   = flag.String("filename", "types.d.ts", "Output file name")
		namingConvention = flag.String("naming", "camelCase", "Naming convention (camelCase/PascalCase/snake_case)")
	)

	flag.Parse()

	// Create output directory
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Set generation options
	opts := &typegen.Options{
		UseInterface:     *useInterface,
		UseReadonly:      *useReadonly,
		OutputFileName:   *outputFileName,
		NamingConvention: typegen.NamingConvention(*namingConvention),
	}

	// Create generator
	generator := typegen.NewGenerator(*inputDir, *outputDir, opts)

	// Generate type definitions
	if err := generator.Generate(); err != nil {
		log.Fatalf("Failed to generate type definitions: %v", err)
	}

	fmt.Printf("TypeScript type definitions generated: %s\n",
		filepath.Join(*outputDir, *outputFileName))
}
