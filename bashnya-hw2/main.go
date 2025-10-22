package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidNum   = errors.New("invalid num")
	ErrServiceError = errors.New("service error")
)

func Task2(num int) (int, error) {
	if num >= 12307 {
		return -1, fmt.Errorf("error: %v", ErrInvalidNum)
	}

	for num < 12307 {
		if num < 0 {
			num *= -1
		} else if num%7 == 0 {
			num *= 39
		} else if num%9 == 0 {
			num = num*13 + 1
			continue
		} else {
			num = (num + 2) * 3
		}

		if num%9 == 0 && num%13 == 0 {
			return -1, fmt.Errorf("error: %v", ErrServiceError)
		} else {
			num++
		}
	}

	return num, nil
}

func main() {
	r := bufio.NewReader(os.Stdin)

	var num int
	fmt.Fscan(r, &num)

	res, err := Task2(num)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("\nРезультат программы: %d", res)
}
