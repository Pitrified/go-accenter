package accenter

type WeiWikiRecord struct {
	WikiRecord WikiRecord
	Weight     int
}

func NewWeiWikiRecord(wr WikiRecord, weight int) WeiWikiRecord {
	return WeiWikiRecord{
		WikiRecord: wr,
		Weight:     weight,
	}
}

// we weigh the records using
// the number of error done on this word
// the frequency of a word, unnormalized
type InfoWord struct {
	Word      string `json:"w"`
	Errors    int    `json:"e"`
	Frequency int    `json:"f"`
}

// basically we are reinventing a database, the key is the word
