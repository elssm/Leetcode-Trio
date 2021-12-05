#include <iostream>
#include <vector>

using std::vector;
using std::swap;

/*
 * vector<int>下标应该定义为std::vector<int>::size_type，或使用decltype(nums.size())来推断
 * 但刷题为了省事，均为int
 */

class Solution {
public:
    // 1. 暴力法（开了新空间，不是原地操作）
    void moveZeroes1(vector<int>& nums) {
        vector<int> non_0;  // 存放非0元素

        // 遍历原数组，把非0元素存到non_0中
        for (auto i : nums) if (i) non_0.push_back(i);
        // 把非0元素存回原数组的前面
        for (int i = 0; i != non_0.size(); ++i)
            nums[i] = non_0[i];
        // 从非0元素的下一位开始补0
        for (int i = non_0.size(); i != nums.size(); ++i)
            nums[i] = 0;
    }

    // 2. 不开新空间
    void moveZeroes2(vector<int>& nums) {
        int k = 0;  // 指向非0元素该存的位置，从0开始
        for (int i = 0; i < nums.size(); ++i) {
            // i指针往后寻找非0元素
            if (nums[i]) {
                if (i != k) nums[k++] = nums[i];    // 将非零元素复制到k指向的位置，k++
                else ++k;   // 小优化：如果i找到非零元素，但此时i，k同时指向此处，说明该非零元素已就位，无需原地复制
            }
        }

        // 将后面的补0
        for (; k != nums.size();) nums[k++] = 0;
    }

    // 3. 在2的基础上，直接交换，无需补0操作
    void moveZeroes3(vector<int>& nums) {
        int k = 0;
        for (int i = 0; i < nums.size(); ++i) {
            if (nums[i]) {
                if (i != k) swap(nums[k++], nums[i]);
                else ++k;
            }
        }
    }
};

void print_vector(vector<int>& nums) {
    for (auto i : nums)
        std::cout << i << " ";
    std::cout << std::endl;
}

// 测试
int main() {
    vector<int> v1 = {0, 1, 0, 3, 12};
    // Solution().moveZeroes1(v1);
    // Solution().moveZeroes2(v1);
    Solution().moveZeroes3(v1);
    print_vector(v1);

    return 0;
}