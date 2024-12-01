package main

import "fmt"

func main() {
	age := 45;

	fmt.Println(age <= 50);
	fmt.Println(age >= 50);
	fmt.Println(age == 45);
	fmt.Println(age != 50);

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is greater than 40")
	}

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index == 3 {
			fmt.Println("breaking at pos", index)
			break
		}

		fmt.Printf("the value at pos %d is %s\n", index, value)
	}

}
