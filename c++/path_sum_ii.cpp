#include <bits/stdc++.h>

using namespace std;

struct TreeNode {
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode() : val(0), left(nullptr), right(nullptr) {}
	TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
	TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution {
public:
	vector<vector<int>> pathSum(TreeNode *root, int targetSum) {
		vector<vector<int>> result;
		vector<int> path;
		findPaths(root, targetSum, path, result);
		return result;
	}

	void findPaths(TreeNode *n, int acc, vector<int> &path, vector<vector<int>> &result) {
		if (!n)
			return;
		acc -= n->val;
		path.push_back(n->val);
		if (!n->left && !n->right && acc == 0)
			result.push_back(vector<int>(path));
		findPaths(n->left, acc, path, result);
		findPaths(n->right, acc, path, result);
		path.pop_back();
	}
};
