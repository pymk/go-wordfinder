package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

type SearchConfig struct {
	CaseSensitive bool
	WholeWord     bool
}

type SearchResult struct {
	LineNumber  int
	CurrentLine string
}

func searchWord(filePath, searchTerm string, config SearchConfig) ([]SearchResult, error) {
	if searchTerm == "" {
		return nil, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := make([]SearchResult, 0)
	compiledPattern, err := buildSearchPattern(searchTerm, config)
	if err != nil {
		return nil, fmt.Errorf("invalid search pattern: %w", err)
	}

	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		currentLine := scanner.Text()
		if compiledPattern.MatchString(currentLine) {
			result = append(result, SearchResult{
				LineNumber:  lineNumber,
				CurrentLine: currentLine,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	return result, nil
}

func buildSearchPattern(term string, config SearchConfig) (*regexp.Regexp, error) {
	pattern := regexp.QuoteMeta(term)
	if config.WholeWord {
		pattern = `\b` + pattern + `\b`
	}

	if !config.CaseSensitive {
		return regexp.Compile("(?i)" + pattern)
	}
	return regexp.Compile(pattern)
}

func printResults(results []SearchResult) {
	if len(results) == 0 {
		fmt.Println("No matches found")
		return
	}

	for _, result := range results {
		fmt.Printf("%6d: %s\n", result.LineNumber, result.CurrentLine)
	}
}

func main() {
	filePath := flag.String("file", "", "Path to the file")
	searchTerm := flag.String("term", "", "Term of interest")
	caseSensitive := flag.Bool("case-sensitive", true, "Match with case-sensitivity")
	wholeWord := flag.Bool("whole-word", true, "Match whole words only")
	flag.Parse()

	if *filePath == "" || *searchTerm == "" {
		flag.Usage()
		os.Exit(1)
	}

	config := SearchConfig{
		CaseSensitive: *caseSensitive,
		WholeWord:     *wholeWord,
	}

	result, err := searchWord(*filePath, *searchTerm, config)
	if err != nil {
		log.Fatalf("Failed to search: %v", err)
	}

	printResults(result)
}
