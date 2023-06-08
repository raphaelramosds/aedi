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

Abaixo está ilustrada a rotação à esquerda do nó 1

```
    1             3
     \           / \
      3    →    1   4    →  retorne endereço de 3
     / \         \
    2   4         2
```

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

O fator de balanceamento de um nó é a diferença entre a altura da sua subárvore direita (Hd) e a altura da subárvore esquerda (He)

```
Fator de balanceamento = Hd - He
```

## Casos de balanceamento

Existem alguns casos em que apenas uma rotação à esquerda ou a direita não vão rebalancear a árvore cujo nó analisado é a raiz. Para identificar esses casos, você precisa analisar o fator de balanceamento do nó que você quer rotacionar (o Nó raiz `root`) e os fatores de balanceamento dos seus filhos esquerdo e direito.

A seguir estão duas subseções que apresentam as operações de rotações que precisam ser feitas para que a árvore do nó raiz analisado seja balanceada. 

Nessas análises, considere que o objeto `root *BSTNode`, pertencente a uma árvore binária de busca, foi passado como parâmetro para uma função de balanceamento e nela foram efetuadas as operações presentes na coluna Operações.

### Análise do filho esquerdo

| bf Nó raiz | bf Filho esquerdo | Operações                               |
|------------|-------------------|-----------------------------------------|
| -2         | -1                | `root.RotRight()`                       |
| -2         | 0                 | `root.RotRight()`                       |
| -2         | 1                 | `root.left.RotLeft()` `root.RotRight()` |

### Análise do filho direito

| bf Nó raiz | bf Filho direito | Operações                                |
|------------|------------------|------------------------------------------|
| 2          | 1                | `root.RotLeft()`                         |
| 2          | 0                | `root.RotLeft()`                         |
| 2          | -1               | `root.right.RotRight()` `root.RotLeft()` |