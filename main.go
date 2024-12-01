package main

import "fmt"

func main() {
	
	// strings
	const nameOne string = "mario";
	var nameTwo = "luigi";
	var nameThree string;

	fmt.Println(nameOne, nameTwo, nameThree);

	nameTwo = "browser";
	nameThree = "peach";
	//nameOne = "hello";

	fmt.Println(nameOne, nameTwo, nameThree);

	namefour := "yoshi";
	fmt.Println(namefour);

	// ints
	var ageOne int = 20;
	var ageTwo = 30;
	var ageThree int;
	ageFour := 40;

	fmt.Println(ageOne, ageTwo, ageThree, ageFour);

	// bits & memory
	// var numOne int8 = 25;
	// var numTwo int8 = -128;
	// var numThree uint = 25;

	// floats
	var scoreOne float32 = -90.98;
	var scoreTwo float64 = 43282783.23840720238;
	scoreThree := 21332121.3211123;

	fmt.Println(scoreOne, scoreTwo, scoreThree);

}
