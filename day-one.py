import itertools

input = [int(x) for x in open("day-one-input.txt").readlines()]
print(sum(input))

first_repeated_freq = 0
seen = {0}
for num in itertools.cycle(input):
    first_repeated_freq += num
    if first_repeated_freq in seen:
        print(first_repeated_freq)
        break
    seen.add(first_repeated_freq)
