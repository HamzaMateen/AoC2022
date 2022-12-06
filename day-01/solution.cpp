/** https://adventofcode.com/2022/day/1 **/
/** Author@HamzaMateen **/

#include <iostream>
#include <vector>
#include <fstream>
#include <algorithm>
#include <numeric>

int main()
{
    std::ifstream file;
    file.open("input.txt");

    std::vector<long int> caloriesPerElf;
    long int max = 0;

    std::string line;
    while (std::getline(file, line))
    {
        if ((int)line[0] == 13) break;
        else max += std::stoi(line);
    }
    caloriesPerElf.push_back(max);
    
    long int temp = 0;
    while (std::getline(file, line))
    {
        if ((int)line[0] == 13)
        {
            caloriesPerElf.push_back(temp);
            temp = 0;
        }
        else temp += std::stoi(line);
    }
    
    std::sort(caloriesPerElf.begin(), caloriesPerElf.end(), [](long int a, long int b) -> bool { return b < a; });
    std::vector<long int> max3(&caloriesPerElf[0], &caloriesPerElf[3]);

    long int maxSum = std::accumulate(max3.begin(), max3.end(), 0);

    std::cout << "Part 1 ans: " << max3[0] << std::endl;
    std::cout << "Part 2 ans: " << maxSum << std::endl;

    file.close();
    return 0;
}
