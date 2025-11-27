# Binary Trees
class TreeNode:
  def __init__(self, val, left=None, right=None):
    self.val = val
    self.left = left
    self.right = right

  def __str__(self):
    return str(self.val)

#        1
#     2    3
#   4  5  10

A = TreeNode(1)
B = TreeNode(2)
C = TreeNode(3)
D = TreeNode(4)
E = TreeNode(5)
F = TreeNode(10)

A.left = B
A.right = C
B.left = D
B.right = E
C.left = F

print(A)
# 1


# Recursive Pre Order Traversal (DFS) Time: O(n), Space: O(n)
def pre_order(node):
  if not node:
    return

  print(node)
  pre_order(node.left)
  pre_order(node.right)

pre_order(A)

# 1
# 2
# 4
# 5
# 3
# 10

# Recursive In Order Traversal (DFS) Time: O(n), Space: O(n)
def in_order(node):
  if not node:
    return

  in_order(node.left)
  print(node)
  in_order(node.right)

in_order(A)
# 4
# 2
# 5
# 1
# 10
# 3

# Iterative Pre Order Traversal (DFS) Time: O(n), Space: O(n)
def pre_order_iterative(node):
  stk = [node]

  while stk:
    node = stk.pop()
    if node.right: stk.append(node.right)
    if node.left: stk.append(node.left)
    print(node)

pre_order_iterative(A)

# 1
# 2
# 4
# 5
# 3
# 10


# Level Order Traversal (BFS) Time: O(n), Space: O(n)
from collections import deque

def level_order(node):
  q = deque()
  q.append(node)

  while q:
    node = q.popleft()
    print(node)
    if node.left: q.append(node.left)
    if node.right: q.append(node.right)

level_order(A)
# 1
# 2
# 3
# 4
# 5
# 10

# Check if Value Exists (DFS) Time: O(n), Space: O(n)
def search(node, target):
  if not node:
    return False

  if node.val == target:
    return True

  return search(node.left, target) or search(node.right, target)

search(A, 11)
# False