import java.io.File

fun main() {
    println("day 15!")
    val input = File("src/input.txt").readText().trim()
    part1(input)
    part2(input)
}

fun part1(input: String) {
    var sum = 0
    for (seq in input.split(",")) {
        sum += hash(seq)
    }
    println("Part 1. The sum of results is $sum")
}

fun part2(input: String) {
    var focusingPower = 0
    val map = hashMapOf<Int, MutableList<Lens>>()
    for (seq in input.split(",")) {
        if (seq.contains("-")) {
            val label = seq.split("-")[0]
            val hash = hash(label)
            map[hash]?.remove(Lens(label))
        } else {
            val split = seq.split("=")
            val label = split[0]
            val focalStrength = split[1].toInt()
            val hash = hash(label)
            val lens = Lens(label, focalStrength)
            if (map.containsKey(hash)) {
                if (map[hash]!!.contains(lens)) {
                    map[hash]!![map[hash]!!.indexOf(lens)] = lens
                } else {
                    map[hash]!!.add(lens)
                }
            } else {
                map[hash] = mutableListOf(lens)
            }
        }
    }
    for ((key, value) in map.entries) {
        for ((index, lens) in value.withIndex()) {
            focusingPower += (key + 1) * (index + 1) * lens.focalLength
        }
    }
    println("Part 2. The sum of focusing powers is $focusingPower")
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