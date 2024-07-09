#include <queue>

using namespace std;

class Solution {
public:
	int findKthLargest(vector<int> &nums, int k) {
		priority_queue q(nums.begin(), nums.end());
		while (--k > 0) q.pop();
		return q.top();
	}
};
