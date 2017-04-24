package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"strconv"
)

/*
Assume we are in a city, and the city has three police offices,
and some police man walking everywhere.
Now, the police offices need some of police mans back to offices because they need a important meeting,
So, the program will help them to find k police man what is short distance to back to office.
 */

type kNNPoint struct {
	x     int
	y     int
	setNO int
}

func main() {

	setOneFileName := "setOne.csv"
	setTwoFileName := "setTwo.csv"

	var points = [50][50]string{}

	// first the city is empty
	var i = 0
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			points[x][y] = " "
		}
	}

	var setOne = [30]kNNPoint{}

	// read set one
	file, err := os.Open(setOneFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	i = 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		//fmt.Printf("%T\n", record[0])
		number, _ := strconv.Atoi(record[0])
		setOne[i].x = number
		numberTwo, _ := strconv.Atoi(record[1])
		setOne[i].y = numberTwo
		setOne[i].setNO = 2
		i++
	}

	var setTwo = [30]kNNPoint{}

	// read set one
	fileTwo, err := os.Open(setTwoFileName)
	if err != nil {
		panic(err)
	}
	defer fileTwo.Close()
	readerTwo := csv.NewReader(fileTwo)
	i = 0
	for {
		record, err := readerTwo.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		//fmt.Printf("%T\n", record[0])
		number, _ := strconv.Atoi(record[0])
		setTwo[i].x = number
		numberTwo, _ := strconv.Atoi(record[1])
		setTwo[i].y = numberTwo
		setTwo[i].setNO = 3
		i++
	}

	for i := 0; i < 30; i++ {
		x := setOne[i].x
		y := setOne[i].y
		points[x][y] = "X"
	}

	for i := 0; i < 30; i++ {
		x := setTwo[i].x
		y := setTwo[i].y
		points[x][y] = "O"
	}

	// three points are police offices
	var offices = [3]kNNPoint{}

	// office
	// x=9, y=9
	offices[0].x = 9
	offices[0].y = 9
	offices[0].setNO = 1

	points[offices[0].x][offices[0].y] = "@"

	// office
	// x=9, y=19
	offices[1].x = 9
	offices[1].y = 19
	offices[1].setNO = 1

	points[offices[1].x][offices[1].y] = "@"

	// office
	// x=19, y=29
	offices[2].x = 19
	offices[2].y = 29
	offices[2].setNO = 1

	points[offices[2].x][offices[2].y] = "@"

	for x := 49; x >= 0; x-- {
		fmt.Println(" - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - ")
		fmt.Print("|")
		for y := 0; y < 50; y++ {
			fmt.Print(points[y][x] + "|")
		}
		fmt.Print("\n")
	}
	fmt.Println(" - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - ")
}
