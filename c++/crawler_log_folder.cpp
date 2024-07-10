#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
	int minOperations(vector<string> &logs) {
		int depth = 0;
		for (string l : logs)
			if (l[1] == '.')
				depth = max(0, depth - 1);
			else if (l[0] != '.')
				depth++;
		return depth;
	}
};
