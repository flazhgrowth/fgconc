package main

import (
	"fmt"
	"time"
)

func fun1(n int) (int, error) {
	for i := 0; i < n; i += 1 {
		if i == 6 {
			return 0, fmt.Errorf("error aja deh")
		}
		time.Sleep(time.Second)
	}

	return 10, nil
}

func fun2(n int) (string, error) {
	for i := 0; i < n; i += 1 {
		if i == 6 {
			return "", fmt.Errorf("error aja deh")
		}
		time.Sleep(time.Second)
	}
	return "Shani Indira", nil
}

func doSomething() error {
	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan string)
	errCh := make(chan error)

	go func() {
		defer close(ch1)
		data, err := fun1(7)
		if err != nil {
			errCh <- err
			return
		}
		errCh <- nil
		ch1 <- data
	}()
	go func() {
		defer close(ch2)
		data, err := fun2(2)
		if err != nil {
			errCh <- err
			return
		}
		errCh <- nil
		ch2 <- data
	}()
	go func() {
		defer close(ch3)
		data, err := fun2(6)
		if err != nil {
			errCh <- err
			return
		}
		errCh <- nil
		ch3 <- data
	}()

	n := 0
	for n < 3 {
		err := <-errCh
		if err != nil {
			return err
		}
		n += 1
	}

	close(errCh)
	res1 := <-ch1
	res2 := <-ch2
	res3 := <-ch3

	fmt.Println("res1: ", res1)
	fmt.Println("res2: ", res2)
	fmt.Println("res3: ", res3)

	return nil
}

func main() {
	for i := 0; i < 3; i += 1 {
		if err := doSomething(); err != nil {
			fmt.Println("err: ", err)
			time.Sleep(time.Second * 10)
			continue
		}
		fmt.Println("okeeee")
	}
}
