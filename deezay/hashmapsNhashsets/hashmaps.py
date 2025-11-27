#  Hashmaps - Dictionaries in Python

d = {'a': 1, 'b': 2, 'c': 3}
print(d)

# Add key-value pair - O(1) average
d['d'] = 4

# Check if key exists - O(1) average
if 'a' in d:
    print(True)

# Loop over the key:val pairs of the dictionary - O(n)
for key, val in d.items():
    print(key, val)
    # a 1
    # b 2
    # c 3
    # d 4

# Remove key-value pair - O(1) average
del d['d']

#Defaultdict from collections module

from collections import defaultdict

default = defaultdict(list)  # default value of int is 0

default[2]


# Counter from collections module
from collections import Counter
count = Counter("hello world")
print(count)  
# Counter({'l': 3, 'o': 2, 'h': 1, 'e': 1, ' ': 1, 'w': 1, 'r': 1, 'd': 1})