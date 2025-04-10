package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"time"
// )

// // type ConditionFunction func(byte) bool

// // func CombineFunctions(str string, condition ConditionFunction) int {
// // 	words := 0
// // 	for i := 0; i < len(str); i++ {
// // 		if condition(str[i]) {
// // 			words++
// // 		}
// // 	}
// // 	return words
// // }

// func VowelsC(chunks []byte) int {
// 	vowels := 0
// 	for i := 0; i < len(chunks); i++ {
// 		switch chunks[i] {
// 		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
// 			vowels++
// 		}
// 	}

// 	return vowels

// }

// func LinesC(chunks []byte) int {
// 	lines := 0
// 	for i := 0; i < len(chunks); i++ {
// 		if chunks[i] == '\n' {
// 			lines++
// 		}
// 	}

// 	return lines

// }

// func SpacesC(chunks []byte) int {
// 	spaces := 0
// 	for i := 0; i < len(chunks); i++ {
// 		if chunks[i] == ' ' {
// 			spaces++
// 		}
// 	}

// 	return spaces

// }

// // func CountSentences(ch byte) bool {
// // 	if ch == '.' {
// // 		return true
// // 	}
// // 	return false
// // }

// // func CountWords(ch byte) bool {
// // 	if ch == ' ' {
// // 		return true
// // 	}
// // 	return false
// // }

// // func CountCharacters(ch byte) bool {
// // 	switch ch {
// // 	case '@', '{', '}', '[', ']', '*', '&', '$', '+', '-', '^', '(', ')', '#', '%', '`':
// // 		return true
// // 	}
// // 	return false
// // }

// // func CountSpaces(ch byte) bool {
// // 	if ch == ' ' {
// // 		return true
// // 	}
// // 	return false
// // }
// // func CountDigits(ch byte) bool {
// // 	if ch >= '0' && ch <= '9' {
// // 		return true
// // 	}
// // 	return false
// // }
// // func CountPunctuation(ch byte) bool {
// // 	if ch == ':' || ch == ';' || ch == ',' || ch == '!' || ch == '.' || ch == '"' || ch == '?' || ch == '/' {
// // 		return true
// // 	}
// // 	return false
// // }

// // func CountConsonants(ch byte) bool {
// // 	if ch != 'a' && ch != 'e' && ch != 'i' && ch != 'o' && ch != 'u' && ch != 'A' && ch != 'E' && ch != 'I' && ch != 'O' && ch != 'U' {
// // 		return true
// // 	}
// // 	return false
// // }

// // func CountLines(ch byte) bool {
// // 	if ch == '\n' {
// // 		return true
// // 	}
// // 	return false
// // }

// // func AnotherCombineFunction(str string, mode string, ) int {
// // 	words := 0
// // 	for i := 0; i < len(str); i++ {
// // 		switch mode {
// // 		case "vowels":
// // 			switch str[i] {
// // 			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
// // 				words++
// // 			}
// // 		case "digits":
// // 			if str[i] >= '0' && str[i] <= '9' {
// // 				words++
// // 			}
// // 		case "words":
// // 			if str[i] == ' ' {
// // 				words++
// // 			}
// // 		case "sentences":
// // 			if str[i] == '.' {
// // 				words++
// // 			}
// // 		case "characters":
// // 			switch str[i] {
// // 			case '@', '{', '}', '[', ']', '*', '&', '$', '+', '-', '^', '(', ')', '#', '%', '`':
// // 				words++
// // 			}
// // 		case "spaces":
// // 			if str[i] == ' ' {
// // 				words++
// // 			}
// // 		case "punc":
// // 			if str[i] == ':' || str[i] == ';' || str[i] == ',' || str[i] == '!' || str[i] == '.' || str[i] == '"' || str[i] == '?' || str[i] == '/' {
// // 				words++
// // 			}
// // 		case "lines":
// // 			if str[i] == '\n' {
// // 				words++
// // 			}
// // 		case "consonants":
// // 			if str[i] != 'a' && str[i] != 'e' && str[i] != 'i' && str[i] != 'o' && str[i] != 'u' && str[i] != 'A' && str[i] != 'E' && str[i] != 'I' && str[i] != 'O' && str[i] != 'U' {
// // 				words++
// // 			}

// // 		}

// // 	}
// // 	return words

// // }

// func main() {
// 	start := time.Now()
// 	fmt.Println("Reading file")
// 	file_name := "Dummy_text.txt"
// 	data, err := os.Open(file_name)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer data.Close()

// 	chunkSize := 1024
// 	buffer := make([]byte, chunkSize)

// 	// for {
// 	bytesRead, err := data.Read(buffer)
// 	if err != nil && err != io.EOF {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}
// 	startChunks := 1000
// 	endChunks := 5000
// 	chunkNo := 1

// 	if chunkNo >= startChunks && chunkNo <= endChunks {
// 		chunks := buffer[:bytesRead]
// 		IsVowels := VowelsC(chunks)

// 		isLines := LinesC(chunks)
// 		isSpaces := SpacesC(chunks)

// 		fmt.Println("Vowels are : ", IsVowels)
// 		fmt.Println("Lines are : ", isLines)
// 		fmt.Println("Words are : ", isSpaces)
// 	}

// 	// fmt.Printf("File content %s", data)
// 	// fmt.Println("File Length :", len(data))
// 	// str := string(data)

// 	// if chunkNo > endChunks {
// 	// 	break
// 	// }
// 	chunkNo++

// 	// combinechannel := make(chan int)
// 	// totalword := AnotherCombineFunction(str, "words", combinechannel)
// 	// totalword := combinechannel
// 	// fmt.Println("Words are : ", AnotherCombineFunction(str, "words"))

// 	// fmt.Println("Sentences are : ", AnotherCombineFunction(str, "sentences"))

// 	// fmt.Println("Digits are : ", AnotherCombineFunction(str, "digits"))

// 	// fmt.Println("Punctuation are : ", AnotherCombineFunction(str, "punc"))

// 	// fmt.Println("Lines are : ", AnotherCombineFunction(str, "lines"))

// 	// fmt.Println("Vowels are : ", AnotherCombineFunction(str, "vowels"))

// 	// fmt.Println("Consonants are : ", AnotherCombineFunction(str, "consonants"))

// 	// fm
// 	// t.Println("Spaces are : ", AnotherCombineFunction(str, "spaces"))

// 	// fmt.Println("Special Characters are : ", AnotherCombineFunction(str, "characters"))

// 	// fmt.Println("Paragraphs are : ", ParaCount(str))

// 	elapse := time.Since(start)
// 	fmt.Printf("The total time it takes is : %s", elapse)
// }
