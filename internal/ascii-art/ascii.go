package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// @done check args if input is a ascii valid character (btw 32 <> 126)
// @done check if `\n`, return to line (print a table after another)
// @done check if heigh is correct
// @done if empty, return one \n
// @done > not needed git config core search for auto-crlf (no conversion wanted)
// todo split the main file (ASCII), create the functions osrgsToString, fileToMultiArray in external lib
// goal: go run . ""

// takes an argument and print it as ascii art
func Ascii(input string, theme string) string {
	// READ arguments
	// args := os.Args[1:]
	// Check arguments and return a string
	// input := ""
	// ? debug only
	// fmt.Println("theme:", theme, "input:", input)
	// Open the theme file
	theme = AsciiTheme(theme)
	theme_file, err := os.Open(theme)
	if err != nil {
		panic(err)
	}
	// Close the theme file (theme_file) on exit and check for its returned error
	defer func() {
		if err := theme_file.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
			os.Exit(1)
		}
	}()
	// Read lines using bufio.Scanner
	scanner := bufio.NewScanner(bufio.NewReader(theme_file))
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		return bufio.ScanLines(data, atEOF)
	})
	// Build the multi-dimensional array
	var series [][]string
	var currentSeries []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // Empty line indicates a new series
			if len(currentSeries) > 0 {
				series = append(series, currentSeries)
				currentSeries = nil // Reset for the next series
			}
		} else {
			currentSeries = append(currentSeries, line)
		}
	}
	// Handle potential last series without empty line at the end
	if len(currentSeries) > 0 {
		series = append(series, currentSeries)
	}

	// Detect if input contains one or multiple \n
	multi := strings.Split(input, "\\n") // split on `\n`
	if len(multi) > 1 {
		var result_multi string
		for _, str := range multi {
			if len(str) == 0 { // if word is empty, return `\n`
				result_multi += "\n"
			} else {
				result_multi += InputToASCII(str, series)
			}
		}
		return result_multi
	} else {
		return InputToASCII(input, series)
	}
}

// accept a 2D-array representing ASCII art (theme) and print it
func InputToASCII(input string, series [][]string) string {
	var result string
	indexes := []rune(input)
	// conversion, first ascii letter 'space' dec:32 becomes dec:0
	for i, rune := range indexes {
		indexes[i] = rune - 32
	}
	// First loop, 8 lines from the theme
	for i := 0; i < 8; i++ {
		// Second loop each characters from the input
		for _, lineIndex := range indexes {
			if int(lineIndex) >= len(series) {
				result += " " // Placeholder for missing series
				continue
			}
			if i < len(series[lineIndex]) {
				result += series[lineIndex][i]
			} else {
				result += " " // Placeholder for missing lines within a series
			}
		}
		result += "\n"
	}
	return result
}

func AsciiTheme(str string) (theme string) {
	switch str {
	case "Standard":
		theme = "internal/ascii-art/src/themes/standard.txt"
	case "Shadow":
		theme = "internal/ascii-art/src/themes/shadow.txt"
	case "Thinkertoy":
		theme = "internal/ascii-art/src/themes/thinkertoy.txt"
	}
	return theme
}

func Error(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
