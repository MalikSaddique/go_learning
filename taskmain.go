package main

import (
	"fmt"
	"os"
)

func WordCount(str string) int {
	word := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			word++
		}
	}

	return word
}

func SentenceCount(str string) int {
	sentence := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '.' {
			sentence++
		}
	}

	return sentence

}

func SpecialCharactersCount(str string) int {
	characters := 0
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '@', '{', '}', '[', ']', '*', '&', '$', '+', '-', '^', '(', ')', '#', '%', '`':
			characters++
		}
	}

	return characters

}

func SpacesCount(str string) int {
	spaces := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			spaces++
		}
	}

	return spaces

}

func DigitsCount(str string) int {
	digits := 0
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			digits++
		}
	}

	return digits

}

func PunctuationCount(str string) int {
	symbols := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ':' || str[i] == ';' || str[i] == ',' || str[i] == '!' || str[i] == '.' || str[i] == '"' || str[i] == '?' || str[i] == '/' {
			symbols++
		}
	}

	return symbols

}

func LinesCount(str string) int {
	lines := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			lines++
		}
	}

	return lines

}

func VowelsCount(str string) int {
	vowels := 0
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowels++
		}
	}

	return vowels

}

func ConsonantsCount(str string) int {
	consonants := 0
	for i := 0; i < len(str); i++ {
		if str[i] != 'a' && str[i] != 'e' && str[i] != 'i' && str[i] != 'o' && str[i] != 'u' && str[i] != 'A' && str[i] != 'E' && str[i] != 'I' && str[i] != 'O' && str[i] != 'U' {
			consonants++
		}
	}
	return consonants
}

func main() {

	fmt.Println("Reading file")
	file_name := "Dummy_text.txt"
	data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("File content %s", data)
	fmt.Println("File Length :", len(data))
	str := string(data)
	fmt.Println("Words are : ", WordCount(str))

	fmt.Println("Sentences are : ", SentenceCount(str))

	fmt.Println("Digits are : ", DigitsCount(str))

	fmt.Println("Punctuation are : ", PunctuationCount(str))

	fmt.Println("Lines are : ", LinesCount(str))

	fmt.Println("Vowels are : ", VowelsCount(str))

	fmt.Println("Consonants are : ", ConsonantsCount(str))

	fmt.Println("Spaces are : ", SpacesCount(str))

	fmt.Println("Special Characters are : ", SpecialCharactersCount(str))

}
