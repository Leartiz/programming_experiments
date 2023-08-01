#include <vector>
#include <iostream>
#include <cassert>

class Solution final
{
public:
    std::vector<std::vector<int>> combinationSum(const std::vector<int>& candidates, int target) {
        std::vector<std::vector<int>> result;
        composeCombinations(result, candidates, target,
                            {}, 0, 0);
        return result;
    }

    void composeCombinations(std::vector<std::vector<int>>& result, const std::vector<int>& candidates,
                             int target, std::vector<int> current, int interimSum, size_t index) {

        for (size_t i = index; i < candidates.size(); ++i) {

            const auto one = candidates[i];
            std::vector<int> next{ current };
            next.push_back(one);

            // ***

            if (interimSum + one == target) {
                result.push_back(next);
            }
            else if (interimSum + one < target) {
                composeCombinations(result, candidates, target,
                        next, interimSum + one, i);
            }
        }
    }
};

void printMx(const std::vector<std::vector<int>>& mx)
{
    for (const auto& line : mx) {
        for (const auto& one : line) {
            std::cout << one << " ";
        } std::cout << std::endl;
    }
}

int main()
{
    Solution s;
    printMx(s.combinationSum({ 1, 2, 3 }, 2));

    {
        const std::vector<std::vector<int>> result = {
            { 2, 2, 3 },
            { 7 }
        };
        const std::vector<std::vector<int>> got = s.combinationSum(
                    { 2, 3, 6, 7 }, 7);
        assert(result == got);
    }
    std::cout << ".";
    {
        const std::vector<std::vector<int>> result = {
            { 2, 2, 2, 2 },
            { 2, 3, 3 },
            { 3, 5 }
        };
        const std::vector<std::vector<int>> got = s.combinationSum(
                    { 2, 3, 5 }, 8);
        assert(result == got);
    }
    std::cout << ".";
    {
        std::vector<std::vector<int>> result = {};
        assert(s.combinationSum({ 2 }, 1) == result);
    }
    std::cout << ".";

    std::cout << "\n[OK]\n";
    return 0;
}
