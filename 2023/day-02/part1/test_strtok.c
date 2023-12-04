#include <stdio.h>
#include <string.h>

int
main ()
{
  printf ("Hello World");
  char myString[8] = " 3 blue"; // strtok won't be able to work here!

  char *firstPart = strtok (myString, " ");
  char *lastPart = strtok (NULL, " ");

  printf("%s:%s\n", firstPart, lastPart);
  return 0;
}
