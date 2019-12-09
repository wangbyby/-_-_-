#include <unistd.h>
#include <stdio.h>

int main()
{
    pid_t t1, t2;
    t1 = fork();
    if (t1 < 0)
    {
        printf("fork error t1=%d\n", t1);
        return 0;
    }
    else if (t1 == 0)
    {
        printf("C");
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
            printf("B");
        }
        else
        {
            printf("A\n");
        }
    }
    return 0;
}
