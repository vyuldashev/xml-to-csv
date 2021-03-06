package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

var columns = []string{
	"AOGUID=\"",
	"FORMALNAME=\"",
	"REGIONCODE=\"",
	"AUTOCODE=\"",
	"AREACODE=\"",
	"CITYCODE=\"",
	"CTARCODE=\"",
	"PLACECODE=\"",
	"PLANCODE=\"",
	"STREETCODE=\"",
	"EXTRCODE=\"",
	"SEXTCODE=\"",
	"OFFNAME=\"",
	"POSTALCODE=\"",
	"IFNSFL=\"",
	"TERRIFNSFL=\"",
	"IFNSUL=\"",
	"TERRIFNSUL=\"",
	"OKATO=\"",
	"OKTMO=\"",
	"UPDATEDATE=\"",
	"SHORTNAME=\"",
	"AOLEVEL=\"",
	"PARENTGUID=\"",
	"AOID=\"",
	"PREVID=\"",
	"NEXTID=\"",
	"CODE=\"",
	"PLAINCODE=\"",
	"ACTSTATUS=\"",
	"LIVESTATUS=\"",
	"CENTSTATUS=\"",
	"OPERSTATUS=\"",
	"CURRSTATUS=\"",
	"STARTDATE=\"",
	"ENDDATE=\"",
	"NORMDOC=\"",
	"CADNUM=\"",
	"DIVTYPE=\"",
}

type StringParser struct{}

func (p *StringParser) Parse(id int, in <-chan string, out chan<- string, done chan<- int) {
	for value := range in {
		var line strings.Builder

		for _, attribute := range columns {
			start := strings.Index(value, attribute)

			if start == -1 {
				_ = line.WriteByte(';')
				continue
			}

			start += len(attribute)

			end := strings.Index(value[start:], "\"") + start

			_, _ = line.WriteString(value[start:end])
			_ = line.WriteByte(';')
		}

		_ = line.WriteByte('\n')

		out <- line.String()
	}

	done <- id
}

func main() {
	parser := StringParser{}

	workers := 3
	scanned, parsed, done := make(chan string), make(chan string), make(chan int)

	for id := 0; id < workers; id++ {
		go parser.Parse(id, scanned, parsed, done)
	}

	go func() {
		f, err := os.Open("./files/AS_ADDROBJ_20190324_a1a706ea-4ac7-43e7-b65b-68de81a57ddb.XML")

		if err != nil {
			panic(err)
		}

		defer f.Close()

		startOfXml := []byte("<Object")
		endOfXml := []byte("/>")

		scanner := bufio.NewScanner(f)
		split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			start := bytes.Index(data, startOfXml)

			if start == -1 {
				return 0, nil, nil
			}

			end := bytes.Index(data[start:], endOfXml)

			if end == -1 {
				return 0, nil, nil
			}

			width := start + end + len(endOfXml)

			return width, data[start:width], nil
		}
		scanner.Split(split)

		for scanner.Scan() {
			scanned <- scanner.Text()
		}

		close(scanned)
	}()

	go func() {
		result, err := os.Create("result.csv")

		if err != nil {
			panic(err)
		}

		defer result.Close()

		w := bufio.NewWriterSize(result, 4096*5)

		for v := range parsed {
			_, _ = w.WriteString(v)
		}

		_ = w.Flush()

		close(done)
	}()

	waits := workers

	for range done {
		waits--
		if waits == 0 {
			close(parsed)
		}
	}

	fmt.Println("Done")
}
