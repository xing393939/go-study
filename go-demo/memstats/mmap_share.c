// 参考：https://phpor.net/blog/post/893
#include<stdio.h>
#include<stdlib.h>
#include<unistd.h>
#include<sys/mman.h>
#include<sys/types.h>
#include<sys/stat.h>
#include<fcntl.h>

long long ccc[4];

#ifdef _WIN32
void *mmap(void *addr, size_t length, int protect, int flags, int fd, off_t offset) {
    return 0;
}
#endif

int main2(int ac,char**av) {
	char*fp=(char*)mmap(NULL,1024*1024*200,PROT_READ,MAP_SHARED|MAP_ANONYMOUS,-1,0);
	char c;
	int i=0;
	if(fp==MAP_FAILED) {
		printf("error\n");
		exit(1);
	}
	while(i++<1024*1024*100) {
		c=*(fp+i);
	}
	printf("%p\n", &ccc);
	sleep(20000000);
	return 0;
}