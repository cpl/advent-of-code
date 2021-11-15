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


    // Read values while not EOF
    while(!feof(input)) {

    }

    // Close file
    fclose(input);

    return 0;
}