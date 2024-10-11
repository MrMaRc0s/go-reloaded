package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run <program.go> <input_file> <output_file>")
		return
	}
	if len(os.Args) > 3 {
		fmt.Println("Too many arguments")
		return
	}
	inputFilename := os.Args[1]
	outputFilename := os.Args[2]

	content, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		fmt.Println("Error reading input file")
		return
	}

	var Words []string
	var n int
	Words = SplitWhiteSpaces(string(content))

	for i := 0; i < len(Words); i++ {
		if i+1 <= len(Words)-1 {
			Runes := []rune((Words[i+1]))
			if Runes[len(Runes)-1] == ')' {
				var str string = string(Runes[:len(Runes)-1])
				x, _ := strconv.Atoi(str)
				n = x
			}
		}
		if Words[i] == "(cap)" || Words[i] == "(cap," {
			for loop := 0; loop <= n; loop++ {
				if loop <= i {
					Words[i-loop] = Capitalize(Words[i-loop])
				}
			}
			if n == 0 {
				Words[i-1] = Capitalize(Words[i-1])
			}
			Words[i] = ""
			if n > 0 {
				Words[i+1] = ""
			}
		} else if Words[i] == "(up)" || Words[i] == "(up," {
			for loop := 0; loop <= n; loop++ {
				if loop <= i {
					Words[i-loop] = ToUpper(Words[i-loop])
				}
			}
			if n == 0 {
				Words[i-1] = ToUpper(Words[i-1])
			}
			Words[i] = ""
			if n > 0 {
				Words[i+1] = ""
			}
		} else if Words[i] == "(low)" || Words[i] == "(low," {
			for loop := 0; loop <= n; loop++ {
				if loop <= i {
					Words[i-loop] = ToLower(Words[i-loop])
				}
			}
			if n == 0 {
				Words[i-1] = ToLower(Words[i-1])
			}
			Words[i] = ""
			if n > 0 {
				Words[i+1] = ""
			}
		} else if Words[i] == "(bin)" {
			Words[i-1] = BinConv((Words[i-1]))
			Words[i] = ""
		} else if Words[i] == "(hex)" {
			Words[i-1] = HexConv((Words[i-1]))
			Words[i] = ""
		} else if Words[i] == "a" || Words[i] == "A" {
			if i < len(Words)-1 {
				if checkifvowel(Words[i+1]) {
					Words[i] = Words[i] + "n"
				}
			}
		}
	}

	finalContent := strings.Join(Words, " ")
	finalContent = formatPunctuation(finalContent)
	finalContent = formatQuotes(finalContent)
	finalContent = trimExtraSpaces(finalContent)

	// Write output to the specified file
	err = ioutil.WriteFile(outputFilename, []byte(finalContent), 0o644)
	if err != nil {
		fmt.Println("Error writing to output file")
		return
	} else {
		fmt.Println("Writing completed")
	}
}

func formatPunctuation(input string) string {
	// Remove spaces before punctuation and ensure punctuation is adjacent to the previous word
	re := regexp.MustCompile(`\s*([.,!?;:]+)`)
	result := re.ReplaceAllString(input, "$1")

	// Ensure exactly one space follows punctuation, if not at the end of the string
	result += " " // Add a space to the end to handle the last punctuation correctly
	re2 := regexp.MustCompile(`([.,!?;:]+)([^\s])`)
	result = re2.ReplaceAllString(result, "$1 $2")

	// Remove any extra spaces after punctuation marks
	re3 := regexp.MustCompile(`([.,!?;:]+)\s+`)
	result = re3.ReplaceAllString(result, "$1 ")

	return strings.TrimSpace(result) // Trim any leading/trailing spaces
}

func trimExtraSpaces(input string) string {
	re := regexp.MustCompile(`\s+`)
	result := re.ReplaceAllString(input, " ")
	return strings.TrimSpace(result)
}

func formatQuotes(input string) string {
	re := regexp.MustCompile(`'([^']*)'`)
	result := re.ReplaceAllStringFunc(input, func(match string) string {
		cleaned := strings.TrimSpace(match[1 : len(match)-1])
		return "'" + cleaned + "'"
	})
	return result
}

func IsUpper(s string) bool {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] < 'A' || slice[i] > 'Z' {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] < 'a' || slice[i] > 'z' {
			return false
		}
	}
	return true
}

func IsNumeric(s string) bool {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] < '0' || slice[i] > '9' {
			return false
		}
	}
	return true
}

func ToUpper(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'a' && c <= 'z' {
			b[i] = c - ('a' - 'A')
		}
	}
	return string(b)
}

func ToLower(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'A' && c <= 'Z' {
			b[i] = c + ('a' - 'A')
		}
	}
	return string(b)
}

func Capitalize(s string) string {
	var str string
	first := true
	for i := 0; i < len(s); i++ {
		if IsLower(string(s[i])) || IsUpper(string(s[i])) {
			if first {
				str += ToUpper(string(s[i]))
				first = false
			} else {
				str += ToLower(string(s[i]))
			}
		} else {
			str += string(s[i])
			first = !IsNumeric(string(s[i]))
		}
	}
	return str
}

func SplitWhiteSpaces(s string) []string {
	var word string
	var result []string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' {
			word += string(s[i])
		} else {
			if word != "" {
				result = append(result, word)
			}
			word = ""
		}
		if i == len(s)-1 && s[i] != ' ' && s[i] != '\t' && s[i] != '\n' {
			result = append(result, word)
		}
	}
	return result
}

func BinConv(s string) string {
	bin, _ := strconv.ParseInt(s, 2, 64)
	return strconv.FormatInt(bin, 10)
}

func HexConv(s string) string {
	dec, _ := strconv.ParseInt(s, 16, 64)
	return strconv.FormatInt(dec, 10)
}

func checkifvowel(s string) bool {
	firstrune := []rune((s))
	return firstrune[0] == 'a' || firstrune[0] == 'e' || firstrune[0] == 'o' || firstrune[0] == 'i' || firstrune[0] == 'u' || firstrune[0] == 'h'
}
