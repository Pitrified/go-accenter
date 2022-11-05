package accenter

import "unicode/utf8"

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
	Word     Word       `json:"word"`
}

// Possibly not the same category as the one in Sense.
type Category struct {
	Name string `json:"name"`
}

type FormOf struct {
	Word Word `json:"word"`
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

type Word string

func (w Word) Len() int {
	return utf8.RuneCountInString(string(w))
}

func (w Word) AppendRune(r rune) Word {
	// return w + (string(r))
	return w + (Word(r))
}

func (w Word) Prefix(prefixLen int) Word {
	if prefixLen < 0 || prefixLen > w.Len() {
		return w
	}
	return Word([]rune(w)[:prefixLen])
}

func (w Word) PrefixStr(prefixLen int) string {
	return string(w.Prefix(prefixLen))
}

// Get all the Glosses for this record.
//
// For each sense in the record,
// if there are RawGlosses, return those,
// else return the Glosses.
func (wr *WikiRecord) GetAllGlosses() [][]string {
	allGlosses := make([][]string, len(wr.Senses))
	for si, sense := range wr.Senses {
		// allGlosses[si] = make([]string, sense.numGlosses())
		allGlosses[si] = sense.GetGlosses()
	}
	return allGlosses
}

// Return the RawGlosses if available, else the Glosses.
func (s *Sense) GetGlosses() []string {
	if len(s.RawGlosses) > 0 {
		return s.RawGlosses
	} else {
		return s.Glosses
	}
}

// Get the number of RawGlosses, if any, else the number of Glosses.
func (s *Sense) NumGlosses() int {
	if len(s.RawGlosses) > 0 {
		return len(s.RawGlosses)
	} else {
		return len(s.Glosses)
	}
}
