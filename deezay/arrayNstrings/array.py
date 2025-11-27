A= [1,2,3]
# Append - Insert element at the end of array - O(1)
A.append(4)
# Pop - Deleting element at end of array - O(1)
A.pop()
# Insert - Insert element at specified index - O (n)
A.insert(2,5)
# Modify an element - o(1)
A[0]=6
# Accessing element given index
print(A[0])
# Checking if array has an element - O(n) same for strings 
if 5 in A:
    print(True)

print(len(A))

s= 'hello'
# Append - 0(n)
b= s+'z'
print(b)


