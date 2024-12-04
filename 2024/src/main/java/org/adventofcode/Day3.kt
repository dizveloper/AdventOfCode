package org.adventofcode

import java.io.File
import kotlin.math.abs

fun Day3() {
    println(" ~ Day 3 ~ ")
    println("Part 1: ${runLogic(false, "src/main/resources/day3_input.txt")}")
    println("Part 2: ${runLogic(true, "src/main/resources/day3_input.txt")}")
}

fun String?.indexesOf(pat: String): List<Int> =
    pat.toRegex()
        .findAll(this?: "")
        .map { it.range.first }
        .toList()

fun runLogic(checkForDoing: Boolean, filename: String): Int {
    var totalMulSum = 0
    var doing = true
    val input = File(filename).readText()
        val lineMulSum = input.indexesOf("mul\\(").map { it: Int ->
            if (checkForDoing) {
                val finalDontInd = input.substring(0, it).lastIndexOf("don't()")
                val finalDoInd = input.substring(0, it).lastIndexOf("do()")
                doing = finalDontInd <= finalDoInd
            }

            val commaInd = input.indexOf(",", it)
            val endInd = input.indexOf(")", it) + 1
            var res = 0

            if (doing) {
                if (input.substring(it + 4, commaInd).toIntOrNull() != null
                    && input.substring(commaInd + 1, endInd - 1).toIntOrNull() != null
                ) {
                    res = input.substring(it + 4, commaInd).toInt() * input.substring(commaInd + 1, endInd - 1).toInt()
                }
            }
            res
        }.sum()

        totalMulSum += lineMulSum

    return totalMulSum
}