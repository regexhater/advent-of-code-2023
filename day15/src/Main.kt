import java.io.File

fun main() {
    println("day 15!")
    val input = File("src/input.txt").readText().trim()
    var sum = 0
    for (seq in input.split(",")) {
        sum += hash(seq)
    }
    println("Part 1. The sum of results if $sum")
}

fun hash(input: String): Int {
    var currentValue = 0
    for (c in input.toCharArray()) {
        currentValue += c.code
        currentValue *= 17
        currentValue %= 256
    }
    return currentValue
}