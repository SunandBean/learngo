package prev_2

import (
	"fmt"
	"time"
)

func prev_2() {
	// go sexyCount("nico")
	// go sexyCount("flynn")
	// time.Sleep(time.Second * 5)

	// c := make(chan bool)
	// people := [2]string{"nico", "flynn"}
	// for _, person := range people {
	// 	go isSexy(person, c)
	// }
	// // result := <-c
	// // fmt.Println(result)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// c := make(chan string)
	// people := [2]string{"nico", "flynn"}
	// for _, person := range people {
	// 	go isSexy(person, c)
	// }
	// // fmt.Println("Waiting for messages")
	// // fmt.Println("Received this message: ", <-c)
	// // fmt.Println("Received this message: ", <-c)

	// fmt.Println("Waiting for messages")
	// resultOne := <-c
	// resultTwo := <-c
	// fmt.Println("Received this message: ", resultOne)
	// fmt.Println("Received this message: ", resultTwo)

	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}

	fmt.Println("Waiting for messages")
	for i := 0; i < len(people); i++ {
		fmt.Println("Waiting for ", i)
		fmt.Println("Received this message: ", <-c)
	}
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

// func isSexy(person string, c chan bool) {
// 	time.Sleep(time.Second * 5)
// 	fmt.Println(person)
// 	c <- true
// }

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is sexy"
}
