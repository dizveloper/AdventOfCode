import json
import datetime
import re

input
with open("day-four-sorted-input.json") as json_data:
    input = json.load(json_data)
GuardSchedule = input

with open("day-four-input.txt") as data:
    GuardIDs = data.read()
# GuardSchedule.sort(
#     key=lambda x: (
#         datetime.datetime.strptime(x["date"], "%Y-%m-%d"),
#         datetime.datetime.strptime(x["time"], "%H:%M"),
#     )
# )
# with open("day-four-sorted-input.json", "w") as f:
#     json.dump(GuardSchedule, f)

AllGuards = re.findall(r"\#(.[0-9]+)", GuardIDs)
TimeChart = [["." for x in range(120)] for y in range(len(AllGuards))]

AsleepTimeChart = [[0 for x in range(2)] for y in range(len(AllGuards))]

for i, id in enumerate(AllGuards):
    AsleepTimeChart[i][0] = int(id)
    TimeChart[i][0] = int(id)

index = 0
current_guard_index = 0
sleep_check_index = 0
for event in GuardSchedule:

    if re.search(r"\#(.[0-9]+)", event["happening"]):
        start_shift = re.search(r"\#(.[0-9]+)", event["happening"])
        current_guard = start_shift.group(1)
        sleep_check_index = 1
        # This needs  to change to be any with current_guard
        print(
            TimeChart.index(
                [
                    int(current_guard),
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                    ".",
                ]
            )
        )
    #     current_guard_index = # index where key = current_guard

    if re.search(r"asleep*", event["happening"]):
        TimeChart[current_guard_index][sleep_check_index] += 1

    index += 1
    sleep_check_index += 1

# print(TimeChart)

# fill AsleepTimeChart[1] with totals of each TimeChart[guard_id] ignoring the first
