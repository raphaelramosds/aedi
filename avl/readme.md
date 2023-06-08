# Árvores Balanceadas AVL

Árvores AVL são árvores de busca binária balanceadas: isso significa que minimizam o número de comparações efetuadas no pior caso para operações de busca, inserção e remoção.

> O nome AVL vem de seus criadores soviéticos Adelson Velsky e Landis, e sua primeira referência encontra-se no documento "Algoritmos para organização da informação" de 1962.

## Rotações

A estratégia que uma árvore AVL utiliza para se rebalancear são as rotações de nós desbalanceados.

Na **rotação à esquerda** de um nó raiz, seguem-se os passos

- Seu filho da direita vira a nova raiz
- O filho da esquerda da nova raiz vira filho da direita da raiz original
- A raiz original vira filho da esquerda da nova raiz
- Retornamos o endereço da nova raiz `right`

Já na **rotação à direita** de um nó raiz, seguem-se os passos

- Seu filho da esquerda vira a nova raiz
- O filho da direita da nova raiz vira filho da esquerda da raiz original
- A raiz original vira filho da direita da nova raiz
- Retornamos o endereço da nova raiz `left`

## Rebalanceamento

Para fazer o rebalanceamento, adicionamos mais duas propriedades ao tipo estruturado BSTNode: a altura do nó `height` e seu fator de balanceamento `bf`

```go
type BSTNode struct {
	left   *BSTNode
	value  int
	height int
	bf     int
	right  *BSTNode
}
```

O fator de balanceamento de um nó é a diferença entre a altura da sua subárvore direita (Hr) e a altura da subárvore esquerda (He)

```
Fator de balanceamento = Hr - He
```