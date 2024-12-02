package org.adventofcode

import java.io.File
import kotlin.math.abs

fun Day1() {
    println(" ~ Day 1 ~ ")

    // Part 1
    val left = ArrayList<Int>()
    val right = ArrayList<Int>()

    val input = File("src/main/resources/day1_input.txt").forEachLine { line ->
        val line = line.split("   ").toTypedArray()
        left.add(line[0].toInt())
        right.add(line[1].toInt())
    }

    // sort both lists in ascending order
    left.sort()
    right.sort()

    var diffTotal = 0
    for (i in left.indices) {
        val diff = abs(left[i] - right[i])

        diffTotal += diff
    }

    println("Part 1: $diffTotal")

    // Part 2
    val similarityScores = ArrayList<Int>()
    for (i in left.indices) {
        var appearances = 0
        for (j in right.indices) {
            if (left[i] == right[j]) {
                appearances++
            }
        }

        similarityScores.add(left[i]*appearances)
    }

    println("Part 2: ${similarityScores.sum()}")
}
