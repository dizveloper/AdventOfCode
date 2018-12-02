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
# total = 0
for line in infile:
    num = line.split()
    nums.append(int(num[0]))
    print(int(num[0]))

total = sum(nums)
print(len(nums))
print(total)

infile.close()
