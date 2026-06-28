#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
public:
    vector<vector<int>> flipAndInvertImage(vector<vector<int>> image) {
        const int rows = image.size();
        for (int i = 0; i < rows; ++i) {

            const int cols = image[i].size();
            for (int j = 0; j < cols / 2; ++j) {
                swap(image[i][j], image[i][cols - j]);
                image[i][j] ^= 1;
                image[i][cols - i] ^= 1;
            }
        }
        return image;
    }
};

int main()
{
    vector<vector<int>> image = {
        {1, 1, 0},
        {1, 0, 1},
        {0, 0, 0},
    };

    Solution s;
    s.flipAndInvertImage(image);
    return 0;
}
