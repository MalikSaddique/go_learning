package models

import "gorm.io/gorm"

type Result struct {
	gorm.Model
	id          uint `json:"-"`
	Words       int  `json:"words"`
	Digits      int  `json:"digits"`
	SpecialChar int  `json:"special_char"`
	Lines       int  `json:"lines"`
	Spaces      int  `json:"spaces"`
	Sentences   int  `json:"sentences"`
	Punctuation int  `json:"punctuation"`
	Consonants  int  `json:"consonants"`
	Vowels      int  `json:"vowels"`
}
