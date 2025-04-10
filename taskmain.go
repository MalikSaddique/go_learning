package main

// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// //Separate Functions

// func WordCount(str string, ch chan int) {
// 	word := 0
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == ' ' {
// 			word++
// 		}
// 	}

// 	ch <- word
// }

// func SentenceCount(str string, ch chan int) {
// 	sentence := 0
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == '.' {
// 			sentence++
// 		}
// 	}

// 	ch <- sentence

// }

// func SpecialCharactersCount(str string, ch chan int) {
// 	characters := 0
// 	for i := 0; i < len(str); i++ {
// 		switch str[i] {
// 		case '@', '{', '}', '[', ']', '*', '&', '$', '+', '-', '^', '(', ')', '#', '%', '`':
// 			characters++
// 		}
// 	}

// 	ch <- characters

// }

// func SpacesCount(str string, ch chan int) {
// 	spaces := 0
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == ' ' {
// 			spaces++
// 		}
// 	}

// 	ch <- spaces

// }

// func DigitsCount(str string, ch chan int) {
// 	digits := 0
// 	for i := 0; i < len(str); i++ {
// 		if str[i] >= '0' && str[i] <= '9' {
// 			digits++
// 		}
// 	}

// 	ch <- digits

// }

// func PunctuationCount(str string, ch chan int) {
// 	symbols := 0
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == ':' || str[i] == ';' || str[i] == ',' || str[i] == '!' || str[i] == '.' || str[i] == '"' || str[i] == '?' || str[i] == '/' {
// 			symbols++
// 		}
// 	}

// 	ch <- symbols

// }

// func LinesCount(str string, ch chan int) {
// 	lines := 0
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == '\n' {
// 			lines++
// 		}
// 	}

// 	ch <- lines

// }

// func VowelsCount(str string, ch chan int) {
// 	vowels := 0
// 	for i := 0; i < len(str); i++ {
// 		switch str[i] {
// 		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
// 			vowels++
// 		}
// 	}

// 	ch <- vowels

// }

// func ConsonantsCount(str string, ch chan int) {
// 	consonants := 0
// 	for i := 0; i < len(str); i++ {
// 		if str[i] != 'a' && str[i] != 'e' && str[i] != 'i' && str[i] != 'o' && str[i] != 'u' && str[i] != 'A' && str[i] != 'E' && str[i] != 'I' && str[i] != 'O' && str[i] != 'U' {
// 			consonants++
// 		}
// 	}
// 	ch <- consonants
// }

// func ParaCount(str string, ch chan int) {
// 	para := 0
// 	newlineCount := 0

// 	for i := 0; i < len(str); i++ {
// 		if str[i] == '\n' {
// 			newlineCount++
// 		} else if newlineCount >= 2 {
// 			para++
// 			newlineCount = 0
// 		}
// 	}
// 	if newlineCount < 2 {
// 		para++
// 	}
// 	ch <- para
// }

// //Combine Function

// func CombineFunction(str string) (int, int, int, int, int, int, int, int, int) {

// 	words := 0
// 	vowels := 0
// 	digits := 0
// 	SpecialChar := 0
// 	lines := 0
// 	spaces := 0
// 	sentences := 0
// 	Punctuation := 0
// 	consonants := 0

// 	for i := 0; i < len(str); i++ {
// 		if str[i] == ' ' {
// 			spaces++
// 		}
// 		if str[i] == ' ' {
// 			words++
// 		} else if str[i] == '.' {
// 			sentences++
// 		}
// 		switch str[i] {
// 		case '@', '{', '}', '[', ']', '*', '&', '$', '+', '-', '^', '(', ')', '#', '%', '`':
// 			SpecialChar++
// 		}
// 		if str[i] >= '0' && str[i] <= '9' {
// 			digits++
// 		} else if str[i] == ':' || str[i] == ';' || str[i] == ',' || str[i] == '!' || str[i] == '.' || str[i] == '"' || str[i] == '?' || str[i] == '/' {
// 			Punctuation++
// 		} else if str[i] == '\n' {
// 			lines++
// 		}
// 		switch str[i] {
// 		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
// 			vowels++
// 		}
// 		if str[i] != 'a' && str[i] != 'e' && str[i] != 'i' && str[i] != 'o' && str[i] != 'u' && str[i] != 'A' && str[i] != 'E' && str[i] != 'I' && str[i] != 'O' && str[i] != 'U' {
// 			consonants++
// 		}

// 	}

// 	return words, digits, SpecialChar, lines, spaces, sentences, Punctuation, consonants, vowels

// }

// func main() {
// 	start := time.Now()
// 	fmt.Println("Reading file")
// 	file_name := "Dummy_text.txt"
// 	data, err := os.ReadFile(file_name)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("\t\t\t--------Separate functions-------- ")
// 	fmt.Println("File Length :", len(data))
// 	str := string(data)
// 	wordCountChannel := make(chan int)
// 	go WordCount(str, wordCountChannel)

// 	sentenceCountChannel := make(chan int)
// 	go SentenceCount(str, sentenceCountChannel)

// 	digitsCountChannel := make(chan int)
// 	go DigitsCount(str, digitsCountChannel)

// 	puncCountChannel := make(chan int)
// 	go PunctuationCount(str, puncCountChannel)

// 	linesCountChannel := make(chan int)
// 	go LinesCount(str, linesCountChannel)

// 	vowelsCountChannel := make(chan int)
// 	go VowelsCount(str, vowelsCountChannel)

// 	consonantsCountChannel := make(chan int)
// 	go ConsonantsCount(str, consonantsCountChannel)

// 	spacesCountChannel := make(chan int)
// 	go SpacesCount(str, spacesCountChannel)

// 	characterCountChannel := make(chan int)
// 	go SpecialCharactersCount(str, characterCountChannel)

// 	ParagraphsCountChannel := make(chan int)
// 	go ParaCount(str, ParagraphsCountChannel)

// 	totalWords := <-wordCountChannel
// 	totalSentences := <-sentenceCountChannel
// 	totalDigits := <-digitsCountChannel
// 	totalPunctuation := <-puncCountChannel
// 	totalLines := <-linesCountChannel
// 	totalVowels := <-vowelsCountChannel
// 	totalConsonants := <-consonantsCountChannel
// 	totalSpaces := <-spacesCountChannel
// 	totalCharacters := <-characterCountChannel
// 	totalParagraphs := <-ParagraphsCountChannel

// 	fmt.Println("Words are : ", totalWords)
// 	fmt.Println("Sentences are : ", totalSentences)
// 	fmt.Println("Digits are : ", totalDigits)
// 	fmt.Println("Punctuation are : ", totalPunctuation)
// 	fmt.Println("Lines are : ", totalLines)
// 	fmt.Println("Vowels are : ", totalVowels)
// 	fmt.Println("Consonants are : ", totalConsonants)
// 	fmt.Println("Spaces are : ", totalSpaces)
// 	fmt.Println("Special Characters are : ", totalCharacters)
// 	fmt.Println("Paragraphs are : ", totalParagraphs)

// 	elapse := time.Since(start)
// 	fmt.Printf("The total time it takes is : %s", elapse)

// 	str1 := string(data)
// 	start1 := time.Now()

// 	words, digits, specChar, lines, spaces, sentences, punctuation, consonants, vowels := CombineFunction(str1)
// 	fmt.Println("\n\t\t\t--------Combine functions-------- ")
// 	fmt.Println("Words are: ", words, "\nSpecial Characters are: ", specChar, "\nLines are: ", lines, "\nSpaces are: ", spaces, "\nSentences ", sentences, "\nPunctuation are: ", punctuation, "\nConsonants are: ", consonants, "\nVowels : ", vowels, "\nDigits are : ", digits)
// 	elapse1 := time.Since(start1)
// 	fmt.Printf("The total time it takes is : %s", elapse1)
// }
