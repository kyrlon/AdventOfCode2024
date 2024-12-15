package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertToIntArray(arry []string) []int {
	new_a := []int{}

	for i := 0; i < len(arry); i++ {
		x, _ := strconv.Atoi(arry[i])
		new_a = append(new_a, x)
	}
	return new_a
}

func countTrue(arry []bool) int {

	n := 0
	for _, val := range arry {
		if val {
			n++
		}

	}
	return n
}

func arrayHasZero(arry []int) bool {
	state := false
	for _, val := range arry {
		if val == 0 {
			state = true
		}
	}
	return state
}

func arrayAllPos(arry []int) bool {
	same := true

	for _, val := range arry {
		if val < 0 {
			same = false
			break
		}
	}
	return same
}

func arrayAllNeg(arry []int) bool {
	same := true
	for _, val := range arry {
		if val > 0 {
			same = false
			break
		}
	}
	return same

}
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func levels_differ_safe(arry []int) bool {
	safe := true

	for _, val := range arry {
		val = AbsInt(val)
		if val > 3 || val == 0 {
			safe = false
			break
		}

	}
	return safe
}

func calc_dist_per_n(arry []int) []int {
	tmp_bin := []int{}

	for i := 0; i < len(arry); i++ {
		//diff of each index
		//
		if i == len(arry)-1 {
			break
		}
		r := arry[i] - arry[i+1]
		tmp_bin = append(tmp_bin, r)
	}
	return tmp_bin
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	true_bin := []bool{}
	for scanner.Scan() {
		line := scanner.Text()
		line_array := strings.Split(line, " ")
		int_array := convertToIntArray(line_array)
		tmp_bin := []int{}

		for i := 0; i < len(int_array); i++ {
			//diff of each index
			//
			if i == len(int_array)-1 {
				break
			}
			r := int_array[i] - int_array[i+1]
			tmp_bin = append(tmp_bin, r)
		}
		if arrayAllNeg(tmp_bin) || arrayAllPos(tmp_bin) {
			if !arrayHasZero(tmp_bin) {

				safe := levels_differ_safe(tmp_bin)
				true_bin = append(true_bin, safe)
				continue
			}
			true_bin = append(true_bin, false)

		} else {
			true_bin = append(true_bin, false)
		}
	}
	result := countTrue(true_bin)
	fmt.Println("PArt 1 result: ", result)

	true_bin = []bool{}
	// part 2
	file, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line_array := strings.Split(line, " ")
		int_array := convertToIntArray(line_array)
		tmp_bin := []int{}

		// for i := 0; i < len(int_array); i++ {
		// 	//diff of each index
		// 	//
		// 	if i == len(int_array)-1 {
		// 		break
		// 	}
		// 	r := int_array[i] - int_array[i+1]
		// 	tmp_bin = append(tmp_bin, r)
		// }
		tmp_bin = calc_dist_per_n(int_array)
		if (arrayAllNeg(tmp_bin) || arrayAllPos(tmp_bin)) && levels_differ_safe(tmp_bin) {
			true_bin = append(true_bin, true)
		} else {
			// not all same inc/dec
			state := false
			for i := 0; i < len(int_array); i++ {

				new_array := []int{}
				for ii, val := range int_array {
					if i == ii {
						continue
					}
					new_array = append(new_array, val)
				}

				tmp_bin = calc_dist_per_n(new_array)
				if (arrayAllNeg(tmp_bin) || arrayAllPos(tmp_bin)) && levels_differ_safe(tmp_bin) {
					state = true
					break
				}

			}

			true_bin = append(true_bin, state)
		}
	}
	result2 := countTrue(true_bin)
	fmt.Println("PArt 2 result: ", result2)

}
