import java.io.File

fun main() {
    println("day 7!")
    val lines = File("src/input.txt").readLines()
    val hands = mutableListOf<Hand>()
    lines.forEach{
        val split = it.split(" ")
        hands.add(
            Hand(
                split[0].toCharArray(),
                split[1].toInt()
            )
        )

    }
    hands.forEach{
        it.calculateType()
    }
    hands.sortDescending()
    var counter = 0
    for ((i, hand) in hands.withIndex()) {
        counter += (i + 1) * hand.bid
    }
    println("Part 1. Total winnings: $counter")
}