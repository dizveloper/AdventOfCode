package org.adventofcode

import java.io.File
import kotlin.math.abs

fun Day4() {
    println(" ~ Day 4 ~ ")

    val crosswordPuzzle = mutableMapOf<Int, List<Char>>()
    File("src/main/resources/day4_input.txt").readLines().forEachIndexed() { row, line ->
        crosswordPuzzle[row] = line.toList()
    }

    // Part 1
    var foundXmas = 0
    crosswordPuzzle.forEach { (row, line) ->
        line.forEachIndexed { col, char ->
            if (char == 'X') {
                // Check forwards
                if (col + 3 < line.size) {
                    if (line[col + 1] == 'M' && line[col + 2] == 'A' && line[col + 3] == 'S') {
                        foundXmas++
                    }
                }
                // Check backwards
                if (col - 3 >= 0) {
                    if (line[col - 1] == 'M' && line[col - 2] == 'A' && line[col - 3] == 'S') {
                        foundXmas++
                    }
                }
                // Check above
                if (row - 3 >= 0) {
                    if (crosswordPuzzle[row - 1]!![col] == 'M' && crosswordPuzzle[row - 2]!![col] == 'A' && crosswordPuzzle[row - 3]!![col] == 'S') {
                        foundXmas++
                    }
                }
                // Check below
                if (row + 3 < crosswordPuzzle.size) {
                    if (crosswordPuzzle[row + 1]!![col] == 'M' && crosswordPuzzle[row + 2]!![col] == 'A' && crosswordPuzzle[row + 3]!![col] == 'S') {
                        foundXmas++
                    }
                }
                // Check diagonally
                if (row - 3 >= 0 && col + 3 < line.size) {
                    if (crosswordPuzzle[row - 1]!![col + 1] == 'M' && crosswordPuzzle[row - 2]!![col + 2] == 'A' && crosswordPuzzle[row - 3]!![col + 3] == 'S') {
                        foundXmas++
                    }
                }
                if (row + 3 < crosswordPuzzle.size && col + 3 < line.size) {
                    if (crosswordPuzzle[row + 1]!![col + 1] == 'M' && crosswordPuzzle[row + 2]!![col + 2] == 'A' && crosswordPuzzle[row + 3]!![col + 3] == 'S') {
                        foundXmas++
                    }
                }
                if (row - 3 >= 0 && col - 3 >= 0) {
                    if (crosswordPuzzle[row - 1]!![col - 1] == 'M' && crosswordPuzzle[row - 2]!![col - 2] == 'A' && crosswordPuzzle[row - 3]!![col - 3] == 'S') {
                        foundXmas++
                    }
                }
                if (row + 3 < crosswordPuzzle.size && col - 3 >= 0) {
                    if (crosswordPuzzle[row + 1]!![col - 1] == 'M' && crosswordPuzzle[row + 2]!![col - 2] == 'A' && crosswordPuzzle[row + 3]!![col - 3] == 'S') {
                        foundXmas++
                    }
                }
            }
        }
    }

    println("Part 1: $foundXmas")

    // Part 2
    foundXmas = 0
    crosswordPuzzle.forEach { (row, line) ->
        line.forEachIndexed { col, char ->
            val rotation = listOf('M','M','S','S','M','S','S','M','S','S','M','M','S','M','M','S')
            val sequences = rotation.windowed(4, 4, false)
            if (char == 'A') {
                if (row - 1 >= 0 && col - 1 >= 0 && row + 1 < crosswordPuzzle.size && col + 1 < line.size) {
                    for (sequence in sequences) {
                        if (crosswordPuzzle[row - 1]!![col - 1] == sequence[0]
                            && crosswordPuzzle[row - 1]!![col + 1] == sequence[1]
                            && crosswordPuzzle[row + 1]!![col + 1] == sequence[2]
                            && crosswordPuzzle[row + 1]!![col - 1] == sequence[3]) {
                            foundXmas++
                        }
                    }
                }
            }
        }
    }

    println("Part 2: $foundXmas")
}