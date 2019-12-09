#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <sys/msg.h>
#include <sys/wait.h>
#define MAX_TEXT_LEN 512
//wsl不能运行
struct user_msg
{
    long int msg_type;
    char text[MAX_TEXT_LEN];
};

int main(int argc, char **argv)
{
    int msgid = -1;
    // 建立消息队列
    msgid = msgget((key_t)1234, 0666 | IPC_CREAT);
    if (msgid == -1)
    {
        printf("msgget error");
        exit(EXIT_FAILURE);
    }
    pid_t reciver = fork();
    if (reciver < 0)
    {
        printf("fork error");
        return -1;
    }
    else if (reciver == 0)
    {
        struct user_msg data;
        //server 进程
        while (1)
        {
            if (msgrcv(msgid, (void *)&data, BUFSIZ, 0, 0) == -1)
            {
                printf("msgrcv error\n"); //some error here
                goto del;
            }
            printf("You wrote: %s\n", data.text);
            // 遇到end结束
            if (strncmp(data.text, "end", 3) == 0)
            {
                break;
            }
        }
    del:
        // 删除消息队列
        if (msgctl(msgid, IPC_RMID, 0) == -1)
        {
            printf("msgctl(IPC_RMID) failed\n");
        }
        exit(EXIT_SUCCESS);
    }
    else
    {
        pid_t sender = fork();
        if (sender < 0)
        {
            printf("fork error");
            return -1;
        }
        else if (sender == 0)
        {
            struct user_msg data;
            char buffer[BUFSIZ];
            // 向消息队里中写消息，直到写入end
            while (1)
            {
                printf("Enter some text: \n");
                fgets(buffer, BUFSIZ, stdin);
                data.msg_type = 1; // 注意2
                strcpy(data.text, buffer);
                // 向队列里发送数据I
                if (msgsnd(msgid, (void *)&data, MAX_TEXT_LEN, 0) == -1)
                {
                    printf("msgget error");
                    exit(EXIT_FAILURE);
                }
                // 输入end结束输入
                if (strncmp(buffer, "end", 3) == 0)
                {
                    break;
                }
                sleep(1);
            }
            exit(EXIT_SUCCESS);
        }
        else
        {
            while (wait(NULL) != -1)
            {
            }
        }
    }
}
