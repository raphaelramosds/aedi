package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"dca0208.com/sortings/helper"
)

// Struct da simulação

type AlgorithmPerfomance struct {

	// Inner properties
	v       []int
	elapsed string

	// Fillable properties
	label        string
	sorting      string
	fashion      string
	num_elements int
}

func (algorithm *AlgorithmPerfomance) init() {
	algorithm.v = make([]int, algorithm.num_elements)
	switch algorithm.fashion {
	case "random":
		for i := 0; i < algorithm.num_elements; i++ {
			algorithm.v[i] = rand.Intn(100)
		}
	case "asc":
		for i := 0; i < algorithm.num_elements; i++ {
			algorithm.v[i] = i
		}
	case "desc":
		for i := 0; i < algorithm.num_elements; i++ {
			algorithm.v[i] = algorithm.num_elements - i
		}
	case "sci":
		lower := int(math.Pow10(7))
		upper := int(math.Pow10(8))

		for i := 0; i < algorithm.num_elements; i++ {
			algorithm.v[i] = rand.Intn(upper-lower) + lower
		}
	default:
		fmt.Println("Forma inválida")
	}
}

func (algorithm *AlgorithmPerfomance) sort() {

	var elapsed time.Duration

	switch algorithm.sorting {
	case "bubble":
		start := time.Now()
		helper.BubbleSort(algorithm.v)
		elapsed = time.Since(start)
	case "selection":
		start := time.Now()
		helper.SelectionSortIP(algorithm.v)
		elapsed = time.Since(start)
	case "insertion":
		start := time.Now()
		helper.InsertionSortIPV4(algorithm.v)
		elapsed = time.Since(start)
	case "quick":
		start := time.Now()
		helper.QuickSort(algorithm.v, 0, len(algorithm.v)-1)
		elapsed = time.Since(start)
	case "counting":
		start := time.Now()
		helper.CountingSort(algorithm.v)
		elapsed = time.Since(start)
	case "quick-random":
		start := time.Now()
		helper.QuickSortRandom(algorithm.v, 0, len(algorithm.v)-1)
		elapsed = time.Since(start)
	case "merge":
		start := time.Now()
		helper.MergeSort(algorithm.v)
		elapsed = time.Since(start)
	default:
		fmt.Println("Algoritmo inválido")
	}

	algorithm.elapsed = elapsed.String()
}

func (algorithm *AlgorithmPerfomance) Simulate() {
	algorithm.init()
	algorithm.sort()
	fmt.Printf("%s (%d elementos) %s", algorithm.label, algorithm.num_elements, algorithm.elapsed)
	fmt.Printf("\n")
}

/*
* Fato 1: Para vetores de tamanho pequeno, a performance da maioria dos algoritmos
* de ordenação não vai influenciar, independente da disposição dos elementos
 */

func first_experiment(n int) {

	// Conclusão: o três algoritmos parecem ter perfomance muito rápidas para vetores abaixo de 1000 elementos, independente da disposição

	// Random
	counting_random := AlgorithmPerfomance{label: "Counting Sort", sorting: "counting", num_elements: n, fashion: "random"}
	quick_random := AlgorithmPerfomance{label: "Quick Sort", sorting: "quick", num_elements: n, fashion: "random"}
	bubble_random := AlgorithmPerfomance{label: "Bubble Sort", sorting: "bubble", num_elements: n, fashion: "random"}

	counting_random.Simulate()
	quick_random.Simulate()
	bubble_random.Simulate()

	fmt.Println("")

	// ASC
	counting_asc := AlgorithmPerfomance{label: "Counting Sort (ASC)", sorting: "counting", num_elements: n, fashion: "asc"}
	quick_asc := AlgorithmPerfomance{label: "Quick Sort (ASC)", sorting: "quick", num_elements: n, fashion: "asc"}
	bubble_asc := AlgorithmPerfomance{label: "Bubble Sort (ASC)", sorting: "bubble", num_elements: n, fashion: "asc"}

	counting_asc.Simulate()
	quick_asc.Simulate()
	bubble_asc.Simulate()

	fmt.Println("")

	// DESC
	counting_desc := AlgorithmPerfomance{label: "Counting Sort (DESC)", sorting: "counting", num_elements: n, fashion: "desc"}
	quick_desc := AlgorithmPerfomance{label: "Quick Sort (DESC)", sorting: "quick", num_elements: n, fashion: "desc"}
	bubble_desc := AlgorithmPerfomance{label: "Bubble Sort (DESC)", sorting: "bubble", num_elements: n, fashion: "desc"}

	counting_desc.Simulate()
	quick_desc.Simulate()
	bubble_desc.Simulate()
}

/*
* Fato 2: Vetor de tamanho grande, a performance do algoritmo influencia de forma significativa.
* Além disso, dependendo da disposição (e valores) dos elementos no vetor, podemos experimentar
* performances bem diferentes (melhor e pior caso)
 */

func second_experiment() {

	// # of elements
	n := 100000

	// Bubble sort ASC x Quick Sort (pivô aleatório) ASC
	// Conclusão: o bubble sort é melhor que o quick sort (pivô aleatório) nos arrays ordenados ASC

	// bubble_asc := AlgorithmPerfomance{label: "Bubble Sort (ASC)", sorting: "bubble", num_elements: n, fashion: "asc"}
	// quickrandom_desc := AlgorithmPerfomance{label: "Quick Sort Random (ASC)", sorting: "quick-random", num_elements: n, fashion: "asc"}

	// bubble_asc.Simulate()
	// quickrandom_desc.Simulate()

	// Counting Sort (k >> n) e Quick Sort (pivô aleatório)
	// Conclusão: o counting sort é pior do que o quick sort para um array com valores muito grandes

	// counting_sci := AlgorithmPerfomance{label: "Counting Sort (SCI)", sorting: "counting", num_elements: n, fashion: "sci"}
	// quick_sci := AlgorithmPerfomance{label: "Quick Sort (SCI)", sorting: "quick-random", num_elements: n, fashion: "sci"}

	// counting_sci.Simulate()
	// quick_sci.Simulate()

	// Quick Sort ASC x Quick Sort DESC x Insertion Sort
	// Conclusão: o pior caso do quick sort deve ser próximo ao desempenho do insertion sort. Para isso, coloquei
	// o quick sort para ordenar um vetor ordenado decrescentemente com muitos elementos, e o insertion para fazer
	// essas mesma tarefa, só que com um vetor desordenado

	quick_asc := AlgorithmPerfomance{label: "Quick Sort (ASC)", sorting: "quick", num_elements: n, fashion: "asc"}
	quick_desc := AlgorithmPerfomance{label: "Quick Sort (DESC)", sorting: "quick", num_elements: n, fashion: "desc"}
	insertion := AlgorithmPerfomance{label: "Insertion Sort", sorting: "insertion", num_elements: n, fashion: "random"}

	quick_asc.Simulate()
	quick_desc.Simulate()
	insertion.Simulate()
}

/*
* Fato 3: MergeSort tem sempre um desempenho muito bom, independente da
* disposição dos elementos no vetor.
 */

func third_experiment(n int) {

	// Merge Sort x Merge Sort ASC x Merge Sort DESC
	merge_random := AlgorithmPerfomance{label: "Merge Sort", sorting: "merge", num_elements: n, fashion: "random"}
	merge_asc := AlgorithmPerfomance{label: "Merge Sort (ASC)", sorting: "merge", num_elements: n, fashion: "asc"}
	merge_desc := AlgorithmPerfomance{label: "Merge Sort (DESC)", sorting: "merge", num_elements: n, fashion: "desc"}

	merge_random.Simulate()
	merge_asc.Simulate()
	merge_desc.Simulate()
}

/*
* Fato 4: O pior caso do Quicksort é com o vetor ordenado de forma
* crescente/decrescente. O Quicksort com randomização de pivô resolve
* esse mau desempenho.
 */

func fourth_experiment(n int) {

	// Quick Sort ASC x Quick Sort (pivô aleatório) ASC

	quick_asc := AlgorithmPerfomance{label: "Quick Sort (ASC)", sorting: "quick", num_elements: n, fashion: "asc"}
	quick_random_asc := AlgorithmPerfomance{label: "Quick Sort Random (ASC)", sorting: "quick-random", num_elements: n, fashion: "asc"}

	quick_asc.Simulate()
	quick_random_asc.Simulate()

	fmt.Println("")

	// Quick Sort DESC x Quick Sort (pivô aleatório) DESC

	quick_desc := AlgorithmPerfomance{label: "Quick Sort (DESC)", sorting: "quick", num_elements: n, fashion: "desc"}
	quick_random_desc := AlgorithmPerfomance{label: "Quick Sort Random (DESC)", sorting: "quick-random", num_elements: n, fashion: "desc"}

	quick_desc.Simulate()
	quick_random_desc.Simulate()
}

/*
* Fato 5: Explique quando o CountingSort tem bom desempenho e quando
* tem mau desempenho mostrando os resultados através dos experimentos
 */

func fifth_experiment(n int) {

	// Counting Sort (k >> n) x Counting Sort (k < n)
	// Conclusão: quanto mais alto for os valores dos elementos do array, pior a perfomance do counting

	counting_worst := AlgorithmPerfomance{label: "Counting Sort (k >> n)", sorting: "counting", num_elements: n, fashion: "sci"}
	counting_best := AlgorithmPerfomance{label: "Counting Sort (k < n)", sorting: "counting", num_elements: n, fashion: "random"}

	counting_worst.Simulate()
	counting_best.Simulate()
}
