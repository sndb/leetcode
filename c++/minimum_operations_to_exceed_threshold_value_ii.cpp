#include <queue>

using namespace std;

class Solution {
public:
	int minOperations(vector<int> &nums, int k) {
		priority_queue<long, vector<long>, greater<long>> q(nums.begin(), nums.end());
		int c = 0;
		while (q.top() < k) {
			long x = q.top();
			q.pop();
			long y = q.top();
			q.pop();

			q.push(min(x, y) * 2 + max(x, y));
			c++;
		}
		return c;
	}
};
