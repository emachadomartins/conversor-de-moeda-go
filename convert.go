package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type file struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		checkErr(fmt.Errorf("expected 2 arguments, got %d", len(args)))
	}

	v, err := strconv.ParseFloat(args[0], 64)
	checkErr(err)

	c := strings.ToUpper(args[1])

	basePath := filepath.Join("base.json")

	f, err := os.ReadFile(basePath)
	checkErr(err)

	var base file
	err = json.Unmarshal(f, &base)
	checkErr(err)

	if base.Base == "" {
		checkErr(fmt.Errorf("base currency is empty"))
	}
	if len(base.Rates) == 0 {
		checkErr(fmt.Errorf("rates are empty"))
	}

	r := 0.0

	for currency, rate := range base.Rates {
		if currency == c {
			r = rate
			break
		}
	}

	if r == 0.0 {
		checkErr(fmt.Errorf("currency %s not found in rates", c))
	}

	result := v * r

	fmt.Printf("%.2f\n", result)
}

// 8043837003
// 7999469003
// 7973191003
// 8037007003
// 7999468003
//
//
//
//
//
//
