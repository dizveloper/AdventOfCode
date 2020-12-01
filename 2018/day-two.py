from collections import Counter
from itertools import combinations, compress


twos = 0
threes = 0

for term in open("day-two-input.txt").readlines():
    count = [char for term, char in Counter(term).most_common()]
    if 3 in count:
        twos += 1
    if 2 in count:
        threes += 1

print(twos * threes)

for first_term, sec_term in combinations(open("day-two-input.txt"), 2):
    diff = [char1 == char2 for char1, char2 in zip(first_term, sec_term)]
    if sum(diff) == (len(first_term) - 1):
        box = "".join(list(compress(first_term, diff)))
        break

print(box)
