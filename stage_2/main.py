def fibonacci(limit):
    sequence = [0, 1]
    while True:
        next_num = sequence[-1] + sequence[-2]
        if next_num > limit:
            break
        sequence.append(next_num)
    return sequence

fib = fibonacci(100)
print("Ряд Фибоначчи до 100:")
print(fib)

