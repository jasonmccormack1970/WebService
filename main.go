package main

import (
	"WebService/models"
	"fmt"
)

func main() {

	u := models.User{
		ID:        2,
		FirstName: "Jason",
		LastName:  "McCormack",
	}

	fmt.Println(u)

	/* arr := [3]int{1, 2, 3}
	slice := arr[:]
	fmt.Println(arr, slice)

	arr[1] = 27
	fmt.Println(arr, slice)
	slice[2] = 40
	fmt.Println(arr, slice) */

	// slice := []int{1, 2, 3}
	// fmt.Println(slice)

	// slice = append(slice, 4)
	// fmt.Println(slice)

	// s2 := slice[0:]
	// s3 := slice[:2]
	// s4 := slice[1:2]
	// fmt.Println(s2, s3, s4)

	/* m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 373
	fmt.Println(m["foo"])

	delete(m, "foo")
	fmt.Println(m) */

	// type user struct {
	// 	ID        int
	// 	FirstName string
	// 	LastName  string
	// }
	// var u user
	// u.ID = 1
	// u.FirstName = "Bob"
	// u.LastName = "Smith"
	// fmt.Println(u)

	// u2 := user{ID: 1,
	// 	FirstName: "Bob",
	// 	LastName:  "Smith",
	// }

	// fmt.Println(u2)
	port, err := startWebServer(3000, 10)
	fmt.Println(port, err)
}

func startWebServer(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting Webserver...")
	fmt.Println("Webserver started")
	fmt.Println(port, numberOfRetries)
	//return errors.New("Something Went wrong")
	return port, nil

}
