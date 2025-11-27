# Fibonacci

def fibonacci(n):
    if n <= 0:
        return 0
    elif n == 1:
        return 1
    else:
        return fibonacci(n - 1) + fibonacci(n - 2)
    

# Linked Lists
class Node:
    def __init__(self, data,next=None):
        self.data = data
        self.next = next

    def __str__(self):
        return str(self.data)
head = Node(1)
head.next = Node(2)
head.next.next = Node(3)
def reverse(node):
    if not node:
        return 
    reverse(node.next)
    print(node)

reverse(head) 