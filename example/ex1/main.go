package main

import (
	"fmt"
	"time"

	"github.com/flazhgrowth/fgconc/lib"
)

func fun1(n int) (int, error) {
	for i := 0; i < n; i += 1 {
		time.Sleep(time.Second)
	}

	return 10, nil
}

func fun2(n int) (string, error) {
	for i := 0; i < n; i += 1 {
		time.Sleep(time.Second)
	}
	return "Shani Indira", nil
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan string)

	wg := lib.New()
	wg.Go(lib.FgHandler(func() error {
		defer close(ch1)
		data, err := fun1(10)
		if err != nil {
			return err
		}
		ch1 <- data

		return nil
	}))
	wg.Go(lib.FgHandler(func() error {
		defer close(ch2)
		data, err := fun2(5)
		if err != nil {
			return err
		}
		ch2 <- data

		return nil
	}))
	wg.Go(lib.FgHandler(func() error {
		defer close(ch3)
		data, err := fun2(7)
		if err != nil {
			return err
		}
		ch3 <- data

		return nil
	}))
	if err := wg.WhenDone(); err != nil {
		panic(err)
	}

	res1 := <-ch1
	res2 := <-ch2
	res3 := <-ch3

	fmt.Println("res1: ", res1)
	fmt.Println("res2: ", res2)
	fmt.Println("res3: ", res3)
}
