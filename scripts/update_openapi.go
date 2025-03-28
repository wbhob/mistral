package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

const (
	commitURL = "https://api.github.com/repos/mistralai/platform-docs-public/commits?path=openapi.yaml&per_page=1"
	specURL   = "https://raw.githubusercontent.com/mistralai/platform-docs-public/main/openapi.yaml"
)

func main() {
	// Get latest version from commit message
	resp, err := http.Get(commitURL)
	if err != nil {
		fmt.Printf("Error getting commit: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var commits []struct {
		Commit struct {
			Message string `json:"message"`
		} `json:"commit"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&commits); err != nil {
		fmt.Printf("Error parsing commit: %v\n", err)
		os.Exit(1)
	}

	re := regexp.MustCompile(`v\d+\.\d+\.\d+`)
	version := re.FindString(commits[0].Commit.Message)
	if version != "" {
		fmt.Printf("::set-output name=version::%s\n", version)
	}

	// Download the spec
	resp, err = http.Get(specURL)
	if err != nil {
		fmt.Printf("Error downloading spec: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	out, err := os.Create("openapi.yaml")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}
}
