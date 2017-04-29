/*
 * generator.go
 *
 * CSc 372 Spring 2017 - Final Project
 *
 * Author: Yujia Lin, Dong Liang
 *
 * ---
 * The program will generate two random sets and each set includes 30 points.
 */

package main

import (
	"os"
	"math/rand"
	"encoding/csv"
	"strconv"
	"time"
)

type point struct {
	x int
	y int
}

func main() {

	// set random generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	setOne := [30]point{}
	setTwo := [30]point{}

	// first, generate a random point for set one
	var x = 0
	var y = 0
	for {
		x = r.Intn(50)
		y = r.Intn(50)
		if x != 9 && y != 9 {
			if x != 39 && y != 9 {
				if x != 39 && y != 39 {
					break
				}
			}
		}
	}

	// generate whole random points for set one
	setOne[0].x = x
	setOne[0].y = y

	for i := 1; i < 30; i++ {
		for {
			x = r.Intn(50)
			y = r.Intn(50)
			for j := 0; j < i; j++ {
				if setOne[j].x == x && setOne[j].y == y {
					continue
				}
			}
			if x != 9 && y != 9 {
				if x != 39 && y != 9 {
					if x != 39 && y != 39 {
						setOne[i].x = x
						setOne[i].y = y
						break
					}
				}
			}
		}
	}

	// second, generate a random point for set two
	for {
		x = r.Intn(50)
		y = r.Intn(50)
		for j := 0; j < 30; j++ {
			if setOne[j].x == x && setOne[j].y == y {
				continue
			}
		}
		if x != 9 && y != 9 {
			if x != 39 && y != 9 {
				if x != 39 && y != 39 {
					break
				}
			}
		}
	}

	setTwo[0].x = x
	setTwo[0].y = y

	// generate whole random points for set two
	for i := 1; i < 30; i++ {
		for {
			x = r.Intn(50)
			y = r.Intn(50)
			for j := 0; j < 30; j++ {
				if setOne[j].x == x && setOne[j].y == y {
					continue
				}
			}
			for j := 0; j < i; j++ {
				if setTwo[j].x == x && setTwo[j].y == y {
					continue
				}
			}
			if x != 9 && y != 9 {
				if x != 39 && y != 9 {
					if x != 39 && y != 39 {
						setTwo[i].x = x
						setTwo[i].y = y
						break
					}
				}
			}
		}
	}

	// write csv file that name is "setOne.csv"
	file, err := os.Create("setOne.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range setOne {
		var data = []string{strconv.Itoa(value.x), strconv.Itoa(value.y)}
		err := writer.Write(data)
		if err != nil {
			panic(err)
		}
	}

	// write csv file that name is "setTwo.csv"
	fileTwo, err := os.Create("setTwo.csv")
	if err != nil {
		panic(err)
	}
	defer fileTwo.Close()

	writerTwo := csv.NewWriter(fileTwo)
	defer writerTwo.Flush()

	for _, value := range setTwo {
		var data = []string{strconv.Itoa(value.x), strconv.Itoa(value.y)}
		err := writerTwo.Write(data)
		if err != nil {
			panic(err)
		}
	}

}
