package main

import "fmt"

func main() {
	// x := 0

	// for x < 5 {
	// 	fmt.Println("value of x is:", x);
	// 	x++;
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("i:", i);
	// }

	names := []string{"yoshi", "mario", "peach", "bowser"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Printf("names[%d]: %s\n", i, names[i]);
	// }

	// for index, value := range names {
	// 	fmt.Printf("names[%d]: %s\n", index, value);
	// }

	for _, value := range names {
		fmt.Println(value);
		value = "new string"
	}

	fmt.Println(names)
}
