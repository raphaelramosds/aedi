# Árvores de  Busca Binária

A pasta helper possui duas implementações de BST: a primeira está em `bst.go` que faz uma implementação estruturada, e a segunda está em `bst2.go` que faz uma implementação mais próxima da orientação a objetos.

## Navegando na árvore

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