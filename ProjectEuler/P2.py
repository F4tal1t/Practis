# Even Fibonacci Numbers

a, b = 1, 2
total_sum = 0
while a <=4000000:
    if a % 2 == 0:
        total_sum += a
    a, b = b, a+b
print(total_sum)
