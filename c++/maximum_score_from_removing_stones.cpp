#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
	int maximumScore(int a, int b, int c) {
		priority_queue<int> q;
		q.push(a);
		q.push(b);
		q.push(c);

		for (int score = 0;; score++) {
			int x = q.top();
			q.pop();
			int y = q.top();
			q.pop();

			if (x == 0 || y == 0)
				return score;

			q.push(x - 1);
			q.push(y - 1);
		}
	}
};
