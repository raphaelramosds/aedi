#include "linkedlist2.h"

int main()
{
  struct linkedlist *ll = inicializar();

  inserirElementoNoInicio(ll, 12);
  inserirElementoNoInicio(ll, 14);
  inserirElementoNoInicio(ll, 15);
  inserirElementoNoFim(ll, 90);
  inserirElementoNoFim(ll, 100);
  inserirElementoNoInicio(ll, 20);
  inserirElementoNoInicio(ll, 30);
  imprimirLista(ll);
  printf("\n");
  removerElementoEmPosicao(ll, 0);
  removerElementoEmPosicao(ll, -2);
  imprimirLista(ll);
  printf("\n");
  obterElementoEmPosicao(ll,12);
}