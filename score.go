package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: score <filename>")
	}

	fileName := os.Args[1]
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(absPath)

	readCSV(absPath)

}

func readCSV(fname string) {
	fp, err := os.Open(fname)
	csvReader := csv.NewReader(fp)

	if err != nil {
		panic(err)
	}

	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	dt := time.Now()

	dates := make(map[string]int)

	fmt.Println(dt.Format("2006-01-02"))

	for i := 0; i < 42; i++ {
		newTime := dt.AddDate(0, 0, i*-1)
		dates[newTime.Format("2006-01-02")] = 0
	}

	fmt.Println(dates)

	defer fp.Close()

}
