width, height = 1000, 1000

import json

input
with open("day-three-input.json") as json_data:
    input = json.load(json_data)

Quilt = [[0 for x in range(width)] for y in range(height)]

for inner_line in input:
    coordinates = inner_line["coordinates"]
    inner_x = coordinates["x"]
    inner_y = coordinates["y"]

    size = inner_line["size"]
    inner_w = size["width"]
    inner_h = size["height"]

    for row in range(inner_h):
        for col in range(inner_w):
            Quilt[inner_y + row][inner_x + col] = (
                Quilt[inner_y + row][inner_x + col]
            ) + 1


union = 0
for Fabric in Quilt:
    for inch in Fabric:
        if inch >= 2:
            union += 1
print(union)


for inner_line in input:
    coordinates = inner_line["coordinates"]
    inner_x = coordinates["x"]
    inner_y = coordinates["y"]

    size = inner_line["size"]
    inner_w = size["width"]
    inner_h = size["height"]

    failed = False
    for row in range(inner_h):
        for col in range(inner_w):
            if Quilt[inner_y + row][inner_x + col] != 1:
                failed = True
                break
    if not failed:
        print(inner_line["id"])
