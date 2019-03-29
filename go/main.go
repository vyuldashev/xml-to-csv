package main

import (
	"bufio"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"os"
)

var columns = []string{
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

func main() {
	f, _ := os.Open("../AS_ADDROBJ_20190324_a1a706ea-4ac7-43e7-b65b-68de81a57ddb.XML")
	defer f.Close()

	csvRawChannel := make(chan []string)
	go write(csvRawChannel)

	r := bufio.NewReader(f)
	d := xml.NewDecoder(r)

	for {
		token, err := d.Token()

		if token == nil || err != nil {
			break
		}

		go read(csvRawChannel, token)
	}

	fmt.Println("Waiting for all writes...")

	fmt.Println("Done")
}

func read(csvRawChannel chan []string, token xml.Token) {
	switch se := token.(type) {
	case xml.StartElement:
		if se.Name.Local != "Object" {
			return
		}

		data := make(map[string]string, len(se.Attr))
		for _, a := range se.Attr {
			data[a.Name.Local] = a.Value
		}

		item := make([]string, len(columns))

		for index, attr := range columns {
			if v, ok := data[attr]; ok {
				item[index] = v
			} else {
				item[index] = ""
			}
		}

		csvRawChannel <- item
	}
}

func write(csvRawChannel chan []string) {
	result, _ := os.Create("result.csv")
	defer result.Close()

	writer := csv.NewWriter(result)
	defer writer.Flush()

	for item := range csvRawChannel {
		writer.Write(item)
	}
}
