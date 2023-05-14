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

func simulate(label string, sorting string, an int, fashion string) {
	algorithm := AlgorithmPerfomance{name: label}
	for i := 0; i < an; i++ {
		length := int(math.Pow(5, float64(i)))
		v := create_list(length, fashion)
		time_start := time.Now()
		switch sorting {
		case "selection":
			helper.SelectionSortIP(v)
		case "bubble":
			helper.BubbleSort(v)
		case "insertion":
			helper.InsertionSortIPV5(v)
		case "merge":
			helper.MergeSort(v)
		case "quick":
			helper.QuickSort(v, 0, len(v)-1)
		case "quick-random":
			helper.QuickSortRandomPivot(v, 0, len(v)-1)
		case "counting":
			helper.CountingSort(v)
		default:
			fmt.Println("I don't know this method")
			return
		}
		time_elapsed := time.Since(time_start)
		save_statistics(&algorithm, time_elapsed.String(), length)
	}
	display_statistics(&algorithm)
}
