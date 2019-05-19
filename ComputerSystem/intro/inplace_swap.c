#include <stdio.h>

void inplace_swap(int * x, int * y) {
    *x = *x ^ *y;
    *y = *x ^ *y;
    *x = *x ^ *y;
}

int main() {
    int x = 1;
    int y = 2;
    inplace_swap(&x, &y);
    printf("x: %d, y: %d\n", x , y);
}
