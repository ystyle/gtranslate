package gtranslate

import (
	"fmt"
	"testing"
	"time"
)

type testTable struct {
	inText   string
	langFrom string
	langTo   string
	outText  string
}

var testingTable = []testTable{
	{"Hello", "en", "es", "Hola"},
	{"Bye", "en", "es", "Adiós"},
	{"Hola", "es", "en", "Hello"},
	{"Adios", "es", "en", "Goodbye"},
	{"World", "en", "es", "Mundo"},
}

func TestTranslate(t *testing.T) {
	N := 5
	var totalDur time.Duration
	for i := 0; i < N; i++ {
		for _, ta := range testingTable {
			start := time.Now()
			translated, err := translate(ta.inText, ta.langFrom, ta.langTo, true, 5, time.Second)
			if err != nil {
				t.Error(err.Error())
			}
			if len(translated) < 2 {
				t.Fail()
			}
			dur := time.Since(start)
			fmt.Print(".")
			totalDur += dur
			if translated != ta.outText {
				t.Error("translated text is not the expected", ta.outText, " != ", translated)
			}
		}
	}
	fmt.Println("\nMean time:", time.Duration(int(totalDur)/(len(testingTable)*N)))
}
