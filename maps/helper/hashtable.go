package helper

// Fator de carga máximo (%)
const MAX_LF = 0.75

// Estruturas básicas

// Obs. É necessário guardar o valor e a chave na tupla
// pois mais de uma chave podem ser mapeadas para uma
// mesma posição da tabela hash

type Tuple struct {
	value string
	key   int
}

type HashTable struct {
	length  int
	buckets [][]Tuple
}

// Funções principais

func (table *HashTable) Init(size int) {
	// Aloque size posições (buckets) para a tabela hash
	table.buckets = make([][]Tuple, size)
	// Ela começa com zero elementos
	table.length = 0
}

func (table *HashTable) Size() int {
	return table.length
}

func (table *HashTable) LoadFactor() float32 {
	return float32(table.length) / float32(len(table.buckets)) // Qtd adicionados / Qtd total
}

func (table *HashTable) Put(key int, value string) {

	// Se fator de carga maior que o permitido, aumente o numero de buckets
	if table.LoadFactor() > MAX_LF {
		table.increase_buckets()
	}

	// Criar nova tupla
	new_tuple := &Tuple{}
	new_tuple.key = key
	new_tuple.value = value

	// Calcular index do bucket onde ela vai ser adicionada
	index := key % len(table.buckets)

	// Adicionar tupla no vetor de tuplas do bucket
	table.buckets[index] = append(table.buckets[index], *new_tuple)

	// Atualizar qtd de elemento inseridos
	table.length++
}

func (table *HashTable) Get(key int) string {
	// Encontre o index do bucket
	index := key % len(table.buckets)
	// Recupere o array de tuplas que se encontra nesse bucket
	array := table.buckets[index]
	// Encontre a chave procurada e retorne o conteúdo dela
	for _, tuple := range array {
		if tuple.key == key {
			return tuple.value
		}
	}
	return "Chave inexistente"
}

func (table *HashTable) Remove(key int) {
	// Encontre o index do bucket onde essa key se encontra
	index := key % len(table.buckets)
	// Recupere o array de tuplas desse bucket
	array := table.buckets[index]
	// Ache o index da chave procurada
	for i, tuple := range array {
		if tuple.key == key {
			// Remova o elemento
			table.buckets[index] = append(array[:i], array[i+1:]...)
			table.length--
			break
		}
	}
}

// Funções auxiliares

func (table *HashTable) increase_buckets() {
	// Criar nova tabela hash com oito vezes o tamanho da original
	new_array := make([][]Tuple, 8*len(table.buckets))

	// Mover os elementos da tabela anterior
	for _, bucket := range table.buckets {
		for _, tuple := range bucket {
			// Mapear chave da tupla no novo array
			index := tuple.key % len(new_array)
			// Inseri-la no bucket
			new_array[index] = append(new_array[index], tuple)
		}
	}

	// Reatribuir endereço da nova tabela
	table.buckets = new_array
}
