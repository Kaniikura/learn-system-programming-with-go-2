package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	file, err := os.Create("sample.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)

	records := [][]string{
		[]string{"名前", "年齢", "身長", "体重"},
		[]string{"Tanaka", "31", "190cm", "97kg"},
		[]string{"Suzuki", "46", "180cm", "79kg"},
		[]string{"Matsui", "45", "188cm", "95kg"},
	}

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			log.Fatal(err)
		}
	}

	writer.Flush()
}
