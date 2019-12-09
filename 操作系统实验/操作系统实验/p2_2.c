#include <unistd.h>
#include <stdio.h>

int main()
{
	int m = 10;
	int i = 0;
	pid_t t1, t2;
	t1 = fork();

	if (t1 < 0)
	{
		printf("fork error t1=%d\n", t1);
		return 0;
	}
	else if (t1 == 0)
	{
		for (int i = 0; i < m; i++)
		{
			printf("C");
		}

		printf("\n");
	}
	else
	{
		t2 = fork();
		if (t2 < 0)
		{
			printf("fork error t2=%d\n", t2);
			return 1;
		}
		else if (t2 == 0)
		{
			for (int i = 0; i < m; i++)
			{
				printf("B");
			}

			printf("\n");
		}
		else
		{
			for (int i = 0; i < m; i++)
			{
				printf("A");
			}

			printf("\n");
		}
	}
	return 0;
}

