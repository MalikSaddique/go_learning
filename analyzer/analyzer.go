package analyzer

import (
	"io"
	"net/http"
)

type Result struct {
	UserID      int `json:"user_id"`
	Words       int `json:"words"`
	Digits      int `json:"digits"`
	SpecialChar int `json:"special_char"`
	Lines       int `json:"lines"`
	Spaces      int `json:"spaces"`
	Sentences   int `json:"sentences"`
	Punctuation int `json:"punctuation"`
	Consonants  int `json:"consonants"`
	Vowels      int `json:"vowels"`
}

func CombineFunctions(str string) Result {
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
		if str[i] != 'a' && str[i] != 'e' && str[i] != 'i' && str[i] != 'o' && str[i] != 'u' &&
			str[i] != 'A' && str[i] != 'E' && str[i] != 'I' && str[i] != 'O' && str[i] != 'U' {
			consonants++
		}
	}

	return Result{
		Words:       words,
		Digits:      digits,
		SpecialChar: SpecialChar,
		Lines:       lines,
		Spaces:      spaces,
		Sentences:   sentences,
		Punctuation: Punctuation,
		Consonants:  consonants,
		Vowels:      vowels,
	}
}

func AnalyzeFile(r *http.Request) (Result, error) {
	// Parse the multipart form (adjust the max size if needed)
	// err := r.ParseMultipartForm(10 << 20) // 10 MB
	// if err != nil {
	// 	return Result{}, err
	// }

	file, _, err := r.FormFile("file")
	if err != nil {
		return Result{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Result{}, err
	}

	return CombineFunctions(string(data)), nil
}
