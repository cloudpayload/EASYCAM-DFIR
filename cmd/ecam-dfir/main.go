package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	
	"ecam-dfir/internal/parser"
)

func main() {
	// Define command line flags
	inputDir := flag.String("input-dir", "", "Directory containing exported logs to parse")
	outputFile := flag.String("output", "events.csv", "Output CSV file for normalized events")
	
	// Parse command line arguments
	flag.Parse()
	
	// Validate required input directory
	if *inputDir == "" {
		fmt.Println("Error: input-dir is required")
		fmt.Println("Usage: ecam-dfir --input-dir /path/to/exported-logs --output events.csv")
		flag.PrintDefaults()
		os.Exit(1)
	}
	
	// Check if input directory exists
	if _, err := os.Stat(*inputDir); os.IsNotExist(err) {
		log.Fatalf("Input directory does not exist: %s", *inputDir)
	}
	
	fmt.Printf("Parsing logs from: %s\n", *inputDir)
	fmt.Printf("Output file: %s\n", *outputFile)
	
	// Create parser instance
	parser := parser.NewParser("maps")
	
	// Parse the input directory
	if err := parser.ParseDirectory(*inputDir); err != nil {
		log.Fatalf("Error parsing directory: %v", err)
	}
	
	fmt.Println("Log parsing completed successfully!")
}
