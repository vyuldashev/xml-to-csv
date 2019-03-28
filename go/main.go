package main

import (
	"bufio"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"os"
	"sync"
)

var attrs = []string{
	"AOGUID",
	"FORMALNAME",
	"REGIONCODE",
	"AUTOCODE",
	"AREACODE",
	"CITYCODE",
	"CTARCODE",
	"PLACECODE",
	"PLANCODE",
	"STREETCODE",
	"EXTRCODE",
	"SEXTCODE",
	"OFFNAME",
	"POSTALCODE",
	"IFNSFL",
	"TERRIFNSFL",
	"IFNSUL",
	"TERRIFNSUL",
	"OKATO",
	"OKTMO",
	"UPDATEDATE",
	"SHORTNAME",
	"AOLEVEL",
	"PARENTGUID",
	"AOID",
	"PREVID",
	"NEXTID",
	"CODE",
	"PLAINCODE",
	"ACTSTATUS",
	"LIVESTATUS",
	"CENTSTATUS",
	"OPERSTATUS",
	"CURRSTATUS",
	"STARTDATE",
	"ENDDATE",
	"NORMDOC",
	"CADNUM",
	"DIVTYPE",
}

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	f, _ := os.Open("../AS_ADDROBJ_20190324_a1a706ea-4ac7-43e7-b65b-68de81a57ddb.XML")
	result, _ := os.Create("result.csv")

	defer f.Close()
	defer result.Close()

	writer := csv.NewWriter(result)
	defer writer.Flush()

	r := bufio.NewReader(f)
	d := xml.NewDecoder(r)

	for {
		t, err := d.Token()

		if t == nil || err != nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local != "Object" {
				continue
			}

			item := make([]string, 0)

			for _, a := range se.Attr {
				if contains(attrs, a.Name.Local) {
					item = append(item, a.Value)
				}
			}

			wg.Add(1)
			go write(writer, item)
		}
	}

	fmt.Println("Waiting for all writes...")

	wg.Wait()

	fmt.Println("Done")
}

func write(w *csv.Writer, item []string) {
	mu.Lock()
	defer mu.Unlock()

	w.Write(item)
	wg.Done()
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
