#include <stdio.h>
#include <stdlib.h>
#define datatype int
typedef struct{
    datatype *arr;
    int front, tail;
    int cap ,len;
}Queue;
Queue  NewQueue(int cap){
    Queue  q = { 
        .front = 0,
        .tail = 0,
        .len = 0,
        .cap = cap,
        .arr = (datatype *)malloc(sizeof(datatype)*cap),
    };
    return q;
}
int queue_full(Queue *q){
    return (q->cap == q->len);
}
int queue_empty(Queue *q) {
    return (q->len ==0);
}
void queue_push_back(Queue *q,datatype element){
    if(queue_full(q)){
        return ;
    }
    q->arr[q->tail] = element;
    q->tail = (q->tail+1)%q->cap;
    q->len++;
}
datatype queue_pop_head(Queue *q) {
    if (!queue_empty(q))
    {
        datatype tmp = q->arr[q->front];
        q->front = (q->front +1)%q->cap;
        q->len--;
        return tmp;
    }
    
}

int main(){
    Queue q = NewQueue(2);
    int a =20;
    for (int i = 0; i < a; i++)
    {
        queue_push_back(&q,i);
        
        int c = pop(&q);
        printf("value = %d\n",c);
    }
    return 0;
}