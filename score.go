package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Runner struct {
	Fatigue      float64
	Form         float64
	Fitness      float64
	FitnessDates map[string]int
	FatigueDates map[string]int
}

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

	var a Runner
	getDates(&a)

	readCSV(absPath, &a)

	fmt.Printf("Fitness\n>\t%f\nFatigue\n>\t%f\nForm\n>\t%f\n", a.Fitness, a.Fatigue, a.Form)
}

func getDates(runner *Runner) {

	dt := time.Now()

	fitDates := make(map[string]int)

	// fmt.Println(dt.Format("2006-01-02"))

	for i := 0; i < 42; i++ {
		newTime := dt.AddDate(0, 0, i*-1)
		fitDates[newTime.Format("2006-01-02")] = 0
	}

	runner.FitnessDates = fitDates

	fatigueDates := make(map[string]int)

	for i := 0; i < 7; i++ {
		newTime := dt.AddDate(0, 0, i*-1)
		fatigueDates[newTime.Format("2006-01-02")] = 0
	}

	runner.FatigueDates = fatigueDates
}

func readCSV(fname string, runner *Runner) {
	fp, err := os.Open(fname)
	csvReader := csv.NewReader(fp)

	if _, err := csvReader.Read(); err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	for {
		rec, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		// TSS idx 22
		// date idx 6

		date := rec[5]
		s, err := strconv.ParseFloat(rec[22], 32)

		if err != nil {
			panic(err)
		}

		if _, ok := runner.FitnessDates[date]; ok {
			runner.Fitness += s
		}

		if _, ok := runner.FatigueDates[date]; ok {
			runner.Fatigue += s
		}
	}

	runner.Fatigue = runner.Fatigue / 7
	runner.Fitness = runner.Fitness / 42

	runner.Form = runner.Fitness - runner.Fatigue

	defer fp.Close()

}
