package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open json file
	file, err := os.Open("gopher.json")
	if err != nil {
		fmt.Println("Open file error:", err)
	}
	defer file.Close()

	// Parse json data to a map of Chapter
	m, err := parseJSON(file)
	if err != nil {
		fmt.Println("Parse JSON error:", err)
	}

	chapter := m["intro"]
	printChapter(chapter)
	// Read user input
	fmt.Print("Enter the number: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		opt, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Print("Invalid input. Please reenter the number: ")
			continue
		}

		// Check if user input within the range, option starts from 1 to number of options
		if opt < 1 || opt > len(chapter.Options) {
			fmt.Print("Invalid option. Please reenter the number: ")
			continue
		}

		nextArc := chapter.Options[opt-1].Arc
		chapter = m[nextArc]
		printChapter(chapter)

		if len(chapter.Options) == 0 {
			fmt.Println("------- END OF STORY -------")
			break
		} else {
			fmt.Print("Enter the number: ")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func parseJSON(file *os.File) (map[string]Chapter, error) {
	jsonBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var chapters map[string]Chapter
	err = json.Unmarshal(jsonBytes, &chapters)
	if err != nil {
		return nil, err
	}

	return chapters, nil
}

func printChapter(chapter Chapter) {
	fmt.Println("--------------------------------------")
	fmt.Println(chapter.Title)
	fmt.Println("--------------------------------------")
	for _, s := range chapter.Story {
		fmt.Println(s)
	}
	for i, s := range chapter.Options {
		fmt.Println(i+1, " - ", s.Text)
	}
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
