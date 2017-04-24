package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"strconv"
	"math"
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

type sortedPoint struct {
	distance float64
	x        int
	y        int
	setNO    int
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
	// x=39, y=9
	offices[1].x = 39
	offices[1].y = 9
	offices[1].setNO = 1

	points[offices[1].x][offices[1].y] = "@"

	// office
	// x=39, y=39
	offices[2].x = 39
	offices[2].y = 39
	offices[2].setNO = 1

	points[offices[2].x][offices[2].y] = "@"

	fmt.Print("Police offices: \n")
	for i := 0; i < 3; i++ {
		fmt.Printf("(%d,%d)", offices[i].x, offices[i].y)
	}
	fmt.Println()
	fmt.Println()

	fmt.Print("Fist set (Team 1): \n")
	for i := 0; i < 30; i++ {
		fmt.Printf("(%d,%d)", setOne[i].x, setOne[i].y)
	}
	fmt.Println()
	fmt.Println()

	fmt.Print("Second set (Team 2): \n")
	for i := 0; i < 30; i++ {
		fmt.Printf("(%d,%d)", setTwo[i].x, setTwo[i].y)
	}
	fmt.Println()
	fmt.Println()

	for x := 49; x >= 0; x-- {
		fmt.Println(" - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - ")
		fmt.Print("|")
		for y := 0; y < 50; y++ {
			fmt.Print(points[y][x] + "|")
		}
		fmt.Print("\n")
	}
	fmt.Println(" - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - ")

	var getPointsOne = [3][30] sortedPoint{}
	var getPointsTwo = [3][30] sortedPoint{}

	for k := 0; k < 3; k++ {
		for i := 0; i < 30; i++ {
			absX := math.Abs(float64(offices[k].x - setOne[i].x))
			absY := math.Abs(float64(offices[k].y - setOne[i].y))
			distance := math.Sqrt(math.Pow(absX, 2) + math.Pow(absY, 2))
			//fmt.Println(distance)
			getPointsOne[k][i].distance = distance
			getPointsOne[k][i].setNO = setOne[i].setNO
			getPointsOne[k][i].x = setOne[i].x
			getPointsOne[k][i].y = setOne[i].y

		}
		for i := 0; i < 30; i++ {
			absX := math.Abs(float64(offices[k].x - setTwo[i].x))
			absY := math.Abs(float64(offices[k].y - setTwo[i].y))
			distance := math.Sqrt(math.Pow(absX, 2) + math.Pow(absY, 2))
			//fmt.Println(distance)
			getPointsTwo[k][i].distance = distance
			getPointsTwo[k][i].setNO = setTwo[i].setNO
			getPointsTwo[k][i].x = setTwo[i].x
			getPointsTwo[k][i].y = setTwo[i].y
		}
		//fmt.Println()
	}

	var getPoints = [3][60] sortedPoint{}

	for i := 0; i < 3; i++ {
		s := 0
		for j := 0; j < 30; j++ {
			getPoints[i][s].x = getPointsOne[i][j].x
			getPoints[i][s].y = getPointsOne[i][j].y
			getPoints[i][s].distance = getPointsOne[i][j].distance
			//fmt.Println(getPointsOne[i][j].distance)
			getPoints[i][s].setNO = getPointsOne[i][j].setNO
			s++
		}
		for j := 0; j < 30; j++ {
			getPoints[i][s].x = getPointsTwo[i][j].x
			getPoints[i][s].y = getPointsTwo[i][j].y
			getPoints[i][s].distance = getPointsTwo[i][j].distance
			//fmt.Println(getPointsTwo[i][j].distance)
			getPoints[i][s].setNO = getPointsTwo[i][j].setNO
			s++
		}
		//fmt.Println()
	}

	//for i := 0; i < 3; i++ {
	//	for j := 0; j < 60; j++ {
	//		fmt.Println(getPoints[i][j].x)
	//		//fmt.Println(getPoints[i][j].y)
	//		//fmt.Println(getPoints[i][j].distance)
	//	}
	//	fmt.Println("+++++++++++++++++++++++++++++++++")
	//}

	for s := 0; s < 3; s++ {
		for i := 1; i < len(getPoints[s]); i++ {
			for j := 0; j < len(getPoints[s])-i; j++ {
				if getPoints[s][j].distance > getPoints[s][j+1].distance {
					getPoints[s][j], getPoints[s][j+1] = getPoints[s][j+1], getPoints[s][j]
				}
			}
		}
	}

	fmt.Println("3 pointes close office 1: ")

	for i := 0; i < 3; i++ {
		fmt.Printf("(%d,%d) ", getPoints[0][i].x, getPoints[0][i].y)
	}

	fmt.Println()

	fmt.Println("3 pointes close office 2: ")

	for i := 0; i < 3; i++ {
		fmt.Printf("(%d,%d) ", getPoints[1][i].x, getPoints[1][i].y)
	}

	fmt.Println()

	fmt.Println("3 pointes close office 3: ")

	for i := 0; i < 3; i++ {
		fmt.Printf("(%d,%d) ", getPoints[2][i].x, getPoints[2][i].y)
	}

	fmt.Println()

}
