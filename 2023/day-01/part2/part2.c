#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#include <ctype.h>

typedef enum {
    false, true
} bool;

// let's get things done!
char textToNumber(char* numInText);

int main(int argc, char const *argv[])
{
    char firstDigit, lastDigit;
    char resultNum[2] = {'\n', '\n'};

    char buf1[3];
    char buf2[4];
    char buf3[5];

    char c;
    int br; 
    off_t seekBack;

    int fd = open("input.txt", O_RDONLY);
    
    int sum = 0;
    bool isFirstDigit = true;

    while ((br = read(fd, &c, sizeof(char))) > 0) {
        // 
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

        seekBack = lseek(fd, -1, SEEK_CUR);
        if (isdigit(c)) {
            if(isFirstDigit) {
                firstDigit = c;
                isFirstDigit = false;
            }
            else {
                lastDigit = c;
            }

            lseek(fd, 1, SEEK_CUR);
            continue;
        }

        if (isalpha(c)) {
            // seek forward 3 chars because there can be 
            // one, two, six
            // use the buf1 (3 characters)
            br = read(fd, buf1, 3);
            seekBack = lseek(fd, -3, SEEK_CUR);

            
            // seek forward 4 chars because there can be 
            // four, five, nine
            // use the buf2 (4 characters)
            br = read(fd, buf2, 4);
            seekBack = lseek(fd, -4, SEEK_CUR);
            

            // seek forward 5 chars because there can be 
            // three, seven, eight
            // use the buf3 (5 characters)
            br = read(fd, buf3, 5);
            seekBack = lseek(fd, -5, SEEK_CUR);

            // after having read all of these values
            // there can be at most one value at a time 
            // and none at least 

            // seek forward per character
            lseek(fd, 1, SEEK_CUR);

            c = textToNumber(buf1); 
            if (c != '\0') {
                if(isFirstDigit) {
                    firstDigit = c;
                    isFirstDigit = false;
                }
                else 
                    lastDigit = c;

                continue;
            } 

            c = textToNumber(buf2);
            if (c != '\0') {
                if(isFirstDigit) {
                    firstDigit = c;
                    isFirstDigit = false;
                }
                else 
                    lastDigit = c;

                continue;
            } 

            // loop should end when the last two bytes of buf3 are 
            // both \n because the 2nd \n would indicated an empty line 
            if ('\n' == buf3[4]) {
                if ('\n' == buf3[3]) 
                    break;
                else 
                    continue;
            }
            
            c = textToNumber(buf3);
            if (c != '\0') {
                if(isFirstDigit) {
                    firstDigit = c;
                    isFirstDigit = false;
                }
                else 
                    lastDigit = c;

                continue;
            } 
        }
    }


    printf("The sum is %d", sum);
}

char textToNumber(char* numInText) {
    if (strncmp(numInText, "one", 3) == 0) {
        return '1';
    } 
    if (strncmp(numInText, "two", 3) == 0) {
        return '2';
    } 
    if (strncmp(numInText, "three", 5) == 0) {
        return '3';
    } 
    if (strncmp(numInText, "four", 4) == 0) {
        return '4';
    } 
    if (strncmp(numInText, "five", 4) == 0) {
        return '5';
    } 
    if (strncmp(numInText, "six", 3) == 0) {
        return '6';
    } 
    if (strncmp(numInText, "seven", 5) == 0) {
        return '7';
    } 
    if (strncmp(numInText, "eight", 5) == 0) {
        return '8';
    } 
    if (strncmp(numInText, "nine", 4) == 0) {
        return '9';
    } 
    return '\0';
}
