/**
 * @author HamzaMateen
 * @brief solution to Problem No. 1 in C++, day 03, AoC2022 
 * @version 0.1
 * @date 2022-12-08
 * @copyright Copyright (c) 2022
 */

#include <iostream>
#include <set>
#include <fstream>
#include <string>
#include <algorithm>
#include <vector>


int main() {
  std::ifstream file("input.txt");
  std::vector<int> commonElemsInLines;

  if (file.is_open()) {
    std::string line;

    while (std::getline(file, line)) 
    {
      int len = line.length();

      std::string startString = line.substr(0, len/2);
      std::string endString = line.substr(len/2, len);
      
      std::set<char> _1stCompt, _2ndCompt;

      std::for_each(startString.begin(), startString.end(), [&_1stCompt](char c) -> void {
        _1stCompt.emplace(c);
      });

      std::for_each(endString.begin(), endString.end(), [&_2ndCompt](char c) -> void {
        _2ndCompt.emplace(c);
      });

      std::set_intersection(_1stCompt.begin(), _1stCompt.end(), _2ndCompt.begin(), _2ndCompt.end(), std::back_inserter(commonElemsInLines));
    }
  }

  unsigned long priorityItemsSum = 0;

  for (char item : commonElemsInLines) {
    if ((int)item >= 97) 
      priorityItemsSum += item - 'a' + 1;
    else 
      priorityItemsSum += item - 'A' + 27; 
  }

  std::cout << priorityItemsSum << std::endl;

  file.close();
  return 0;
}