#include <stdlib.h>
#include <stdio.h>


#define MATRIX_SIZE 2000


int main() {
    int* matrix = calloc(MATRIX_SIZE*MATRIX_SIZE, sizeof(int));

    // Create file pointer
    FILE* input;

    // Open file for reading
    input = fopen("input.txt", "r");
    if(input == NULL) {
        exit(1);
    }


    // Read values while not EOF
    while(!feof(input)) {
        int id, x, y, h, w;
        fscanf(input, "#%d @ %d,%d: %dx%d\n", &id, &x, &y, &w, &h);
        for(int delta_y = 0; delta_y < h; delta_y++)
            for(int delta_x = 0; delta_x < w; delta_x++)
                matrix[x+delta_x+(y+delta_y)*MATRIX_SIZE]++;
    }

    int overlap = 0;
    for(int index = 0; index < MATRIX_SIZE*MATRIX_SIZE; index++) {
        if(matrix[index] > 1)
            overlap++;
    }

    printf("OVERLAPPING: %d\n", overlap);


    // Close file
    fclose(input);

    return 0;
}