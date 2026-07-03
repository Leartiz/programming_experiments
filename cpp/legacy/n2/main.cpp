#include <vector>
#include <algorithm>
#include <iterator>
#include <iostream>

using namespace std;

void print_vec(const vector<int>& vec)
{
    copy(vec.begin(), vec.end(),
         ostream_iterator<int>(cout, " "));
    cout << endl;
}

int main()
{
    vector vec{1, 2, 3, 4};
    print_vec(vec);
    iter_swap(vec.begin() + 0, vec.begin() + 1); // swap vals.
    print_vec(vec);
    iter_swap(vec.begin() + 1, vec.begin() + 3);
    print_vec(vec);
    return 0;
}
