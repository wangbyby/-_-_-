#include "stdio.h"
#include <pthread.h>
#include <stdlib.h>
#include <semaphore.h>
int less(int a, int b, int flag)
{
    return (!flag) ? a > b : a < b;
}
void bubble_sort(int arr[], int len, int flag)
{
    int i, j, temp;
    for (i = 0; i < len - 1; i++)
        for (j = 0; j < len - 1 - i; j++)
            if (less(arr[j], arr[j + 1], flag))
            {
                temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
            }
}

sem_t lock;
int data[10] = {10, 9, 39, 230, 40, 1, 42, 32, 45, 450};
int len = 10;
void *thread1_task(void *args)
{
    int *flag = (int *)args;
    sem_wait(&lock);
    bubble_sort(data, len, *flag);
    for (int i = 0; i < len; i++)
    {
        printf("%d\t", data[i]);
    }
    printf("\n");
    sem_post(&lock);
}
int main()
{
    int t1 = 1;
    int t2 = 0;
    int *a1 = &t1;
    int *a2 = &t2;
    sem_init(&lock, 0, 1);
    pthread_t tid1, tid2;
    int e1 = pthread_create(&tid1, NULL, (void *)thread1_task, (void *)a1);
    int e2 = pthread_create(&tid2, NULL, (void *)thread1_task, (void *)a2);
    if (e1 || e2)
    {
        printf("pthread_create error\n");
        return -1;
    }
    pthread_join(tid1, NULL);
    pthread_join(tid2, NULL);
    printf("done\n");
    return 0;
}
