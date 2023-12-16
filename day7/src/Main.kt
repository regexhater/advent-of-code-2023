import java.io.File

fun main() {
    println("day 7!")
    val lines = File("src/input.txt").readLines()
    // Part 1
    val winnings1 = calculateWinnings(lines, false)
    println("Part 1. Total winnings: $winnings1")
    // Part 2
    val winnings2 = calculateWinnings(lines, true)
    println("Part 2. Total winnings: $winnings2")
}

fun calculateWinnings(lines: List<String>, enableJokers: Boolean): Int {
    val hands = mutableListOf<Hand>()
    lines.forEach{
        val split = it.split(" ")
        hands.add(
            Hand(
                split[0].toCharArray(),
                split[1].toInt(),
                enableJokers
            )
        )

    }
    hands.forEach{
        it.calculateType()
    }
    hands.sortDescending()
    var winnings = 0
    for ((i, hand) in hands.withIndex()) {
        winnings += (i + 1) * hand.bid
    }
    return winnings
}