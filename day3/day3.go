package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func mulString(val string) int {
	val = strings.TrimLeft(val, "mul(")
	val = strings.TrimRight(val, ")")
	str_array := strings.Split(val, ",")
	x, _ := strconv.Atoi(str_array[0])
	y, _ := strconv.Atoi(str_array[1])
	var product int = x * y
	return product
}

func sumArray(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	pattern := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))`)
	var product_array []int = []int{}
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		allMatches := pattern.FindAllString(line, -1)
		for _, product_s := range allMatches {
			var product int = mulString(product_s)
			product_array = append(product_array, product)

		}
		count++

	}
	result := sumArray(product_array)
	fmt.Println("Part 1 ans: ", result)
	//part2
	file, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(file)

	do_mul_pattern := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))`)
	product_bin := []int{}
	p2_count := 0
	ans := 0
	var text string = ""
	for scanner.Scan() {
		line := scanner.Text()
		// accurate_result := dont_to_do_regex.ReplaceAllString(line, "*****")
		// tmp := strings.Split(accurate_result, "don't()")
		// accurate_result = tmp[0]
		text += line
		p2_count++

	}
	matches := do_mul_pattern.FindAllString(text, -1)
	enabled := true
	// Check if we are at the last line by looking ahead
	for _, product_s := range matches {
		if product_s == "do()" {
			enabled = true
			continue
		}
		if product_s == "don't()" {
			enabled = false
			continue
		}
		if enabled {
			var product int = mulString(product_s)
			ans += product
			fmt.Println(p2_count, " :product: ", ans)
			product_bin = append(product_bin, product)
		}
	}

	result = sumArray(product_bin)
	fmt.Println("Part 2 ans: ", result)

}
