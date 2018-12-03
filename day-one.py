import itertools

data = [int(x) for x in open("day-one-input.txt").readlines()]
print(sum(data))

freq = 0
seen = {0}
for num in itertools.cycle(data):
    freq += num
    if freq in seen:
        print(freq)
        break
    seen.add(freq)
