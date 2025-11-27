# Stacks - Last in First Out (Lifo)

stk = []
print(stk)

# Append to top of stack - O(1)
stk.append(5)
stk.append(4)
stk.append(3)

# Pop from stack - O(1)
x = stk.pop()

print(x)
print(stk)


# Ask what's on the top of stack - O(1)
print(stk[-1])


# Ask if something is in the stack - O(1)
if stk:
  print(True)

