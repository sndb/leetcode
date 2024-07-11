#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
	string reverseParentheses(string s) {
		vector<char> a;
		stack<int> v;
		int j = 0;
		for (auto c : s) {
			if (c == '(') {
				v.push(j);
			} else if (c == ')') {
				reverse(a.begin() + v.top(), a.end());
				v.pop();
			} else {
				a.push_back(c);
				j++;
			}
		}
		return string(a.begin(), a.end());
	}
};
