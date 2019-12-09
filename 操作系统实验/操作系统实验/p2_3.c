#include "stdio.h"
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>
int main()
{

	int pipe_fd[2];
	int status = pipe(pipe_fd);
	if (status == -1)
	{
		printf("pipe create failed\n");
		return -1;
	}
	__pid_t pid1, pid2;

	pid1 = fork();

	if (pid1 < 0)
	{
		printf("function fork() error\n");
		printf("pid1 = %d \n ", pid1);
	}
	else if (pid1 == 0)
	{
		// printf("pid1 = %d \n pid2 = %d\n",pid1, pid2);
		char buf[1000] = "Child 1 is sending a message!\n";
		if (write(pipe_fd[1], buf, strlen(buf)) != -1)
		{

			printf("child pid = %d parent pid = %d child 1 send success\n", getpid(), getppid());
		}
		else
		{
			printf("child1 send error\n");
		}
	}

	else
	{
		pid2 = fork();
		if (pid2 < 0)
		{
			printf("function fork() error\n");
			printf("pid2 = %d \n ", pid2);
			return 1;
		}
		else if (pid2 == 0)
		{
			// printf("pid1 = %d \n pid2 = %d\n",pid1, pid2);
			char buf[1000] = "Child 2 is sending a message!\n";
			if (write(pipe_fd[1], buf, strlen(buf)) != -1)
			{

				printf("child pid = %d parent pid = %d child2 send success\n", getpid(), getppid());
			}
			else
			{
				printf("child2 send error\n");
			}
		}
		else
		{
			// printf("parent pid = %d\n",getpid());
			printf("parent gets the mes!\n");
			int rnum = 0, i = 0;
			char buf[100];
			while (wait(NULL) != -1)
			{}
			// wait(pid1, NULL, 0);
			// wait(pid2, NULL, 0);
			close(pipe_fd[1]);
			rnum = read(pipe_fd[0], buf, 100);
			if (rnum > 0)
			{
				for (i = 0; i < rnum; i++)
				{
					printf("%c", buf[i]);
				}
			}
			close(pipe_fd[0]);
		}
	}
	return 0;
}
