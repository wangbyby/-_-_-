#include "stdio.h"
#include <pthread.h>
#include <semaphore.h>
#include <stdlib.h>
int flag = 0;
void compute(int a, int b)
{
    if (flag == 1)
    {
        printf("%d * %d = %d\n", a, b, a * b);
    }
    else if (flag == 0)
    {
        printf("%d + %d = %d\n", a, b, a + b);
    }
}

int value[2] = {0, 0};
sem_t empty, full, mutex, empty2, full2;
pthread_t r1, r2, c1, c2;
void *read_file(void *file_name)
{
    char *filename = (char *)file_name;
    FILE *r_file = fopen(filename, "r");

    if (r_file == NULL)
    {
        return NULL;
    }
    int num = 0;
    char tmp[1024];
    for (int i = 0; i < 10; i++)
    {
        fgets(tmp, 1024, r_file);
        num = atoi(tmp);
        if (filename == "1.dat")
        {
            sem_wait(&empty);
            sem_wait(&mutex);
            value[0] = num;
            sem_post(&mutex);
            sem_post(&full);
        }
        else if (filename == "2.dat")
        {
            sem_wait(&empty2);
            sem_wait(&mutex);
            value[1] = num;
            sem_post(&mutex);
            sem_post(&full2);
        }
    }
}
void *consumer(void *args)
{
    int i, j, a, b;
    for (i = 0; i < 5; i++)
    {
        for (j = 0; j < 2; j++)
        {
            sem_wait(&full);
            sem_wait(&full2);
            sem_wait(&mutex);
            a = value[0];
            b = value[1];
            compute(a, b);
            flag = 1^flag;
            sem_post(&mutex);
            sem_post(&empty);
            sem_post(&empty2);
        }
    }
}
int main()
{
    int t1, t2, t3, t4;

    sem_init(&empty, 0, 1);
    sem_init(&empty2, 0, 1);

    sem_init(&mutex, 0, 1);
    sem_init(&full, 0, 0);
    sem_init(&full2, 0, 0);
    t1 = pthread_create(&c1, NULL, consumer, NULL);
    t2 = pthread_create(&c2, NULL, consumer, NULL);
    t3 = pthread_create(&r1, NULL, read_file, (void *)"1.dat");
    t4 = pthread_create(&r2, NULL, read_file, (void *)"2.dat");
    pthread_join(r1, NULL);
    pthread_join(r2, NULL);
    pthread_join(c1, NULL);
    pthread_join(c2, NULL);
    return 0;
}
