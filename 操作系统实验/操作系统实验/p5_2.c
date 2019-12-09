#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include <string.h>
#define TEXT 2048
struct shared_use_st
{
    int read;
    char text[TEXT];
};
typedef struct shared_use_st shared_use_st;
int main()
{
    pid_t reciver, sender;
    int shmid; //共享内存的id
    //创建共享内存
    shmid = shmget((key_t)1234, sizeof(shared_use_st), 0666 | IPC_CREAT);
    if (shmid == -1)
    {
        printf("shmget error\n");
        return -1;
    }
    reciver = fork();
    if (reciver < 0)
    {
        printf("fork() error\n");
        return -1;
    }
    else if (reciver == 0) //子
    {
        //进程链接到 共享内存
        void *shm_client = NULL;
        shm_client = shmat(shmid, 0, 0);
        if (shm_client == (void *)-1)
        {
            printf("shmat error\n");
            return -1;
        }
        // 指向 share mem的指针=
        shared_use_st *shared = (shared_use_st *)shm_client;
        shared->read = 0;
        while (1)
        {
            if (shared->read == 1)
            {
                printf("you wrote:%s\n", shared->text);
                int a = atoi(shared->text);
                printf("平方根:%d\n", a * a);
                shared->read = 0; // 可写
                sleep(1);
                if (strncmp(shared->text, "end", 3) == 0)
                {
                    break;
                }
            }
            else
            {
                sleep(1);
            }
        }
        // 进程与共享内存分离
        if (shmdt(shm_client) == -1)
        {
            printf("shmdt error\n");
            return -1;
        }
    }
    else
    {
        sender = fork(); 
        if (sender < 0)
        {
            printf("fork() error\n");
            return -1;
        }
        else if (sender == 0) //子
        {
            void *shm_server = shmat(shmid, (void *)0, 0);
            if (shm_server == (void *)-1)
            {
                printf("shmat error\n");
                return -1;
            }
            shared_use_st *shared = (shared_use_st *)shm_server;
            while (1)
            {
                while (shared->read == 1)
                {
                    sleep(1);
                    printf("Please Waiting...\n");
                }
                printf("Enter some text: ");
                char buffer[BUFSIZ + 1];

                fgets(buffer, BUFSIZ, stdin);
                strncpy(shared->text, buffer, TEXT);
                shared->read = 1;
                if (strncmp(buffer, "end", 3) == 0)
                {
                    break;
                }
            }
            if (shmdt(shm_server) == -1)
            {
                printf("shmdt error\n");
                return -1;
            }
        }
        else //父
        {
            while (wait(NULL) != -1)
            {
            }

            //删除共享内存
            if (shmctl(shmid, IPC_RMID, 0) == -1)
            {
                printf("shmctl IPC_RMID error\n");
                return -1;
            }
        }
    }

    return 0;
}
