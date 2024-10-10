package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	filename := os.Args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	var Words []string
	Words = SplitWhiteSpaces(string(content))
	fmt.Println(Words[2])

	for i := 0; i < len(Words); i++ {
		if Words[i] == "(cap)" {
			Words[i-1] = Capitalize(Words[i-1])
			Words[i] = ""
		} else if Words[i] == "(up)" {
			Words[i-1] = ToUpper(Words[i-1])
			Words[i] = ""
		} else if Words[i] == "(low)" {
			Words[i-1] = ToLower(Words[i-1])
			Words[i] = ""
		} else if Words[i] == "(bin)" {
			Words[i-1] = BinConv((Words[i-1]))
			Words[i] = ""
		} else if Words[i] == "(hex)" {
			Words[i-1] = HexConv((Words[i-1]))
			Words[i] = ""
		}
	}
	fmt.Println(Words)
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
