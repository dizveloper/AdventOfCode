package org.adventofcode

import java.io.File
import kotlin.math.abs

fun Day2() {
    println(" ~ Day 2 ~ ")

    val lines = List<List<Int>>(0) { emptyList() }.toMutableList()
    // Part 1
    File("src/main/resources/day2_input.txt").forEachLine { line ->
        val lineAsInt = line.split(" ").map { it.toInt() }
        lines += listOf(lineAsInt)
    }

    var safeTotal = 0
    safeTotal = lines.map { it }
        .count {
            isSafe(it)
        }

    println("Part 1: $safeTotal")

    // Part 2
    safeTotal = lines.map { it }
        .count {
            isSafe(it) || it.indices.any { index ->
                isSafe(it.subList(0, index) + it.subList(index + 1, it.size))
            }
        }
    println("Part 2: $safeTotal")

}

fun isSafe(list: List<Int>): Boolean {
    val direction = list[0] < list[1]

    for (i in 0 until list.size - 1) {
        val currentDifference = abs(list[i] - list[i + 1])
        if ((currentDifference > 3 || currentDifference == 0) || (list[i] < list[i + 1] != direction)) {
            return false
        }
    }

    return true
}


