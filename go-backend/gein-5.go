package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
)

// Dynamic Configuration Models completely divorced from layout strings
type Config struct {
	Meta struct {
		PageTitle  string `json:"page_title"`
		ThemeClass string `json:"theme_class"`
	} `json:"meta"`
	Nav struct {
		LogoText string `json:"logo_text"`
		Links    []struct {
			Label string `json:"label"`
			URL   string `json:"url"`
		} `json:"links"`
	} `json:"nav"`
	Hero struct {
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
	} `json:"hero"`
	About struct {
		Title      string   `json:"title"`
		AlertText  string   `json:"alert_text,omitempty"`
		Paragraphs []string `json:"paragraphs"`
	} `json:"about"`
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [config.json] [output.html]")
		os.Exit(1)
	}

	configFile := os.Args[1]
	outputFile := os.Args[2]

	// 1. Read the data parameter input file dynamically from disk
	configBytes, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}

	var data Config
	if err := json.Unmarshal(configBytes, &data); err != nil {
		fmt.Printf("Error parsing configuration: %v\n", err)
		os.Exit(1)
	}

	// 2. Parse layout components
	tmpl := template.Must(template.ParseFiles("layout.html"))

	// 3. Output standard rendered file code
	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	tmpl.Execute(out, data)
	fmt.Printf("Successfully templated %s -> %s\n", configFile, outputFile)
}