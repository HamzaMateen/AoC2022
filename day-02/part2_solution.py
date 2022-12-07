total_score = 0

ROCK = 1
PAPER = 2
SCISSORS = 3

with open("input.txt") as f:
  
  for line in f:  
    oppChoice, result = line.strip("\n").split(" ")
    if (oppChoice == 'A' and result == 'X'):
      total_score += SCISSORS 
    elif (oppChoice == 'A' and result == 'Y'):
      total_score += ROCK + 3
    elif (oppChoice == 'A' and result == 'Z'):
      total_score += PAPER + 6
    elif (oppChoice == 'B' and result == 'X'):
      total_score += ROCK 
    elif (oppChoice == 'B' and result == 'Y'):
      total_score += PAPER + 3
    elif (oppChoice == 'B' and result == 'Z'):
      total_score += SCISSORS + 6
    elif (oppChoice == 'C' and result == 'X'):
      total_score += PAPER 
    elif (oppChoice == 'C' and result == 'Y'):
      total_score += SCISSORS + 3
    elif (oppChoice == 'C' and result == 'Z'):
      total_score += ROCK + 6
      
print(total_score)
    