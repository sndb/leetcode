#include <queue>

using namespace std;

class Solution {
public:
	int halveArray(vector<int> &nums) {
		double sum = 0;
		priority_queue<double> q;
		for (int x : nums) {
			q.push(x);
			sum += x;
		}

		int ops = 0;
		double half = sum / 2;
		while (sum > half) {
			double h = q.top() / 2;
			q.pop();
			q.push(h);
			sum -= h;
			ops++;
		}
		return ops;
	}
};
