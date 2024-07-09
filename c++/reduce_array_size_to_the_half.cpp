#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
	int minSetSize(vector<int> &arr) {
		unordered_map<int, int> occur;
		for (int x : arr)
			occur[x]++;

		priority_queue<int> count;
		for (auto o : occur)
			count.push(o.second);

		int size = 0;
		int total = 0;
		while (total < arr.size() / 2) {
			total += count.top();
			count.pop();
			size++;
		}
		return size;
	}
};
