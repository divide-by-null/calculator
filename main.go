package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		result := ""
		expression, _ := reader.ReadString('\n')

		if !strings.Contains(expression, " ") {
			log.Fatal("bad expression")
		}

		data, message, err := validate(expression)
		if err != nil {
			log.Fatal(err)
		}

		switch message {
		case "Arab":
			result, err = calcArab(data)
		case "Roman":
			result, err = calcRoman(data)
		default:
			fmt.Println("Not implemented yet :(")
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", result)
	}
}

func validate(expression string) ([]string, string, error) {
	var message = ""

	var tokens = strings.Split(strings.TrimRight(expression, "\n"), " ")

	if len(tokens) != 3 {
		return tokens, message, errors.New("bad expression")
	}

	if isNum(tokens[0]) && isNum(tokens[2]) {
		message = "Arab"
	}

	if isRomanNum(tokens[0]) && isRomanNum(tokens[2]) {
		message = "Roman"
	}

	if message == "" {
		return tokens, message, errors.New("incorrect members of an expression")
	}

	if !(tokens[1] == "+" || tokens[1] == "-" || tokens[1] == "*" || tokens[1] == "/") {
		return tokens, message, errors.New("unknown operation")
	}

	return tokens, message, nil
}

func isNum(s string) bool {
	num, err := strconv.ParseInt(s, 10, 0)
	return (num > 0 && num <= 10) && err == nil
}

func isRomanNum(s string) bool {
	Roman := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for i := 0; i < len(Roman); i++ {
		if s == Roman[i] {
			return true
		}
	}
	return false
}

func add(s1 string, s2 string) (int64, error) {
	num1, err := strconv.ParseInt(s1, 10, 64)
	if err != nil {
		return 0, err
	}
	num2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		return 0, err
	}
	return num1 + num2, nil
}

func sub(s1 string, s2 string) (int64, error) {
	num1, err := strconv.ParseInt(s1, 10, 64)
	if err != nil {
		return 0, err
	}
	num2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		return 0, err
	}
	return num1 - num2, nil
}

func mul(s1 string, s2 string) (int64, error) {
	num1, err := strconv.ParseInt(s1, 10, 64)
	if err != nil {
		return 0, err
	}
	num2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		return 0, err
	}
	return num1 * num2, nil
}

func div(s1 string, s2 string) (int64, error) {
	num1, err := strconv.ParseInt(s1, 10, 64)
	if err != nil {
		return 0, err
	}
	num2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		return 0, err
	}
	if num2 == 0 {
		return 0, errors.New("zero division")
	}
	return num1 / num2, nil
}

func calcArab(data []string) (string, error) {
	var result int64 = 0
	var err error = nil
	switch data[1] {
	case "+":
		result, err = add(data[0], data[2])
	case "-":
		result, err = sub(data[0], data[2])
	case "*":
		result, err = mul(data[0], data[2])
	case "/":
		result, err = div(data[0], data[2])
	default:
		return "", errors.New("unknown operation")
	}
	// fmt.Printf("Result: %d\n", result)
	return strconv.FormatInt(result, 10), err
}

func calcRoman(data []string) (string, error) {
	var num int64 = 0
	result := ""

	num1, err := romanToInteger(data[0])

	if err != nil {
		return "", err
	}
	num2, err := romanToInteger(data[2])

	if err != nil {
		return "", err
	}

	switch data[1] {
	case "+":
		num, err = add(num1, num2)
	case "-":
		num, err = sub(num1, num2)
	case "*":
		num, err = mul(num1, num2)
	case "/":
		num, err = div(num1, num2)
	default:
		return "", errors.New("unknown operation")
	}

	result, err = integerToRoman(num)
	return result, err
}

func romanToInteger(s string) (string, error) {
	arab := ""
	tableRA := map[string]string{"I": "1", "II": "2", "III": "3", "IV": "4", "V": "5", "VI": "6", "VII": "7", "VIII": "8", "IX": "9", "X": "10"}
	if isNum(s) {
		return arab, errors.New("not a roman number")
	}
	arab = tableRA[s]
	return arab, nil
}

func integerToRoman(number int64) (string, error) {
	if number < 1 {
		return "", errors.New("roman number must be greater than zero")
	}
	conversions := []struct {
		value int64
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String(), nil
}
