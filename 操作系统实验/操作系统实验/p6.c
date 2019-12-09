#include <stdio.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <sys/types.h>
#include <signal.h>
#include <unistd.h>
void sig_int(int signo)
{
    printf("\n sig_int get %d\n", signo);
    (void)signal(SIGINT, SIG_DFL); //恢复默认行为
}
void sig_usr(int signo)
{
    if (signo == SIGUSR1)
    {
        printf("receive SIGUSR1 %d\n", signo);
    }
    else
    {
        printf("usr received signal %d\n", signo);
    }

    (void)signal(SIGINT, SIG_DFL); //恢复默认行为
}
int main(void)
{
    pid_t child = fork();
    if (child < 0)
    {
        printf("fork error\n");
        return 1;
    }
    else if (child == 0)
    {
        printf("child id = %d\n", getpid());
        (void)signal(SIGINT, sig_usr);
    }
    else
    {
        printf("parent id = %d\n", getpid());
        (void)signal(SIGINT, sig_int);
    }
    pause();
}
