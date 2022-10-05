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
