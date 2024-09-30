<<<<<<< HEAD
#include <iostream>
#include <vector>
#include <map>

using namespace std;

class KthLargest {
    map<int, int> mp;
    int k;

public:

    // 4 5 8 2

    // 2 4 5 8
    // 2 3 4 5 8 = 4
    // 2 3 4 5 5 8 = 5
    // 2 3 4 5 5 8 10 = 5
    // 2 3 4 5 5 8 9 10 = 8
    // ...

    KthLargest(int k, vector<int>& nums) {
        for (size_t i = 0; i < nums.size(); ++i) {
            mp[nums[i]]++;
        }
    }

    int add(int val) {
        mp[val]++;

        int cur_index = 0;
        for (auto rit = mp.rbegin(); rit != mp.rend(); ++rit) {
            if (cur_index + rit->second > k)
                return rit->first;

            cur_index += rit->second;
        }
        return 0;
    }
};


int main()
{
    cout << "Hello World!" << endl;
    return 0;
}
=======
#include <stdio.h>
#include <stdlib.h>

using namespace std;

int main()
{
    int* nums = (int*)malloc(10*sizeof(int));
    if (!nums)
        return -1;

    // ***

    for (int i = 0; i < 10; ++i)
        nums[i] = 10 + i;

    for (int i = 0; i < 10; ++i)
        if (printf("%d ", nums[i]) < 0)
            return -2;

    // ***

    free(nullptr);
    free(nums);
    free(nums); // !

    return 0;
}

>>>>>>> d26a70116fb70142de5c81e6998d458a135dc7b7
