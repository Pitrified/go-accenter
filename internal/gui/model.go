package accenter

import (
	"fmt"

	persist "example.com/accenter/internal/persist"
	weightedrand "example.com/accenter/internal/weightedrand"
	wiki "example.com/accenter/pkg/wiki"
)

type guiModel struct {
	wr map[wiki.Word]wiki.WikiRecord
	iw map[wiki.Word]weightedrand.InfoWord
}

func newModel() *guiModel {
	// create the model
	m := &guiModel{}

	// load the records and the info
	m.wr, m.iw = persist.LoadDataset()

	// fmt.Printf("%+v\n", m.wr[0])
	fmt.Printf("Picked %+v\n", weightedrand.ExtractWord(m.iw))

	return m
}
