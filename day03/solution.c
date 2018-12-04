#include <stdlib.h>
#include <stdio.h>


#define MATRIX_SIZE 2000
#define MAX_CLAIMS 1375

typedef struct claim {
    int id;
    int x, y;
    int h, w;

    int overlap;
} Claim;


int main() {
    int* matrix = calloc(MATRIX_SIZE*MATRIX_SIZE, sizeof(*matrix));
    Claim claims[MAX_CLAIMS] = {0};

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

        claims[id-1].id = id;
        claims[id-1].x = x;
        claims[id-1].y = y;
        claims[id-1].w = w;
        claims[id-1].h = h;

        for(int delta_y = 0; delta_y < h; delta_y++)
            for(int delta_x = 0; delta_x < w; delta_x++)
                matrix[x+delta_x+(y+delta_y)*MATRIX_SIZE]++;
    }

    int overlap = 0;
    for(int index = 0; index < MATRIX_SIZE*MATRIX_SIZE; index++) {
        if(matrix[index] > 1)
            overlap++;
    }

    for(int claim_index = 0; claim_index < MAX_CLAIMS; claim_index++) {
        Claim c = claims[claim_index];
        int claim_ok = 1;

        for(int delta_y = 0; delta_y < c.h; delta_y++)
            for(int delta_x = 0; delta_x < c.w; delta_x++)
                if(matrix[c.x+delta_x+(c.y+delta_y)*MATRIX_SIZE] != 1)
                    claim_ok = 0;

        if(claim_ok)
            printf("CLAIM ID: %d\n", c.id);
    }


    printf("OVERLAPPING: %d\n", overlap);


    // Close file
    fclose(input);

    return 0;
}