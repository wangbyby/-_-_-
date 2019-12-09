#include "stdio.h"
#include <pthread.h>
#include <stdlib.h>

struct Args
{
    int num;
    char ch;
};

void *thread_task(void *args)
{
    struct Args* b = (struct Args *)args;
    printf("thread %d  %c is running\n",b->num,b->ch);
}

int main()
{

    pthread_t tid1, tid2;
    struct Args* b;
    b = (struct Args *)malloc(sizeof(struct Args));
    b->num = 10;
    b->ch = 'c';
    int i1 = pthread_create(&tid1, NULL, thread_task, (void *)b);
    int i2 = pthread_create(&tid2, NULL, thread_task,(void *)b);

    if (i1 + i2 < 0)
    {
        printf("pthread_create error\n");
        return -1;
    }

    if (pthread_join(tid1, NULL) && pthread_join(tid2, NULL))
    {
        printf("thread is not exit...\n");
        return -2;
    }
    printf("done");
    return 0;
}
