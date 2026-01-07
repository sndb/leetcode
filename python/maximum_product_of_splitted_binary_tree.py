class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def maxProduct(self, root: TreeNode | None) -> int:
        sums = []
        def walk(node):
            if node is None:
                return 0

            sum = node.val + walk(node.left) + walk(node.right)
            sums.append(sum)
            return sum

        total = walk(root)
        return max(sum * (total - sum) for sum in sums) % (10**9 + 7)
