package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrInvalidNum   = errors.New("invalid num")
	ErrServiceError = errors.New("service error")
)

func ToString(num int) string {
	if num == 0 {
		return "ноль"
	}

	if num < 0 || num > 99999 {
		return "число вне диапазона"
	}

	units := []string{"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять",
		"десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать",
		"шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}

	tens := []string{"", "", "двадцать", "тридцать", "сорок", "пятьдесят",
		"шестьдесят", "семьдесят", "восемьдесят", "девяносто"}

	hundreds := []string{"", "сто", "двести", "триста", "четыреста", "пятьсот",
		"шестьсот", "семьсот", "восемьсот", "девятьсот"}

	result := ""

	// тысячи
	if num >= 1000 {
		thousands := num / 1000
		if thousands == 1 {
			result += "одна тысяча "
		} else if thousands == 2 {
			result += "две тысячи "
		} else {
			// от 3 до 19
			if thousands < 20 {
				result += units[thousands] + " "
			} else {
				// от 20 до 99
				result += tens[thousands/10] + " "
				if thousands%10 > 0 {
					result += units[thousands%10] + " "
				}
			}
			// определяем окончание для тысяч
			if thousands >= 5 && thousands <= 19 {
				result += "тысяч "
			} else if thousands%10 == 1 {
				result += "тысяча "
			} else if thousands%10 >= 2 && thousands%10 <= 4 {
				result += "тысячи "
			} else {
				result += "тысяч "
			}
		}

		num %= 1000
	}

	// сотни
	if num >= 100 {
		result += hundreds[num/100] + " "
		num %= 100
	}

	// десятки и единицы
	if num >= 20 {
		result += tens[num/10] + " "
		num %= 10
	}

	// единицы и числа 10-19
	if num > 0 {
		result += units[num]
	}

	return strings.TrimSpace(result)
}

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

	fmt.Printf("\nРезультат программы: %d (%s)", res, ToString(res))
}
