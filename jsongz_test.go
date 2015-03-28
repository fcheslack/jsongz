package jsongz

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

var _ = log.Print

type ItemMeta struct {
	ParsedDate string
}

type Creator struct {
	Name      string
	FirstName string
	LastName  string
}
type ItemData struct {
	Title    string
	Creators []Creator
}

type Item struct {
	Meta ItemMeta
	Data ItemData
}

var item1 = Item{
	Meta: ItemMeta{
		ParsedDate: "2012-09-05",
	},
	Data: ItemData{
		Title: "Water Not Actually Wet",
		Creators: []Creator{
			Creator{
				Name: "Faolan M Cheslack-Postava",
			},
			Creator{
				FirstName: "Sean",
				LastName:  "Takats",
			},
			Creator{
				FirstName: "Dan",
				LastName:  "Stillman",
			},
		},
	},
}
var item2 = Item{
	Meta: ItemMeta{
		ParsedDate: "2012-09-04",
	},
	Data: ItemData{
		Title: "Water is Wet",
		Creators: []Creator{
			Creator{
				FirstName: "Faolan",
				LastName:  "Cheslack-Postava",
			},
			Creator{
				Name: "Sean Takats",
			},
			Creator{
				FirstName: "Dan",
				LastName:  "Stillman",
			},
		},
	},
}
var item3 = Item{
	Meta: ItemMeta{
		ParsedDate: "2012-09-04",
	},
	Data: ItemData{
		Title: "Water is Wet",
		Creators: []Creator{
			Creator{
				FirstName: "Faolan",
				LastName:  "Cheslack-Postava",
			},
		},
	},
}
var item4 = Item{
	Meta: ItemMeta{
		ParsedDate: "2010-09-04",
	},
	Data: ItemData{
		Title: "Water is Wet",
		Creators: []Creator{
			Creator{
				FirstName: "Dan S",
				LastName:  "Stillman",
			},
			Creator{
				FirstName: "Faolan",
				LastName:  "Cheslack-Postava",
			},
			Creator{
				FirstName: "Sean",
				LastName:  "Takats",
			},
		},
	},
}
var item5 = Item{
	Meta: ItemMeta{
		ParsedDate: "2010-09-04",
	},
	Data: ItemData{
		Title: "purple",
		Creators: []Creator{
			Creator{
				FirstName: "Faolan",
				LastName:  "Cheslack-Postava",
			},
		},
	},
}
var item6 = Item{
	Meta: ItemMeta{
		ParsedDate: "2010-09-04",
	},
	Data: ItemData{
		Title: "Water is Not Wet",
		Creators: []Creator{
			Creator{
				FirstName: "Faolan",
				LastName:  "Cheslack-Postava",
			},
			Creator{
				FirstName: "Sean P",
				LastName:  "Takats",
			},
			Creator{
				FirstName: "Dan",
				LastName:  "Stillman",
			},
		},
	},
}

var filename = fmt.Sprintf("%s/jsongz_test_%d.gz", os.TempDir(), time.Now().Unix())

func TestGZJson(t *testing.T) {
	writeData := []*Item{&item1, &item2, &item3, &item4, &item5, &item6}

	gzwriter, err := NewWriter(filename)
	if err != nil {
		t.Fatal(err)
	}

	err = gzwriter.Encode(writeData)
	if err != nil {
		t.Fatal(err)
	}

	readData := make([]*Item, 0)

	gzreader, err := NewReader(filename)
	if err != nil {
		t.Fatal(err)
	}
	err = gzreader.Decode(&readData)
	if err != nil {
		t.Fatal(err)
	}

	if len(readData) != len(writeData) {
		t.Error("Unequal lengths of written/read data")
	}

	for i := 0; i < len(writeData); i++ {
		if readData[i].Data.Title != writeData[i].Data.Title {
			t.Error("non-matching titles")
		}
	}

	err = WriteFile(filename, writeData)
	if err != nil {
		t.Error(err)
	}

	readData2 := make([]*Item, 0)
	ReadFile(filename, &readData2)
	if len(readData2) != len(writeData) {
		t.Error("Unequal lengths of written/read data")
	}

	for i := 0; i < len(writeData); i++ {
		if readData2[i].Data.Title != writeData[i].Data.Title {
			t.Error("non-matching titles")
		}
	}
}
