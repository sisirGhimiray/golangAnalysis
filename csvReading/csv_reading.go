package csvreading

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ExecuteReading(fileName string) []string {
	data := make([]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.LazyQuotes = true

	for {

		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			break
		}

		d := strings.Join(record, "\n")
		e := strings.Split(d, ",")
		if len(e) >= 2 {

			f := strings.Replace(e[1], "\"", "", 2)
			if f != "" {

				data = append(data, f)
			}
		}

	}

	return data
}
