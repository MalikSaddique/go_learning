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

	fmt.Println("Chunck size: ", chunkSize)

	chunks := make([]string, parts)
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	ch4 := make(chan string, 1)

	var wg sync.WaitGroup

	for i := 0; i < parts; i++ {
		startChunk := i * chunkSize     //0*10=0
		endChunk := (i + 1) * chunkSize //(0+1)*10=10

		if i == parts-1 {
			endChunk = totalLength
		}
		chunks[i] = str[startChunk:endChunk] // chunk[0]=str[0:10]

		go func(i int, chunkData string) {
			wg.Add(1)
			defer wg.Done()
			words, digits, specChar, lines, spaces, sentences, punctuation, consonants, vowels := CombineFunctions(chunkData)
			fmt.Println("\n\nChunks \n", i+1)
			result := fmt.Sprintf("Words are: %d\nSpecial Characters are: %d\nLines are: %d\nSpaces are: %d\nSentences: %d\nPunctuation are: %d\nConsonants are: %d\nVowels: %d\nDigits are: %d",
				words, specChar, lines, spaces, sentences, punctuation, consonants, vowels, digits)
			ch1 <- result
			ch2 <- result
			ch3 <- result
			ch4 <- result

		}(i, chunks[i])

		go CombineFunctions(str)

		output1 := <-ch1
		output2 := <-ch2
		output3 := <-ch3
		output4 := <-ch4

		fmt.Println(output1)
		fmt.Println(output2)
		fmt.Println(output3)
		fmt.Println(output4)

		// go func() {

		// 	close(ch)
		// }()

		// for result := range ch {
		// 	fmt.Println(result)
		// }

	}
	wg.Wait()

	fmt.Println("Chunk size is  divided into ", parts, " parts: ", chunkSize)

	elapse := time.Since(start)
	fmt.Printf("The total time it takes is : %s", elapse)
}
