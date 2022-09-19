#include <stdbool.h>

struct TreeNode {
	int              val;
	struct TreeNode *left;
	struct TreeNode *right;
};

bool
gt(struct TreeNode *root, int val) {
	return !root || root->val > val && gt(root->left, val) && gt(root->right, val);
}

bool
lt(struct TreeNode *root, int val) {
	return !root || root->val < val && lt(root->left, val) && lt(root->right, val);
}

bool
isValidBST(struct TreeNode *root) {
	if (!root)
		return true;
	return isValidBST(root->left)
	    && isValidBST(root->right)
	    && lt(root->left, root->val)
	    && gt(root->right, root->val);
}
