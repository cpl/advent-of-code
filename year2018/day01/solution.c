#include <stdlib.h>
#include <stdio.h>


int main() {
    // Create file pointer
    FILE* input;

    // Create frequency array
    int seen[2000000] = {0};

    // Variables
    int value = 0;
    int tempv = 0;
    int has_seen = 0;

    // Open file for reading
    input = fopen("input.txt", "r");
    if(input == NULL) {
        exit(1);
    }

    // Read values until a frequency is seen twice
    while(!has_seen) {
        rewind(input);

        // Read values while not EOF
        while(!feof(input)) {

            // Increment and check if frequency was met before
            seen[value+1000000] += 1;
            if(!has_seen)
                if(seen[value+1000000] > 1) {
                    printf("SEEN VALUE %d\n", value);
                    has_seen = 1;
                }

            // Read digit
            fscanf(input, "%d\n", &tempv);
            // Change value
            value += tempv;
        }
    }

    // Print answer
    printf("FINAL VALUE: %d\n", value);

    // Close file
    fclose(input);

    // for(int index = 0; index < 2000000; index++)
    //     if(seen[index] != 0)
    //         printf("%d ", seen[index]);

    return 0;
}