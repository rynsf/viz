#include <stdio.h>
#include <pthread.h>
#include <unistd.h>
#include <stdlib.h>

void* foo(void* args) {
    while(1) {
        printf("Running\n");
        sleep(1);
    }
}

void* bar(void* args) {
    while(1) {
        printf("Running every 2 sec\n");
        sleep(2);
    }
}

int main() {
    pthread_t t1, t2;
    pthread_create(&t1, NULL, foo, NULL);
    pthread_create(&t2, NULL, bar, NULL);
    pthread_join(t2, NULL);
    pthread_join(t1, NULL);
    return 0;
}
