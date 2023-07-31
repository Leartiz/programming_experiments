#include <vector>
#include <iostream>
#include <algorithm>
#include <iterator>
#include <cassert>

void rotate_impl(std::vector<int>& vec, size_t beg, size_t mid);
void rotate(std::vector<int>& vec, size_t k)
{
    k = k % vec.size();
    rotate_impl(vec, 0, vec.size() - k);
}

void rotate_impl(std::vector<int>& vec, size_t beg, size_t mid)
{
    // why?
    if (beg == mid) return;
    if (vec.size() == mid) return;

    size_t write = beg;
    size_t next_read = beg;

    for (size_t read = mid; read < vec.size(); ++read, ++write) {
        if (write == next_read)
            next_read = read;
        std::swap(vec[write], vec[read]);
    }

    rotate_impl(vec, write, next_read);
}

std::vector<int> rotated(std::vector<int> vec, size_t k)
{
    rotate(vec, k);
    return vec;
}

// Example: [ 1 2 3 4 5 6 |7| 8 9 10 ] -> [ 8 9 10 1 2 3 4 5 6 7 ]

int main()
{
    {
        std::vector<int> vec{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };
        std::copy(vec.begin(), vec.end(), std::ostream_iterator<int>(std::cout, " "));
        std::cout << std::endl;
        std::rotate(vec.begin(), vec.begin() + 7, vec.end());
        std::copy(vec.begin(), vec.end(), std::ostream_iterator<int>(std::cout, " "));
        std::cout << std::endl;
    }
    std::cout << std::endl;
    {
        std::vector<int> vec{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };
        std::copy(vec.begin(), vec.end(), std::ostream_iterator<int>(std::cout, " "));
        std::cout << std::endl;
        rotate(vec, 3);
        std::copy(vec.begin(), vec.end(), std::ostream_iterator<int>(std::cout, " "));
        std::cout << std::endl;
    }

    // tests
    {
        assert(rotated({ 1, 2, 3, 4 ,5 }, 1) == std::vector<int>({ 5, 1, 2, 3, 4 }));
        assert(rotated({ 1, 2, 3, 4 ,5 }, 2) == std::vector<int>({ 4, 5, 1, 2, 3 }));
        assert(rotated({ 1, 2, 3, 4 ,5 }, 4) == std::vector<int>({ 2, 3, 4, 5, 1 }));
        assert(rotated({ 1, 2, 3, 4 ,5 }, 6) == std::vector<int>({ 5, 1, 2, 3, 4 }));

        assert(rotated({ 5, 6, 7, 8, 9, 10, 11 }, 3) == std::vector<int>({ 9, 10, 11, 5, 6, 7, 8 }));
    }

    std::cout << "[OK]\n";
    return 0;
}
