#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
	vector<vector<int>> kClosest(vector<vector<int>> &points, int k) {
		auto cmp = [](vector<int> &a, vector<int> &b) { return dist(a) < dist(b); };
		nth_element(points.begin(), points.begin() + k, points.end(), cmp);
		return vector<vector<int>>(points.begin(), points.begin() + k);
	}

	static double dist(vector<int> &a) {
		return a[0] * a[0] + a[1] * a[1];
	}
};
