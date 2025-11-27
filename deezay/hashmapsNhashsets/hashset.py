# Hashsets
s = set()

# Add elements - O(1) average
s.add(1)
s.add(2)
print(s)

# Lookup - O(1) average
if 1 in s:
    print(True)

# Remove elements - O(1) average
s.remove(2)
print(s)

# Set Construction from iterable - O(n)
string = "hello"
char_set = set(string)
print(char_set)

# Loop over items in set - O(n)
for x in s:
    print(x)
