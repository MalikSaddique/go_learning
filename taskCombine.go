package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func CombineFunctions(str string) (int, int, int, int, int, int, int, int, int) {

	words := 0
	vowels := 0
	digits := 0
	SpecialChar := 0
	lines := 0
	spaces := 0
	sentences := 0
	Punctuation := 0
	consonants := 0

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			spaces++
		}
		if str[i] == ' ' {
			words++
		} else if str[i] == '.' {
			sentences++
		}
		switch str[i] {
		case '@', '{', '}', '[', ']', '*', '&', '$', '+', '-', '^', '(', ')', '#', '%', '`':
			SpecialChar++
		}
		if str[i] >= '0' && str[i] <= '9' {
			digits++
		} else if str[i] == ':' || str[i] == ';' || str[i] == ',' || str[i] == '!' || str[i] == '.' || str[i] == '"' || str[i] == '?' || str[i] == '/' {
			Punctuation++
		} else if str[i] == '\n' {
			lines++
		}
		switch str[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowels++
		}
		if str[i] != 'a' && str[i] != 'e' && str[i] != 'i' && str[i] != 'o' && str[i] != 'u' && str[i] != 'A' && str[i] != 'E' && str[i] != 'I' && str[i] != 'O' && str[i] != 'U' {
			consonants++
		}

	}

	return words, digits, SpecialChar, lines, spaces, sentences, Punctuation, consonants, vowels

}

func main() {
	start := time.Now()
	fmt.Println("Reading file")
	file_name := "Dummy_text.txt"
	data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}
	fmt.Println("File Length :", len(data))

	str := string(data)

	totalLength := len(str)

	parts := 4

	chunkSize := totalLength / parts

	fmt.Println(73, chunkSize)

	chunks := make([]string, parts)
	// ch:=make(map[chan int] parts)

	// var wg sync.WaitGroup

	for i := 0; i < parts; i++ {
		startChunk := i * chunkSize     //0*10=0
		endChunk := (i + 1) * chunkSize //(0+1)*10=10

		if i == parts-1 {
			endChunk = totalLength
		}
		chunks[i] = str[startChunk:endChunk] // chunk[0]=str[0:10]

		// wg.Add(parts)
		// go func(index int, chunkData string) {
		words, digits, specChar, lines, spaces, sentences, punctuation, consonants, vowels := CombineFunctions(chunks[i])
		fmt.Println("\n\nChunks \n", i+1)
		fmt.Println("Words are: ", words, "\nSpecial Characters are: ", specChar, "\nLines are: ", lines, "\nSpaces are: ", spaces, "\nSentences ", sentences, "\nPunctuation are: ", punctuation, "\nConsonants are: ", consonants, "\nVowels : ", vowels, "\nDigits are : ", digits)
		// defer wg.Done()
		// ch<-CombineFunction(chunks[])
		// }(i, chunks[i])
		// wg.Wait()
	}

	fmt.Println("Chunk size is  divided into ", parts, " parts: ", chunkSize)

	elapse := time.Since(start)
	fmt.Printf("The total time it takes is : %s", elapse)
}
