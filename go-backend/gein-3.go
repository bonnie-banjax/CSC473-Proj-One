package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"os"
)

func main() {
	// Define standard CLI flags
	tmplPath := flag.String("template", "template.html", "Path to the Go HTML template file")
	dataPath := flag.String("data", "", "Path to the JSON configuration data file")
	flag.Parse()

	// Validation
	if *dataPath == "" {
		fmt.Fprintln(os.Stderr, "Error: -data flag is required.")
		flag.Usage()
		os.Exit(1)
	}

	// 1. Read and parse the JSON input dataset
	jsonData, err := os.ReadFile(*dataPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading data file: %v\n", err)
		os.Exit(1)
	}

	// Unmarshal into an anonymous interface map so it accepts any structure matching the template
	var data interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON payload: %v\n", err)
		os.Exit(1)
	}

	// 2. Read and parse the target Go HTML Template
	tmpl, err := template.ParseFiles(*tmplPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing template file: %v\n", err)
		os.Exit(1)
	}

	// 3. Render directly to standard output
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing template resolution: %v\n", err)
		os.Exit(1)
	}
}