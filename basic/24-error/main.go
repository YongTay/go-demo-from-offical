package main

import "fmt"

func f1() (int, error) {
	return -1, fmt.Errorf("some error")
}

type myError struct {
	arg interface{}
}

func (e *myError) Error() string {
	return fmt.Sprintf("this is a custom error, %v", e.arg)
}

func f2(n int) (int, error) {
	if n == 10 {
		return -1, &myError{arg: n}
	}

	return n, nil
}

func main() {
	_, err := f1()
	fmt.Println(err)

	v, e := f2(10)
	fmt.Println(v, e)
}
