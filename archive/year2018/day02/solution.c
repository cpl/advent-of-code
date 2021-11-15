#include <stdlib.h>
#include <stdio.h>


void cmp_by_one(char* str0, char* str1) {
    char* iptr0 = str0;
    char* iptr1 = str1;

    int diff = 0;
    while(*str0 && *str1) {
        if(*str0 != *str1)
            diff++;
        str0++; str1++;
    }
    if(diff == 1) {
        printf("STR: %s\nSTR: %s\n", iptr0, iptr1);

        printf("STR FINAL: ");
        while(*iptr0) {
            if(*iptr0 == *iptr1)
                printf("%c", *iptr0);
            iptr0++; iptr1++;
        }
        printf("\n");
    }
}


int main() {
    // Create file pointer
    FILE* input;

    // Variables
    int mul2x = 0;
    int mul3x = 0;

    char* lines[250];
    int line_index = 0;


    // Open file for reading
    input = fopen("input.txt", "r");
    if(input == NULL) {
        exit(1);
    }

    // Read values while not EOF
    while(!feof(input)) {
        // Read line
        lines[line_index] = malloc(100);
        fscanf(input, "%s\n", lines[line_index]);

        char ascii[0xFF] = {0};

        // Count char frequency
        int index = 0;
        while(lines[line_index][index]) {
            ascii[lines[line_index][index]] += 1;
            index++;
        }
        int seen_2x = 0, seen_3x = 0;
        line_index++;

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
    printf("LINES: %d\n", line_index);

    for(int index_0 = 0; index_0 < 250; index_0++)
        for(int index_1 = index_0+1; index_1 < 250; index_1++)
            cmp_by_one(lines[index_0], lines[index_1]);

    // Close file
    fclose(input);

    return 0;
}