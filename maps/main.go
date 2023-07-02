package main

import "dca208.com/hashtable/helper"

func main() {
	table := helper.HashTable{}

	table.Init(6)

	// Inserção de elementos
	table.Put(1884, "Raphael Ramos da Silva")
	table.Put(9647, "Bernardo Brito")
	table.Put(2282, "Rafael Ribeiro")
	table.Put(9048, "Maria Clara Freitas")
	table.Put(9567, "Niedson Fernando Dantas")
	table.Put(9587, "Poliana Ellen")
	table.Put(1507, "Anelma Silva")

	// Remover elementos
	table.Remove(9567)
	table.Remove(9647)

	// Recuperar elementos
	println(table.Get(9587))
	println(table.Get(1507))

	// Recuperar elemento removido
	println(table.Get(9567))
	println(table.Get(9647))

	// Imprimir fator de carga
	println(table.LoadFactor())

}
