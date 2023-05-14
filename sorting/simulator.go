package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"dca0208.com/sortings/helper"
)

type AlgorithmPerfomance struct {
	name         string
	elapsed_arr  []string
	num_elements []int
}

/*
* @brief Creates an array with specifications
* @param n int
* @param fasion string
 */

func create_list(n int, fashion string) []int {
	list := make([]int, n)
	switch fashion {
	case "asc":
		for i := 0; i < n; i++ {
			list[i] = i
		}
	case "desc":
		for i := 0; i < n; i++ {
			list[i] = n - i
		}
	case "scientific":
		min_value := math.Pow(10, 3)
		max_value := math.Pow(10, 6)
		for i := 0; i < n; i++ {
			list[i] = rand.Intn(int(max_value)-int(min_value)) + int(min_value)
		}
	default:
		for i := 0; i < n; i++ {
			list[i] = rand.Intn(100)
		}
	}
	return list
}

/*
* @brief Register statistics
* @param *AlgorithmPerfomance
 */

func save_statistics(s *AlgorithmPerfomance, elapsed string, length int) {
	s.elapsed_arr = append(s.elapsed_arr, elapsed)
	s.num_elements = append(s.num_elements, length)
}

/*
* @brief Displays statistics
* @param *AlgorithmPerfomance
 */

func display_statistics(s *AlgorithmPerfomance) {
	fmt.Printf("Algoritmo: %s \n", s.name)
	for i := 0; i < len(s.elapsed_arr); i++ {
		fmt.Printf("%d elementos \t\t\t %s \n", s.num_elements[i], s.elapsed_arr[i])
	}
	fmt.Printf("\n")
}

/*
* @brief Simulate a sorting algorithm given specifications
* @param algorithm
* @param an int
* @param fashion string
 */

type SortFunc func(v []int)

func simulate(label string, sorting string, an int, fashion string) {

	algorithm := AlgorithmPerfomance{name: label}
	var method SortFunc

	switch sorting {
	case "selection":
		method = helper.SelectionSortIP
	case "bubble":
		method = helper.BubbleSort
	case "insertion":
		method = helper.InsertionSortIPV5
	case "merge":
		method = helper.MergeSort
	case "quick":
		method = helper.QuickSort
	case "quick-random":
		method = helper.QuickSortRandom
	case "counting":
		method = helper.CountingSort
	default:
		fmt.Println("I don't know this method")
		return
	}

	for i := 0; i < an; i++ {
		length := int(math.Pow(3, float64(i)))
		v := create_list(length, fashion)
		time_start := time.Now()
		method(v)
		time_elapsed := time.Since(time_start)
		save_statistics(&algorithm, time_elapsed.String(), length)
	}

	display_statistics(&algorithm)
}
