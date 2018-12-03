#include <stdlib.h>
#include <stdio.h>


int main() {
    // Create file pointer
    FILE* input;

    // Variables
    char buffer[200];
    int mul2x = 0, mul3x = 0;

    // Open file for reading
    input = fopen("input.txt", "r");
    if(input == NULL) {
        exit(1);
    }

    // Read values while not EOF
    while(!feof(input)) {
        // Read line
        fscanf(input, "%s\n", buffer);
        char ascii[0xFF] = {0};

        // Count char frequency
        int index = 0;
        while(buffer[index]) {
            ascii[buffer[index]] += 1;
            index++;
        }
        int seen_2x = 0, seen_3x = 0;

        // Check for 2x or 3x
        for(int a_index = 'a'; a_index <= 'z'; a_index++)
            switch(ascii[a_index]) {
            case 2:
                if(!seen_2x) {
                    mul2x++;
                    seen_2x = 1;
                }
                break;
            case 3:
                if(!seen_3x) {
                    mul3x++;
                    seen_3x = 1;
                }
                break;
            }
    }

    printf("RESULT: %d x %d = %d\n", mul2x, mul3x, mul2x*mul3x);

    // Close file
    fclose(input);

    return 0;
}