#include <stdio.h>
#include <stdlib.h>

struct no
{
  int val;
  struct no *prox;
};

struct linkedlist
{
  struct no *cabeca;
  int qtdade;
};

struct linkedlist *inicializar()
{
  struct linkedlist *ll = (struct linkedlist *)malloc(sizeof(struct linkedlist));
  ll->cabeca = NULL;
  ll->qtdade = 0;
  return ll;
}

struct no *alocarNovoNo(int valor)
{
  struct no *novo = (struct no *)malloc(sizeof(struct no));
  novo->prox = NULL;
  novo->val = valor;
  return novo;
}

void inserirElementoNoFim(struct linkedlist *lista, int valor)
{
  struct no *novo = alocarNovoNo(valor);

  if (lista->cabeca == NULL)
  {
    lista->cabeca = novo;
  }
  else
  {
    struct no *aux = lista->cabeca;
    while (aux != NULL)
    {
      if (aux->prox == NULL)
      {
        aux->prox = novo;
        novo->prox = NULL;
      }
      aux = aux->prox;
    }
  }

  lista->qtdade++;
}

void inserirElementoNoInicio(struct linkedlist *lista, int valor)
{
  struct no *novo = alocarNovoNo(valor);

  if (lista->cabeca == NULL)
  {
    lista->cabeca = novo;
  }
  else
  {
    struct no *temp = lista->cabeca;
    lista->cabeca = novo;
    novo->prox = temp;
  }
  lista->qtdade++;
}

void inserirElementoEmPosicao(struct linkedlist *lista, int valor, int posicao)
{
  // TODO
}

int obterElementoEmPosicao(struct linkedlist *lista, int posicao)
{
  if (posicao < 0 || posicao > lista->qtdade - 1)
  {
    printf("Index invalido");
  }

  if (lista->cabeca == NULL)
  {
    printf("A lista esta vazia");
  }

  struct no *aux = lista->cabeca;
  int count = 0;

  while (aux != NULL)
  {
    if (count == posicao)
    {
      return aux->val;
    }
    aux = aux->prox;
    count++;
  }
}

void removerElementoEmPosicao(struct linkedlist *lista, int posicao)
{
  if (posicao < 0 || posicao > lista->qtdade - 1)
  {
    printf("Posicao invalida: %d", posicao);
    return;
  }

  struct no *aux = lista->cabeca;

  if (posicao == 0)
  {
    lista->cabeca = NULL;
    lista->cabeca = aux->prox;
    free(aux);
    return;
  }

  int count = 0;

  while (aux != NULL)
  {
    if (count == posicao - 1)
    {

      // Guardo endereço do elemento a ser removido e do posterior
      struct no *temp = aux->prox;
      struct no *front = temp->prox;

      // Retiro ponteiro que aponta para o elemento removido
      aux->prox = NULL;

      // E coloco-o no elemento posterior
      aux->prox = front;

      // Libero memória do elemento removido
      free(temp);

      return;
    }
    aux = aux->prox;
    count++;
  }

  lista->qtdade--;
}

void imprimirLista(struct linkedlist *lista)
{
  if (lista->cabeca != NULL)
  {
    struct no *aux = lista->cabeca;
    printf("[");
    do
    {
      printf("%d", aux->val);
      aux = aux->prox;
      if (aux != NULL)
      {
        printf(", ");
      }
    } while (aux != NULL);
    printf("]");
  }
  else
  {
    printf("A lista esta vazia!");
  }
}