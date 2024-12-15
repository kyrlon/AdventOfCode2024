package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
	left_bin := []int{}
	right_bin := []int{}
	for scanner.Scan() {
		ex := scanner.Text()
		result := strings.Split(ex, " ")
		l, _ := strconv.Atoi(result[0])
		r, _ := strconv.Atoi(result[len(result)-1])
		left_bin = append(left_bin, l)
		right_bin = append(right_bin, r)

	}
	sort.Ints(left_bin)
	sort.Ints(right_bin)
	var ans = []int{}
	if len(left_bin) == len(right_bin) {
		for i := range left_bin {
			ans = append(ans, AbsInt(left_bin[i]-right_bin[i]))
		}
	}

	final_ans := sumArray(ans)
	fmt.Println("PArt1 ans: ", final_ans)

	//part2
	dict := make(map[int]int)
	for _, val_l := range left_bin {
		for _, val_R := range right_bin {

			if val_l == val_R {
				dict[val_l]++
			}
		}

	}

	final_container := []int{}

	for key, v := range dict {
		final_container = append(final_container, key*v)
	}

	final_ans = sumArray(final_container)
	println("part2 ans: ", final_ans)

}
