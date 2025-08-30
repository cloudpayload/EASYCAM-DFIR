package parser

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

// Parser handles the parsing of log files and conversion to normalized events
type Parser struct {
	MapsDir string
}

// NewParser creates a new Parser instance
func NewParser(mapsDir string) *Parser {
	return &Parser{
		MapsDir: mapsDir,
	}
}

// ParseDirectory scans a directory for log files and processes them
func (p *Parser) ParseDirectory(inputDir string) error {
	fmt.Printf("Scanning directory: %s\n", inputDir)
	
	// Walk through the input directory
	err := filepath.WalkDir(inputDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		// Skip directories
		if d.IsDir() {
			return nil
		}
		
		// Check if it's a log file (common extensions)
		if isLogFile(path) {
			fmt.Printf("Processing log file: %s\n", path)
			if err := p.processLogFile(path); err != nil {
				fmt.Printf("Error processing %s: %v\n", path, err)
			}
		}
		
		return nil
	})
	
	return err
}

// isLogFile checks if a file is likely a log file based on extension
func isLogFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	logExtensions := []string{".log", ".json", ".csv", ".txt", ".xml", ".yaml", ".yml"}
	
	for _, logExt := range logExtensions {
		if ext == logExt {
			return true
		}
	}
	
	// Also check if filename contains common log patterns
	filename := strings.ToLower(filepath.Base(path))
	logPatterns := []string{"log", "audit", "trail", "event", "access"}
	
	for _, pattern := range logPatterns {
		if strings.Contains(filename, pattern) {
			return true
		}
	}
	
	return false
}

// processLogFile processes a single log file
func (p *Parser) processLogFile(filePath string) error {
	// TODO: Implement log file processing
	// 1. Read file content
	// 2. Determine log type using detector patterns
	// 3. Apply appropriate mapping configuration
	// 4. Convert to normalized events
	
	return nil
}
