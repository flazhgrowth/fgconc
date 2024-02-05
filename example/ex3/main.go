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
	var res1 int
	var res2 string
	var res3 string

	wg := lib.New()
	wg.Go(lib.FgHandler(func() error {
		data, err := fun1(10)
		if err != nil {
			return err
		}
		res1 = data

		return nil
	}))
	wg.Go(lib.FgHandler(func() error {
		data, err := fun2(5)
		if err != nil {
			return err
		}
		res2 = data

		return nil
	}))
	wg.Go(lib.FgHandler(func() error {
		data, err := fun2(7)
		if err != nil {
			return err
		}
		res3 = data

		return nil
	}))
	if err := wg.WhenDone(); err != nil {
		panic(err)
	}

	fmt.Println("res1: ", res1)
	fmt.Println("res2: ", res2)
	fmt.Println("res3: ", res3)
}
