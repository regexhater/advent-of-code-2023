import java.nio.charset.StandardCharsets
import java.nio.file.Files
import java.nio.file.Paths

fun main() {
    println("day 9!")
    val path = Paths.get("src/input.txt")
    val lines = Files.readAllLines(path, StandardCharsets.UTF_8)
    part1(lines)
    part2(lines)
}

fun part1(input: MutableList<String>) {
    var sumOfExtrapolatedValues = 0
    for (line in input) {
        val matrix = generatePredictionMatrix(line)
        matrix.last().add(0)
        for (i in matrix.size -2 downTo 0) {
            val nextVal = matrix[i].last() + matrix[i + 1].last()
            matrix[i].add(nextVal)
        }
        sumOfExtrapolatedValues += matrix.first().last()
    }
    println("Part 1. The sum of Extrapolated values is equal: $sumOfExtrapolatedValues")
}

fun part2(input: MutableList<String>) {
    var sumOfExtrapolatedValues = 0
    for (line in input) {
        val matrix = generatePredictionMatrix(line)
        var newRow = mutableListOf(0)
        newRow.addAll(matrix.last())
        matrix[matrix.size - 1] = newRow
        for (i in matrix.size -2 downTo 0) {
            val prevVal = matrix[i].first() - matrix[i + 1].first()
            newRow = mutableListOf(prevVal)
            newRow.addAll(matrix[i])
            matrix[i] = newRow
        }
        sumOfExtrapolatedValues += matrix.first().first()
    }
    println("Part 2. The sum of Prev Extrapolated values is equal: $sumOfExtrapolatedValues")
}

private fun generatePredictionMatrix(line: String): MutableList<MutableList<Int>> {
    val matrix = mutableListOf<MutableList<Int>>()
    val ml = mutableListOf<Int>()
    for (s in line.split(" ")) {
        ml.add(Integer.parseInt(s))
    }
    matrix.add(ml)
    var diff = getDiffList(ml)
    matrix.add(diff)
    while (!diff.all { x -> x == 0 }) {
        diff = getDiffList(diff)
        matrix.add(diff)
    }
    return matrix
}

fun getDiffList(list: MutableList<Int>) : MutableList<Int> {
    val result = mutableListOf<Int>()
    for (i in 0..<list.size -1) {
        val diff = list[i + 1] - list[i]
        result.add(diff)
    }
    return result
}