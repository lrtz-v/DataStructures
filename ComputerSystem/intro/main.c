#include <unistd.h>
#include <stdio.h> 

int main(void)
{
   int i=0;
   printf("i\tson/pa\tppid\tpid\tfpid\n");
   //ppid指当前进程的父进程pid
   //pid指当前进程的pid,
   //fpid指fork返回给当前进程的值
   for(i=0;i<2;i++){
       pid_t fpid=fork();
       if(fpid==0)
    	   printf("%d\tchild\t%4d\t%4d\t%4d\n",i,getppid(),getpid(),fpid);
       else
    	   printf("%d\tparent\t%4d\t%4d\t%4d\n",i,getppid(),getpid(),fpid);
   }
   return 0;
}