package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	filename := ""
	homeDir, _ := os.UserHomeDir()
	shell := os.Getenv("SHELL")
	switch shell {
	case "/bin/zsh":
		filename = filepath.Join(homeDir, ".zshrc")
	case "/bin/bash":
		filename = filepath.Join(homeDir, ".bashrc")
	default:
		fmt.Fprintf(os.Stderr, "Error unsupported shell: %s\n", shell)
		os.Exit(1)
	}

	aliases, err := readAliases(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filename, err)
		os.Exit(1)
	}

	fmt.Println("ichimoji alias list.")

	for ch := 'a'; ch <= 'z'; ch++ {
		letter := string(ch)
		if value, exists := aliases[letter]; exists {
			fmt.Printf("✅ %s='%s'\n", letter, value)
		} else {
			fmt.Printf("🈳 %s\n", letter)
		}
	}
}

func readAliases(filename string) (map[string]string, error) {
	aliases := make(map[string]string)
	file, err := os.Open(filename)
	if err != nil {
		return aliases, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	aliasRegex := regexp.MustCompile(`^alias\s+([a-z])=['"](.+)['"]$`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if matches := aliasRegex.FindStringSubmatch(line); len(matches) == 3 {
			letter := matches[1]
			value := matches[2]
			aliases[letter] = value
		}
	}

	return aliases, nil
}
