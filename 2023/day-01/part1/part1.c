#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#include <ctype.h>

typedef enum {
    false, true
} bool;


int main(int argc, char const *argv[])
{
    int sum = 0;
    char firstDigit, lastDigit;
    char resultNum[2] = {'\n', '\n'}; // placeholders for nil or no
    
    // open the file 
    int fd = open("input.txt", O_RDONLY);

    if (fd <= 0) {
        perror("Error opening input file: ");
        exit(EXIT_FAILURE);
    }

    // read file line by line 
    char c;
    int bytesRead;

    bool isFirstDigit = true;
    while ((bytesRead = read(fd, &c, sizeof(char))) > 0) {
        if ('\n' == c) {
            // check if there was only 1 number per string 
            if (lastDigit == '\n') 
                lastDigit = firstDigit;

            resultNum[0] = firstDigit;
            resultNum[1] = lastDigit;

            sum += atoi(resultNum);
            isFirstDigit = true;

            // reset digits to 0 
            firstDigit = lastDigit = '\n';
            continue;
        }

        if (isdigit(c)) {
            if(isFirstDigit) {
                firstDigit = c;
                isFirstDigit = false;
            }
            else {
                lastDigit = c;
            }
        }
    }

    // close the opened file 
    close(fd);

    // print the sum value
    printf("Calibrations' sum is: %d\n", sum);


    return 0;
}
