#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#include <ctype.h>

enum config
{
    RED = 12,
    GREEN,
    BLUE,
};

typedef enum
{
    false,
    true
} bool;

struct Cube
{
    char *color;
    int quantity;
};

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

char* my_strdup(char* src) {
    if (src == NULL) 
        return NULL;

    char* dest = (char*)malloc(strlen(src) + 1);

    if (dest == NULL) {
        perror("error allocating space for 'dest' string");
        exit(EXIT_FAILURE);
    }

    strcpy(dest, src);
    return dest; 
}

int main(int argc, char const *argv[])
{
    int idSum = 0;
    int gameId, redCount, greenCount, blueCount;
    int currGameDrawCount;
    bool isGamePossible;

    // assume max size is 10
    const int MAX_DRAWS = 10;
    char *draws[MAX_DRAWS];

    // there are maximum of 3 types of cubes per draw
    char *cubeInfos[3] = {NULL, NULL, NULL};

    struct Cube cubes[3]; // some string representational error
    struct Cube tempCube;

    // init to NULL
    for (int i = 0; i < MAX_DRAWS; i++)
    {
        draws[i] = NULL;
    }

    // let's try read the input file
    char *filename = "input.txt";
    FILE *input = fopen(filename, "r");

    char *line;
    while ((line = readLine(input)) != NULL)
    {
        // %s expects a null terminated string
        // printf("%s\n", line);

        // split the line by :
        char *tokenStart = strtok(line, ":");
        char *tokenEnd = strtok(NULL, ":");

        strtok(tokenStart, " ");
        gameId = atoi(strtok(NULL, " "));

        // separate the part2 by `;` in order to get the ball drawn and its quantity for 3 consecutive times
        char *currDraw;

        currDraw = strtok(tokenEnd, ";");
        draws[0] = &currDraw[1]; // strip the front space character
        for (int i = 1; (currDraw = strtok(NULL, ";")) != NULL; i++)
        {
            draws[i] = &currDraw[1];
        }

        isGamePossible = true;
        // char *cubeInfo;
        for (int i = 0; draws[i] != NULL; i++)
        {
            // for 3 cube types
            // we have an array of cubes we can fill
            // char *initialDrawInfo = my_strdup(draws[i]);

            cubeInfos[0] = strtok(draws[i], ",");
            for (int idx = 1; (cubeInfos[idx] = strtok(NULL, ",")) != NULL; idx++) {
                printf("%s\n", cubeInfos[idx-1]);
            };

            printf("\n\n");

            // reset the cubeInfos 
            cubeInfos[0] = cubeInfos[1] = cubeInfos[2] = NULL;

            // tempCube.quantity = atoi(strtok(cubeInfo, " "));
            // tempCube.color = strtok(NULL, " ");

            // cubeInfo = strtok(initialDrawInfo, ",");
            // free(initialDrawInfo);

            // for (int j = 1; (cubeInfo = strtok(NULL, ",")) != NULL; j++)
            // {
            //     printf("CUBE INFO: %s\n", cubeInfo);
            //     tempCube.quantity = atoi(strtok(cubeInfo, " "));
            //     tempCube.color = strtok(NULL, " ");

            //     cubes[j] = tempCube;
            // }
            
            // for (int k = 0; k < 3; k++)
            // {
            //     if (strcmp(cubes[k].color, "red") == 0 && cubes[k].quantity > RED ||
            //         strcmp(cubes[k].color, "green") == 0 && cubes[k].quantity > GREEN ||
            //         strcmp(cubes[k].color, "blue") == 0 && cubes[k].quantity > BLUE)
            //     {
            //         isGamePossible = false;
            //     }
            // }
        }

        if (isGamePossible)
        {
            idSum += gameId;
        }

        // reset the cubes
        // cubes[0].color = cubes[1].color = cubes[2].color = "";
        // cubes[1].quantity = cubes[1].quantity = cubes[1].quantity = 0;

        // init to NULL again
        for (int i = 0; i < MAX_DRAWS; i++)
        {
            draws[i] = NULL;
        }

        // line is dynamically allocated by the readLine function.
        free(line);
    }

    printf("Sum of possible game's ids: %d\n", idSum);

    return 0;
}
