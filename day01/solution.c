#include <stdlib.h>
#include <stdio.h>


int main() {
    // Create file pointer
    FILE* input;

    // Open file for reading
    input = fopen("input.txt", "r");
    if(input == NULL) {
        exit(1);
    }

    int value = 0;
    int tempv = 0;

    // Read values while not EOF
    while(!feof(input)) {
        // Read digit
        fscanf(input, "%d\n", &tempv);
        // Change value
        value += tempv;
    }

    // Print answer
    printf("FINAL VALUE: %d\n", value);

    // Close file
    fclose(input);

    return 0;
}