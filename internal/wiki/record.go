package accenter

// Wikidictionary record.
//
// '#categories>#name',
// '#form_of>#word',
// '#pos',
// '#senses>#categories>#name',
// '#senses>#examples>#english',
// '#senses>#examples>#text',
// '#senses>#glosses>str',
// '#senses>#raw_glosses>str',
// '#senses>#tags>str',
// '#senses>#topics>str',
// '#word'
type WikiRecord struct {
	Category []Category `json:"categories"`
	FormOfs  []FormOf   `json:"form_of"`
	Pos      string     `json:"pos"`
	Senses   []Sense    `json:"senses"`
	Word     string     `json:"word"`
}

// Possibly not the same category as the one in Sense.
type Category struct {
	Name string `json:"name"`
}

type FormOf struct {
	Word string `json:"word"`
}

type Example struct {
	English string `json:"english"`
	Text    string `json:"text"`
}

type Sense struct {
	Categories []Category `json:"categories"`
	Examples   []Example  `json:"examples"`
	Glosses    []string   `json:"glosses"`
	RawGlosses []string   `json:"raw_glosses"`
	Tags       []string   `json:"tags"`
	Topics     []string   `json:"topics"`
}
