#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
	double minimumAverage(vector<int> &nums) {
		sort(nums.begin(), nums.end());
		double r = numeric_limits<double>::max();
		for (int i = 0; i < nums.size() / 2; i++)
			r = min(r, (nums[i] + nums[nums.size() - 1 - i]) / 2.0);
		return r;
	}
};
