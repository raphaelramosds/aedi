# Algoritmos de ordenação

## Análise de complexidade de tempo

Relembre que a notação do *big O* representa o pior caso, e o *big omega* o melhor caso. O *theta* representa o melhor e pior caso

### Selection sort

- `θ(n²)` pois, independente da lista, esse algoritmo sempre vai percorrer n - i elementos nas n iterações para achar o menor valor, em que n é o tamanho do array e i é o número da iteração atual

### Bubble sort

- `O(n²)` para ordenar uma lista em ordem decrescente, pois os maiores elementos da lista está nas primeiras posições

- `Ω(n)` para ordenar uma lista em ordem ascendente, pois nenhuma troca será feita na primeira iteração

### Insertion sort

- `O(n²)` para ordenar uma lista em ordem decrescente

- `Ω(n)` para ordenar uma lista em ordem ascendente, pois em cada iteração será identificado que o elemento da mão direita a ser inserido é maior que todos os elementos da mão esquerda. Logo, não há necessidade do loop interno varrer os n - i elementos mão esquerda em cada uma das i iterações, em que n é o tamanho da lista 

### Merge sort

- `θ(nlog(n))` pois, independente da lista, esse algoritmo faz chamadas recursivas em partições da lista que são a metade da partição anterior e, dadas duas partições já ordenadas, é utilizada uma função auxiliar que percorre todos os elementos dessas partições. Portanto, são log(n) chamadas e, em cada uma delas, são executadas n - i iterações, em que n é o tamanho das duas partições analisadas

### Quick sort

- `O(n²)` para ordenar uma lista em ordem decrescente e ascendente. No caso da lista em ordem ascendente, o maior elemento sempre vai estar no final da partição, portanto, o index do pivô vai iterar sempre até o último elemento da partição. Nos dois casos vamos ter n chamadas recursivas com n - i operações em cada uma, em que n é o tamanho do array e i é o número da iteração

- `θ(nlog(n))` para ordenar uma lista cujas partições não fiquem ordenadas após a troca dos elementos

### Counting sort

Aqui, considere k como a diferença entre o maior e o menor valor

- `O(n + k), k >> n` quando temos listas cujos elementos são muito grandes. Nesse caso, faríamos o código processar e alocar uma lista auxiliar com muitos elementos na memória

- `Ω(n + k), n > k` pois mesmo com uma lista com muitos elementos para ordenar, o custo computacional para mapear e ordenar os elementos dela ainda vai se linear

## Perguntas rápidas

**1. Por que SelectionSort não consegue melhorar o desempenho para cenários nos quais o vetor já está ordenado?**

Porque o SelectionSort sempre vai precisar procurar o índice com o menor elemento a cada iteração.

**2. Por que BubbleSort consegue melhorar o desempenho para cenários nos quais o vetor já está ordenado?**

Porque na 1° varredura, se não foi necessário nenhuma troca, os elementos já estão na posição correta, pois os antecessores já são menores que seus sucessores.

**3. Por que InsertionSort consegue melhorar o desempenho para cenários nos quais o vetor já está ordenado?**

Porque ao escolher um elemento para ser inserido, todos na mão esquerda já são menores que ele, fazendo o algoritmo pular para o próximo elemento da mão direita.

**4. Por que o MergeSort sempre tem o mesmo desempenho para qualquer cenário (vetor organizado de diferentes formas)?**

Porque ele faz chamadas recursivas sempre com a metade da lista - independente da sua natureza.

**5. Por que o pior caso do QuickSort é O(n²)?**

Porque se a este algoritmo for dado um vetor em ordem ascendente, ele vai sempre percorrer até o último elemento da partição com n - i elementos, em que i é o nùmero da iteração e n o tamanho da lista original

**6. Como mitigar a probabilidade do pior caso acontecer no QuickSort?**

O índice do pivô em cada iteração deve ser escolhido aleatoriamente e a posição correspondente a esse índice deve ser trocada com a última posição da partição

**7. O CountingSort é melhor ou pior do que o MergeSort? E em relação ao QuickSort?**

Ele possui melhor perfomance que os dois no seu melhor caso devido a complexidade linear. Porém, ele se torna pior que os dois caso o intervalo dos seus elementos seja suficientemente grande.

## Experimentos

A seguir são apresentados 5 fatos sobre algoritmos de ordenação. Planeje, execute experimentos, e apresente resultados que evidenciem cada afirmação.

> **Fato 1:** Para vetores de tamanho pequeno, a performance da maioria dos algoritmos de ordenação não vai influenciar, independente da disposição dos elementos

> **Fato 2:** Vetor de tamanho grande, a performance do algoritmo influencia de forma significativa. Além disso, dependendo da disposição (e valores) dos elementos no vetor, podemos experimentar performances bem diferentes (melhor e pior caso).

> **Fato 3:** MergeSort tem sempre um desempenho muito bom, independente da disposição dos elementos no vetor.

> **Fato 4:** O pior caso do Quicksort é com o vetor ordenado de forma crescente/decrescente. O Quicksort com randomização de pivô resolve esse mau desempenho.

> **Fato 5:** Explique quando o CountingSort tem bom desempenho e quando tem mau desempenho mostrando os resultados através dos experimentos.