#include<stdio.h>
#include<stdlib.h>
#include<unistd.h>

#define ONE_MB 1<<20
#define ONE_KB 1<<10

int main(void) {
  int i;
  char *p;

  for (i=0;i<1000;i++) {
    p = malloc(ONE_KB);
    if (p == NULL) {
      printf("Failed to allocate at %d kB\n",i);
    }
  }
}
