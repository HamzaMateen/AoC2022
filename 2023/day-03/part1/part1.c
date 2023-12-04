#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#include <ctype.h>

#include <dirent.h>


char *readLine(FILE *file) {

    if (file == NULL)
    {
        printf("Error: file pointer is null.");
        exit(1);
    }

    int maximumLineLength = 128;
    char *lineBuffer = (char *)malloc(sizeof(char) * maximumLineLength);

    if (lineBuffer == NULL)
    {
        printf("Error allocating memory for line buffer.");
        exit(1);
    }

    char ch = getc(file); // read first byte/char
    int count = 0;

    while ((ch != '\n') && (ch != EOF))
    {
        if (count == maximumLineLength)
        {
            maximumLineLength += 128;

            // double the size in case of filled up buffer.
            lineBuffer = realloc(lineBuffer, maximumLineLength);
            if (lineBuffer == NULL)
            {
                printf("Error reallocating space for line buffer.");
                exit(1);
            }
        }

        lineBuffer[count] = ch;
        count++;

        ch = getc(file);
    }

    lineBuffer[count] = '\0'; // size if count + 1 here
    lineBuffer = realloc(lineBuffer, count + 1);

    if (lineBuffer[0] == '\0')
        return NULL;

    return lineBuffer;
}

typedef enum
{
    false,
    true
} bool;


int main(int argc, char const *argv[])
{
    // demo for the first line 
    char* filename = "input.txt";   
    FILE* input = fopen(filename, "r");

    if (input == NULL) {
        perror("error occured while opening the file: ");
        exit(EXIT_FAILURE);
    }

    char* line1 = readLine(input);
    char* line2 = readLine(input);

    printf("%s\n%s\n", line1, line2);

    int enginePartsIdSum = 0;

    int numStart = -1;
    int numEnd = -1;

    // numbers are either 2 or 3 digits long, so we can have max of 3 maybe?
    // 4th for null termination
    char numString[4]; 
    for(int i=0; line1[i] != '\n'; i++) {

        if (isdigit(line1[i])) {
            if (numStart == -1) numStart = i;
            
            strcat(numStart, )
        }
    }


    free(line1);
    free(line2);

    return 0;
}
