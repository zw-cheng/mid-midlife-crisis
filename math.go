package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var data [][]string
var birthdate []int
var age int

func birthdateParser(s string) ([]int, error) {
	var res []int
	if len(s) != 8 {
		return []int{}, errors.New(fmt.Sprintf("Invalid birthdate: not 8 digits. Your input %s", s))
	}

	month, err := strconv.Atoi(s[0:2])
	if (err != nil) || month > 12 || month < 1 {
		return []int{}, errors.New("Invalid birthdate: month.")
	}

	day, err := strconv.Atoi(s[2:4])
	if (err != nil) || day > 31 || day < 1 {
		return []int{}, errors.New("Invalid birthdate: day.")
	}

	year, err := strconv.Atoi(s[4:8])
	if (err != nil) || year > 2022 || year < 1900 {
		return []int{}, errors.New(fmt.Sprintf("Invalid birthdate: year. %d", year))
	}

	res = append(res, year)
	res = append(res, month)
	res = append(res, day)
	return res, err
}

func getLifes(age int) int {
	return 555
}

func readData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("???")
	}

	return records, nil
}

func calculateLife(age int) string {
	lifeExpect := data[age][3]
	message := fmt.Sprintf("You have %s more years to live.\n", lifeExpect)

	for i := age; i <= 100; i += 10 {
		roundUpAge := (i/10 + 1) * 10
		numberOfLifes1, _ := strconv.Atoi(data[age][2])
		numberOfLifes2, _ := strconv.Atoi(data[roundUpAge][2])
		percentage := float64(numberOfLifes2) / float64(numberOfLifes1) * 100
		message += fmt.Sprintf("You have a chance of %.2f%% to live to %d.\n", percentage, roundUpAge)
	}

	return message
}

func calculate(s string) (string, error) {
	parseresult, err := birthdateParser(s)
	if err != nil {
		return "invalid input", err
	}
	birthdate = parseresult

	age = time.Now().Year() - birthdate[0]

	readresult, err := readData("data.csv")
	if err != nil {
		log.Fatal("read data failed")
	}
	data = readresult

	message := calculateLife(age)

	return message, err
}
