package main

import "fmt"

func main() {

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("Counter : %d \n", i)
	// }

	// fruits := []string{"mango", "apple", "banana"}
	// for i, counter := range fruits {
	// 	fmt.Printf("%d : %d", i, counter)
	// }

	// slice - dynamic array
	arr := []int{2, 2, 3, 4, 5, 7, 9, 12, 24, 29}
	arr = append(arr, 35)
	target := 4
	i, j, err := twoSumWithMap(arr, target)
	if err == nil {
		fmt.Printf("result : (%d, %d) \n", i, j)
	} else {
		fmt.Println("Error : ", err)
	}
}

// classic twoSum problem with sorted array
func twoSum(arr []int, target int) (int, int, error) {
	n := len(arr)

	i := 0
	j := n - 1
	for i < j {
		sum := arr[i] + arr[j]
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			return i, j, nil
		}
	}

	return -1, -1, fmt.Errorf("Indices not found!")
}

func twoSumWithMap(arr []int, target int) (int, int, error) {
	memo := map[int]int{}

	for i, el := range arr {
		complement := target - el

		val, ok := memo[complement]
		if ok {
			return val, i, nil
		}

		memo[el] = i
	}

	return -1, -1, fmt.Errorf("indices not found!")
}
