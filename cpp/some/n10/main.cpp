#include <cassert>
#include <iostream>

#include <unordered_map>
#include <vector>

using namespace std;

class Solution final {
public:
    bool containsNearbyDuplicate(const vector<int>& nums, int k) {
        unordered_map<int, int> mp;
        int n = nums.size();
        mp.reserve(n);

        for(int i = 0; i < n; i++) {
            if(mp.count(nums[i])) {
                if(abs(i - mp[nums[i]]) <= k)
                    return true;
            }

            mp[nums[i]] = i;
        }
        return false;
    }
};

/*

Test cases:

Input: nums = [1,2,3,1], k = 3
Output: true

Input: nums = [1,0,1,1], k = 1
Output: true

Input: nums = [1,2,3,1,2,3], k = 2
Output: false

*/

int main()
{
    Solution s;
    {
        assert(s.containsNearbyDuplicate({ 1, 2, 3, 1 }, 3) == true);
        assert(s.containsNearbyDuplicate({ 1, 0, 1, 1 }, 1) == true);
        assert(s.containsNearbyDuplicate({ 1, 2, 3, 1, 2, 3 }, 2) == false);
    }

    std::cout << "[OK]\n";
    return 0;
}
