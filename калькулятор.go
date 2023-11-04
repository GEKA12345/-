package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	lst := strings.Split(text, " ")
	var err1 error
	if len(lst) < 3 {
		err1 = errors.New("строка не является математической операцией")
	} else if len(lst) > 3 {
		err1 = errors.New("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
	return lst, err1
}

func rome_to_int(x string) (int, error) {
	dct := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	res, ok := dct[x]
	var err3 error
	if !ok {
		err3 = errors.New("на вход принимаются числа от 1 до 10 включительно как в арабском так и римском исчислениях")
	}
	return res, err3

}

func parse(lst []string) ([3]int, bool, error) {
	var flag_is_rome bool
	res := [3]int{0, 0, 0}
	str := "+-/*"
	res[2] = strings.Index(str, lst[1])
	if res[2] < 0 {
		err1 := errors.New("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		return res, flag_is_rome, err1
	}
	var rome_1, rome_2 error
	res[0], rome_1 = strconv.Atoi(lst[0])
	res[1], rome_2 = strconv.Atoi(lst[2])
	var err2 error
	var err3_1, err3_2 error
	if rome_1 == nil && rome_2 != nil || rome_2 == nil && rome_1 != nil {
		err2 = errors.New("используются одновременно разные системы исчисления")
		return res, flag_is_rome, err2
	} else if rome_1 != nil && rome_2 != nil {
		flag_is_rome = true
		res[0], err3_1 = rome_to_int(lst[0])
		res[1], err3_2 = rome_to_int(lst[2])
		if err3_1 != nil {
			return res, flag_is_rome, err3_1
		} else if err3_2 != nil {
			return res, flag_is_rome, err3_2
		}
	} else if res[0] < 1 || res[0] > 10 || res[1] < 1 || res[1] > 10 {
		err3_1 = errors.New("на вход принимаются числа от 1 до 10 включительно как в арабском так и римском исчислениях")
		return res, flag_is_rome, err3_1
	}
	return res, flag_is_rome, nil
}

func calcul(lst [3]int) int {
	var res int
	switch lst[2] {
	case 0:
		res = lst[0] + lst[1]
	case 1:
		res = lst[0] - lst[1]
	case 2:
		res = lst[0] / lst[1]
	case 3:
		res = lst[0] * lst[1]
	}
	return res
}

func int_to_rome(x int) (string, error) {
	dct := map[int]string{
		0:   "",
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		20:  "XX",
		30:  "XXX",
		40:  "XL",
		50:  "L",
		60:  "LX",
		70:  "LXX",
		80:  "LXXX",
		90:  "XC",
		100: "C",
	}
	var err4 error
	var res string
	if x < 1 {
		err4 = errors.New("результат вычисления меньше 1 - в римской системе нет отрицательных чисел")
		return res, err4
	}
	hund := x / 100
	res = res + dct[hund*100]
	x = x - hund*100
	tens := x / 10
	res = res + dct[tens*10]
	x = x - tens*10
	res = res + dct[x]
	return res, err4
}

func main() {
	lst, err := input()
	if err != nil {
		fmt.Println(err)
	} else {
		lst_parsed, flag_is_rome, err := parse(lst)
		if err != nil {
			fmt.Println(err)
		} else {
			res := calcul(lst_parsed)
			if flag_is_rome {
				ans, err := int_to_rome(res)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(ans)
				}
			} else {
				fmt.Println(res)
			}
		}
	}
}
