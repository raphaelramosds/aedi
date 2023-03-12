#include "arraylist2.h"

int main()
{
    struct arraylist *l = inicializar(3);

    inserirElementoNoFim(l, 1);
    inserirElementoNoFim(l, 9);
    inserirElementoNoFim(l, 3);
    inserirElementoNoFim(l, 3);
    inserirElementoNoFim(l, 2);
    removerElementoNoFim(l);
    removerElementoNoFim(l);
    removerElementoNoFim(l);
    inserirElementoNoFim(l, 9);
    inserirElementoNoFim(l, 10);
    atualizarElemento(l, 12, 2);
    exibirLista(l);
}