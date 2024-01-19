package main

import (
	"errors"
	"fmt"
	"math"
)

type booking_struct struct {
	id           string
	booking_type string
	date         string
}

func main() {
	var x int
	x = 5
	fmt.Println("hello world")
	fmt.Println(x)
	if x > 6 {
		fmt.Println("More than 6")
	} else if x < 2 {
		fmt.Println("Bite me")
	} else {
		print("brudddaaaaaaaa")
	}
	a := [5]int{1, 2, 3, 4, 5}                            //fixed length numerical array
	b := [5]string{"one", "two", "three", "four", "five"} //fixed length string array
	fmt.Println(a)
	fmt.Println(b)
	c := []int{1, 2, 3, 4, 5} //variable length numerical array
	c = append(c, 12)         //add an item to the end of an array
	fmt.Println(c)
	//map is like a dictionary in python
	booking := make(map[string]string)
	booking["id"] = "1"
	booking["user"] = "Tzenin@email.com"
	booking["type"] = "insane ass whoopin"
	booking["square"] = "this is not supposed to be here"
	delete(booking, "square") //delete the accidental option
	fmt.Println(booking)
	//loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// iterate over an array
	for index, value := range c {
		fmt.Println("index:", index, "value:", value)
	}
	// iterate over a map
	for index, value := range booking {
		fmt.Println(index, ":", value)
	}
	sum(1, 2)
	fmt.Println(sqrt(2))
	//using the booking struct
	new_booking := booking_struct{id: "g269", booking_type: "consultation", date: "in 5 minutes"}
	fmt.Println(new_booking)
	fmt.Println(new_booking.id)
	//pointers
	i := 7
	inc2(i)
	fmt.Println(i, "(See it doesn't affect the code)")
	inc(&i)
	fmt.Println(i, "(now it affects the code)")
}

// function with arguments
func sum(x int, y int) int {
	return x + y
}

// function with error handling
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

// function to demonstrate how pointers allow you to affect variables in other functions
func inc(x *int) []int {
	*x++
	return nil
}
func inc2(x int) []int {
	x++
	return nil
}
