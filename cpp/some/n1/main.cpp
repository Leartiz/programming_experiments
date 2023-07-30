#include <vector>
#include <iostream>
#include <algorithm>
#include <iterator>

void rotate(std::vector<int>& vec, size_t mid)
{
    mid = mid % vec.size();

}

// [1 2 3 4 5 6 7] -> [6 7 1 2 3 4 5]

int main()
{
    std::vector<int> vec{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };
    std::copy(vec.begin(), vec.end(), std::ostream_iterator<int>(std::cout, " "));
    std::cout << std::endl;
    std::rotate(vec.begin(), vec.begin() + 6, vec.end());
    std::copy(vec.begin(), vec.end(), std::ostream_iterator<int>(std::cout, " "));
    std::cout << std::endl;
    return 0;
}
