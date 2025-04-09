package main

import (
	"fmt"
	"os"
	"time"
)

// type ConditionFunction func(byte) bool

// func CombineFunctions(str string, condition ConditionFunction) int {
// 	words := 0
// 	for i := 0; i < len(str); i++ {
// 		if condition(str[i]) {
// 			words++
// 		}
// 	}
// 	return words
// }

// func CountVowels(ch byte) bool {
// 	switch ch {
// 	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
// 		return true
// 	}
// 	return false
// }

// func CountSentences(ch byte) bool {
// 	if ch == '.' {
// 		return true
// 	}
// 	return false
// }

// func CountWords(ch byte) bool {
// 	if ch == ' ' {
// 		return true
// 	}
// 	return false
// }

// func CountCharacters(ch byte) bool {
// 	switch ch {
// 	case '@', '{', '}', '[', ']', '*', '&', '$', '+', '-', '^', '(', ')', '#', '%', '`':
// 		return true
// 	}
// 	return false
// }

// func CountSpaces(ch byte) bool {
// 	if ch == ' ' {
// 		return true
// 	}
// 	return false
// }
// func CountDigits(ch byte) bool {
// 	if ch >= '0' && ch <= '9' {
// 		return true
// 	}
// 	return false
// }
// func CountPunctuation(ch byte) bool {
// 	if ch == ':' || ch == ';' || ch == ',' || ch == '!' || ch == '.' || ch == '"' || ch == '?' || ch == '/' {
// 		return true
// 	}
// 	return false
// }

// func CountConsonants(ch byte) bool {
// 	if ch != 'a' && ch != 'e' && ch != 'i' && ch != 'o' && ch != 'u' && ch != 'A' && ch != 'E' && ch != 'I' && ch != 'O' && ch != 'U' {
// 		return true
// 	}
// 	return false
// }

// func CountLines(ch byte) bool {
// 	if ch == '\n' {
// 		return true
// 	}
// 	return false
// }

func AnotherCombineFunction(str string, mode string) int {
	words := 0
	for i := 0; i < len(str); i++ {
		switch mode {
		case "vowels":
			switch str[i] {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				words++
			}
		case "digits":
			if str[i] >= '0' && str[i] <= '9' {
				words++
			}
		case "words":
			if str[i] == ' ' {
				words++
			}
		case "sentences":
			if str[i] == '.' {
				words++
			}
		case "characters":
			switch str[i] {
			case '@', '{', '}', '[', ']', '*', '&', '$', '+', '-', '^', '(', ')', '#', '%', '`':
				words++
			}
		case "spaces":
			if str[i] == ' ' {
				words++
			}
		case "punc":
			if str[i] == ':' || str[i] == ';' || str[i] == ',' || str[i] == '!' || str[i] == '.' || str[i] == '"' || str[i] == '?' || str[i] == '/' {
				words++
			}
		case "lines":
			if str[i] == '\n' {
				words++
			}
		case "consonants":
			if str[i] != 'a' && str[i] != 'e' && str[i] != 'i' && str[i] != 'o' && str[i] != 'u' && str[i] != 'A' && str[i] != 'E' && str[i] != 'I' && str[i] != 'O' && str[i] != 'U' {
				words++
			}

		}

	}
	return words

}

func main() {
	start := time.Now()
	fmt.Println("Reading file")
	file_name := "Dummy_text.txt"
	data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("File content %s", data)
	fmt.Println("File Length :", len(data))

	str := string(data)
	// combinechannel := make(chan int)
	// totalword := AnotherCombineFunction(str, "words", combinechannel)
	// totalword := combinechannel
	fmt.Println("Words are : ", AnotherCombineFunction(str, "words"))

	fmt.Println("Sentences are : ", AnotherCombineFunction(str, "sentences"))

	fmt.Println("Digits are : ", AnotherCombineFunction(str, "digits"))

	fmt.Println("Punctuation are : ", AnotherCombineFunction(str, "punc"))

	fmt.Println("Lines are : ", AnotherCombineFunction(str, "lines"))

	fmt.Println("Vowels are : ", AnotherCombineFunction(str, "vowels"))

	fmt.Println("Consonants are : ", AnotherCombineFunction(str, "consonants"))

	fmt.Println("Spaces are : ", AnotherCombineFunction(str, "spaces"))

	fmt.Println("Special Characters are : ", AnotherCombineFunction(str, "characters"))

	fmt.Println("Paragraphs are : ", ParaCount(str))

	elapse := time.Since(start)
	fmt.Printf("The total time it takes is : %s", elapse)
}
