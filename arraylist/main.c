#include "arraylist2.h"

int main()
{
    struct arraylist* l = inicializar(10);

    inserirElementoNoFim(l,10);
    inserirElementoNoFim(l,20);
    inserirElementoNoFim(l,50);
    inserirElementoNoFim(l,30);
    inserirElementoNoFim(l,2);
    inserirElementoNoFim(l,9);
    exibirLista(l);

    printf("\n");
    printf("l[2] = %d", obterElementoEmPosicao(l,2));

    printf("\n");
    printf("l[0] = %d", obterElementoEmPosicao(l,0));

    removerElementoNoFim(l);
    removerElementoNoFim(l);
    removerElementoNoFim(l);

    printf("\n");    
    exibirLista(l);


    return 0;
}