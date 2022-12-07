total_score = 0

with open("input.txt") as f:
  choice_combo_1 = list(map(lambda e : (e, 'X'), ['A', 'B', 'C']));
  choice_combo_2 = list(map(lambda e : (e, 'Y'), ['A', 'B', 'C']));
  choice_combo_3 = list(map(lambda e : (e, 'Z'), ['A', 'B', 'C']));
  
  for line in f:
    current_choice = tuple(line.strip("\n").split(" "))

    if (current_choice == choice_combo_1[0]):
      total_score += 1 + 3 # rock and draw
    elif (current_choice == choice_combo_1[1]):
      total_score += 1 + 0 # rock and lose
    elif (current_choice == choice_combo_1[2]):
      total_score += 1 + 6 # rock and win 
    elif (current_choice == choice_combo_2[0]):
      total_score += 2 + 6 # paper and win 
    elif (current_choice == choice_combo_2[1]):
      total_score += 2 + 3 # paper and draw
    elif (current_choice == choice_combo_2[2]):
      total_score += 2 + 0 # paper and lose 
    elif (current_choice == choice_combo_3[0]):
      total_score += 3 + 0 # scissor and lose
    elif (current_choice == choice_combo_3[1]):
      total_score += 3 + 6 # scissor and win 
    else:
      total_score += 3 + 3

  print(total_score)
  # NOTE: I didn't have python3.10 at the time of implementation so used if-else, match-case is recommended otherwise.
