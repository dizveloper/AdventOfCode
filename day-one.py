# total = 0

# with open("day-one-input.txt", "r") as inp:
#     for line in inp:
#         try:
#             num = float(line)
#             total += num
#         except ValueError:
#             print("{} is not a number!".format(line))

# print("Total of all numbers: {}".format(total))


infile = open("day-one-input.txt", "r")
infile.readline()  # skip the first line
nums = []
running_total = 0
freqs = [0]
first_repeated = 0

# while first_repeated == 0:
for line in infile:
    num = line.split()
    nums.append(int(num[0]))
    running_total += int(num[0])
    if running_total in freqs:
        first_repeated = running_total
        print(running_total)
    freqs.append(running_total)


print("PART 1: total", running_total)
print("PART 2: first repeat frequency", first_repeated)

infile.close()
