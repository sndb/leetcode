from math import inf

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def maxLevelSum(self, root: TreeNode | None) -> int:
        max_sum = -inf
        max_level = 1
        curr_level = 1
        queue = [root]
        while queue:
            curr_sum = 0
            next_level = []
            for node in queue:
                curr_sum += node.val
                if node.left is not None:
                    next_level.append(node.left)
                if node.right is not None:
                    next_level.append(node.right)

            if curr_sum > max_sum:
                max_level = curr_level
                max_sum = curr_sum

            curr_level += 1
            queue = next_level

        return max_level
