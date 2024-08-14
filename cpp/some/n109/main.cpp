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

