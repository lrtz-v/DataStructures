# include <stdio.h>
# include <string.h>

typedef unsigned char *byte_pointer;

void show_bytes(byte_pointer start, int len) {
    int i;
    for (i = 0; i < len; i++) {
        printf(" %0.2x", start[i]);
    }
    printf("\n");
}

void show_int(int x) {
    show_bytes((byte_pointer) &x, sizeof(x));
}

void show_float(float x) {
    show_bytes((byte_pointer) &x, sizeof(float));
}

void show_pointer(void *x) {
    show_bytes((byte_pointer) &x, sizeof(void *));
}

void test_show_bytes(int val) {
    // int ival = val;
    // show_int(ival);
    // show_float((float) ival);
    // int *pval = &ival;
    // show_pointer(pval);
    char *s = "ABCDEF";
    show_bytes(s, strlen("ABCDEF"));
}

int main() {
    test_show_bytes(12345);
    
    return 0;
}
