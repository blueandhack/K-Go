package main

import "fmt"

/*
Assume we are in a city, and the city has three police offices,
and some police man walking everywhere.
Now, the police offices need all of police mans back to offices because they need a important meeting,
So, the program will help them to find what is short distance back to office.
 */

type kNNPoint struct {
	x     int
	y     int
	setNO int
}

func main() {

	setOneFileName := "setOne.csv"
	setTwoFileName := "setTwo.csv"

	var points = [2500]kNNPoint{}

	// first the city is empty
	var i = 0
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			points[i].x = x
			points[i].y = y
			points[i].setNO = 0
			i++
		}
	}

	// three points are police offices
	var offices = [3]kNNPoint{}

	offices[0].x = 9
	offices[0].y = 9
	offices[0].setNO = 1

	offices[0].x = 9
	offices[0].y = 19
	offices[0].setNO = 1

	offices[0].x = 19
	offices[0].y = 29
	offices[0].setNO = 1

	for i := 0; i < 2500; i++ {
		fmt.Printf("X: %d, Y: %d, Set NO: %d\n", points[i].x, points[i].y, points[i].setNO)
	}
}
