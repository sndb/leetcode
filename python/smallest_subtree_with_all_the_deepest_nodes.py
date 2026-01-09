class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def subtreeWithAllDeepest(self, root: TreeNode) -> TreeNode:
        def search(node):
            if node is None:
                return (None, 0)

            left_lca, left_depth = search(node.left)
            right_lca, right_depth = search(node.right)

            if left_depth > right_depth:
                return (left_lca, left_depth + 1)
            elif left_depth < right_depth:
                return (right_lca, right_depth + 1)
            else:
                return (node, left_depth + 1)

        return search(root)[0]
