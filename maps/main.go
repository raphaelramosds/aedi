package main

import "dca208.com/hashtable/helper"

func main() {
	table := helper.HashTable{}

	table.Init(10)

	// Inserção de elementos
	table.Put(1884, "Raphael Ramos da Silva")
	table.Put(9647, "Bernardo Brito")
	table.Put(2282, "Rafael Ribeiro")
	table.Put(9048, "Maria Clara Freitas")
	table.Put(9567, "Niedson Fernando Dantas")
	table.Put(9587, "Poliana Ellen")

	// Busca por chave
	println(table.Get(9647))

	// Buscar elemento com colisão
	println(table.Get(9567))
}
