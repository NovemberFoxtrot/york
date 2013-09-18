package yindex

import (
	"strings"
	"ynet"
	"ytext"
)

var theindex index

type indexData map[string]ytext.Documents

type index struct {
	data indexData
}

type QueryResult struct {
	Location string
	Sentence string
}

func Query(query string) []QueryResult {
	terms := strings.Split(query, ` `)

	results := make([]QueryResult, 0)

	for _, doc := range ytext.TheDocuments {
		for index := 0; index < len(doc.Sentences)-1; index++ {
			s := doc.Sentences[index]
			e := doc.Sentences[index+1]

			found := false

			for _, term := range terms {
				if found = strings.Contains(doc.Text[s:e], term); found != true {
					break
				}
			}

			if found == true {
				qr := QueryResult{Location: doc.Location, Sentence: doc.Text[s:e]}
				results = append(results, qr)
			}
		}
	}

	return results
}

func (i *index) update(w *ytext.Document) {
	for _, gram := range w.Grams {
		if i.data[gram] == nil {
			i.data[gram] = make(ytext.Documents, 0)
		}

		i.data[gram] = append(i.data[gram], w)
	}
}

func Add(website string) {
	var d ytext.Document
	d.Location = website
	d.Text = ynet.FetchUrl(website)
	d.CalcGrams()
	ytext.TheDocuments = append(ytext.TheDocuments, d)
	theindex.update(&d)
}

func IndexDataLen() int {
	return len(theindex.data)
}

func init() {
	theindex.data = make(indexData)
}
