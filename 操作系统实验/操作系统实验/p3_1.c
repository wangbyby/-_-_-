#include "stdio.h"
#include <pthread.h>
#include <stdlib.h>

int count =0;
static void * thread_task(void *args)
{
    count++;
    printf("thread %d is running\n", count);
    
}

int main()
{

    pthread_t tid1, tid2;

    int i1 = pthread_create(&tid1, NULL, thread_task, NULL);
    int i2 = pthread_create(&tid2, NULL, thread_task, NULL);

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

