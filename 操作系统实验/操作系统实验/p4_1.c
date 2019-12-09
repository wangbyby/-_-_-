#include "stdio.h"
#include <pthread.h>
#include <semaphore.h>
#include <stdlib.h>
#include <time.h>
#define __compute() ((rand() % 2) >= 1 ? 0 : 1)
void compute(int a, int b)
{
	if (__compute())
	{
		printf("%d * %d = %d\n", a, b, a * b);
	}
	else
	{
		printf("%d + %d = %d\n", a, b, a + b);
	}
}
#define datatype int
typedef struct
{
	datatype *arr;
	int front, tail;
	int cap, len;
} Queue;
Queue NewQueue(int cap)
{
	Queue q = {
		.front = 0,
		.tail = 0,
		.len = 0,
		.cap = cap,
		.arr = (datatype *)malloc(sizeof(datatype) * cap),
	};
	return q;
}
int queue_full(Queue *q)
{
	return (q->cap == q->len);
}
int queue_empty(Queue *q)
{
	return (q->len == 0);
}
void queue_push_back(Queue *q, datatype element)
{
	if (queue_full(q))
	{
		return;
	}
	q->arr[q->tail] = element;
	q->tail = (q->tail + 1) % q->cap;
	q->len++;
}
datatype queue_pop_head(Queue *q)
{
	if (!queue_empty(q))
	{
		datatype tmp = q->arr[q->front];
		q->front = (q->front + 1) % q->cap;
		q->len--;
		return tmp;
	}
}
Queue *queue;
int value[2] = {0, 0};
sem_t empty, full, mutex;
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
		sem_wait(&empty);
		sem_wait(&mutex);
		queue_push_back(queue, num);
		sem_post(&mutex);
		sem_post(&full);
	}
}
void *consumer(void *args)
{

	int i, j;
	for (i = 0; i < 5; i++)
	{
		for (j = 0; j < 2; j++)
		{
			sem_wait(&full);
			sem_wait(&mutex);

			value[j] = queue_pop_head(queue);
			sem_post(&mutex);
			sem_post(&empty);
		}

		int a = value[0];
		int b = value[1];

		compute(a, b);
	}
}
int main()
{
	srand((int)time(0));
	int t1, t2, t3, t4;
	Queue queue1 = NewQueue(2);
	queue = &queue1;
	sem_init(&empty, 0, 2);
	sem_init(&mutex, 0, 1);
	sem_init(&full, 0, 0);
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

// struct Array* list1= (struct Array*)malloc(sizeof(struct Array));
// list1->flag = 0;
// list1->arr = data;
// list1->len = 10;
