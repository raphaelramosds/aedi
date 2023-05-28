# Árvores de  Busca Binária

A pasta helper possui duas implementações de BST: a primeira está em `bst.go` que faz uma implementação estruturada, e a segunda está em `bst2.go` que faz uma implementação mais próxima da orientação a objetos.

## Perguntas e respostas

**1. Qual a diferença entre uma árvore binária e uma árvore de busca binária (BST)?**

Uma árvore de busca binária é uma árvore binária onde cada nó possui a sua subárvore esquerda com valores numéricos inferiores a ele, e sua subárvore direita possui valores numéricos superiores a ele. Um árvore binária pode ou não respeitar essa propriedade.

**2. O que significam os níveis de uma árvore?**

Os níveis de uma árvore representam o número de arestas entre cada um dos nós até a raiz. O nível da raiz, por exemplo, é zero.

**3. O que é uma BST completa?**

É uma BST onde cada nó, com exceção das folhas, possuem exatamente dois filhos e as folhas estão no mesmo nível

**4. Como calculamos o fator de balanceamento de uma árvore?**

O fator de balanceamento de uma árvore é a diferença entre a altura da subárvore esquerda (He) e a altura da subárvore direita (Hr)

```
Fator de balanceamento = He - Hr
```

**5. Qual a altura de uma árvore balanceada contendo n nós? Mostre matematicamente**

Vamos aproximar essa árvore balanceada como uma árvore balanceada completa. Assim, cada nível terá o número de nós como uma potência de dois.

```
Nível 0 (i = 0) → n = 2° nós

Nível 1 (i = 1) → n = 2¹ nós

Nível 2 (i = 2) → n = 2² nós

Nível 3 (i = 3) → n = 2³ nós

...

Nível h (i = h) → n = 2ʰ nós
```

Como a árvore é completa, veja que h é a sua altura pois ela já representa o número máximo de arestas da última folha até a raiz. Somando o número de nós até o nível h, teremos uma progressão geométrica de razão r = 2

```
n = 2° + 2¹ + ... + 2ʰ → n = 2ʰ⁺¹ - 1
```

Aplicando o logaritmo na base 2 dos dois lados, temos uma expressão para a altura da árvore

```
h = log(n + 1) - 1
```

**6. Qual a complexidade de tempo das seguintes operações em uma árvore completamente desbalanceada? Add(value int), Search(value int) bool, Min() int, Max() int, PrintPre(),PrintIn(), PrintPos(), Height() int, Remove(value int) \*BstNode**

Todas elas são O(n) no pior caso

**7. Qual a complexidade de tempo das seguintes operações em uma árvore completa? Add(value int), Search(value int) bool, Min() int, Max() int, PrintPre(), PrintIn(), PrintPos(), Heigh () int, Remove(value int) \*BstNode**

| Operação | Árvore completa (BST) | Árvore completa (não BST) |
|----------|-----------------------|---------------------------|
| Add      | O(h)                  | O(n)                      |
| Search   | O(h)                  | O(n)                      |
| Min      | O(h)                  | O(n)                      |
| Max      | O(h)                  | O(n)                      |
| PrintPre | O(n)                  | O(n)                      |
| PrintIn  | O(n)                  | O(n)                      |
| PrintPos | O(n)                  | O(n)                      |
| Height   | O(h)                  | O(h)                      |
| Remove   | O(h)                  | O(n)                      |

**Observação:** Remove e Add se tornam O(n) na árvore completa (não BST) porque, no pior caso, o critério para fazer essas operações pode se basear em percorrer todos os elementos da árvore.

**8.Explique os possíveis casos de remoção em uma BST. Como deve-se proceder em cada caso?**

- Caso 1 (nó removido é uma folha)

- Caso 2 (nó removido tem um filho)

- Caso 3 (nó removido tem dois filhos)

## Dicas de implementação

### Inserindo elementos na BST

Para adicionar elementos em uma árvore de busca binária sempre gosto de raciocinar assim

1. Devo adicionar na sub-arvores direita ou esquerda?

2. Se a subarvore esquerda ou direita for diferente de NULL, chama recursivamente; se for NULL, simplesmente adiciona.

### Navegando na árvore

Há duas formas para percorrer uma árvore: em largura e em profundidade. A que foi implementada aqui foi os métodos em profundidade, que são três: pré ordem, em ordem e pós ordem.

Para lembrar desses três tipos de navegação, lembre que o prefixo desses nomes dizem sobre em qual ordem a raiz vai ser explorada

1. **Pré ordem:** a raíz é imprimida antes do nó esquerdo e do direito serem explorados

```
Raiz → Nó esquerdo → Nó direito
```

2. **Em ordem:** a raiz é imprimida logo após o nó esquerdo ser explorado

```
Nó esquerdo → Raiz → Nó direito
```

3. **Pós ordem:** a raiz é imprimida logo após o nó esquerdo e depois o nó direito serem explorados

```
Nó esquerdo → Nó direito → Raiz
```