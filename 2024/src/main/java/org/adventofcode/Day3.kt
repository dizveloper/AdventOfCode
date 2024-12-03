package org.adventofcode

import java.io.File
import kotlin.math.abs

fun Day3() {
    println(" ~ Day 3 ~ ")

    println("Part 1: ${runLogic(false)}")
    println("Part 2: ${runLogic(true)}")
}

fun String?.indexesOf(pat: String): List<Int> =
    pat.toRegex()
        .findAll(this?: "")
        .map { it.range.first }
        .toList()

fun runLogic(checkForDoing: Boolean): Int {
    var totalMulSum = 0
    File("src/main/resources/day3_input_ex2.txt").forEachLine { line ->
        var doing = true
        val lineMulSum = line.indexesOf("mul\\(").map { it: Int ->
            if (checkForDoing) {
                // Look backwards to see if there are any instances of don't() before the current mul()
                val finalDontInd = line.substring(0, it).lastIndexOf("don't()")
                val finalDoInd = line.substring(0, it).lastIndexOf("do()")
                doing = finalDontInd <= finalDoInd
            }

            val commaInd = line.indexOf(",", it)
            val endInd = line.indexOf(")", it) + 1
            var res = 0

            // do if it is within an index window in doWindows
            if (doing) {
                if (line.substring(it + 4, commaInd).toIntOrNull() != null
                    && line.substring(commaInd + 1, endInd - 1).toIntOrNull() != null
                ) {
                    val mulStatement = line.substring(it, endInd)

                    res = line.substring(it + 4, commaInd).toInt() * line.substring(commaInd + 1, endInd - 1).toInt()
                }
            }
            res
        }.sum()

        totalMulSum += lineMulSum
    }

    return totalMulSum
}