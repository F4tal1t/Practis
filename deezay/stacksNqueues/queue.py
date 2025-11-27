
# Queues - First in First out (Fifo)

from collections import deque

q = deque()

print(q)

# Append to end of queue - O(1)
q.append(5)
q.append(4)
q.append(3)

print(q)

# Pop from front of queue - O(1)
x = q.popleft()

print(x)
print(q)

# Ask what's on the front of queue - O(1)
print(q[0])