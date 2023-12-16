class Hand(
    letters: CharArray,
    val bid: Int,
    private val enableJokers : Boolean = false
) :Comparable<Hand> {
    private lateinit var type: HandType
    private var cards: List<Card>

    init {
        val cards = mutableListOf<Card>()
        letters.forEach {
            when (it) {
                'A' -> cards.add(Card.A)
                'K' -> cards.add(Card.K)
                'Q' -> cards.add(Card.Q)
                'J' ->  {
                    if (enableJokers) cards.add(Card.JOKER) else cards.add(Card.J)
                }
                'T' -> cards.add(Card.T)
                '9' -> cards.add(Card.NINE)
                '8' -> cards.add(Card.EIGHTH)
                '7' -> cards.add(Card.SEVEN)
                '6' -> cards.add(Card.SIX)
                '5' -> cards.add(Card.FIVE)
                '4' -> cards.add(Card.FOUR)
                '3' -> cards.add(Card.THREE)
                '2' -> cards.add(Card.TWO)
            }
        }
        this.cards = cards.toList()
    }

    fun calculateType() {
        val cardsMap = mutableMapOf<Card, Int>()
        for (card in cards) {
            when (val count = cardsMap[card]) {
                null -> cardsMap[card] = 1
                else -> cardsMap[card] = count + 1
            }
        }
        when(cardsMap.keys.size) {
            1 -> type = HandType.FiveOfAKind
            2 -> {
                type = if (cardsMap.values.contains(4)) {
                    if (cardsMap.containsKey(Card.JOKER)) {
                        HandType.FiveOfAKind
                    } else {
                        HandType.FourOfAKind
                    }
                } else {
                    if (cardsMap.containsKey(Card.JOKER)) {
                        HandType.FiveOfAKind
                    } else {
                        HandType.FullHouse
                    }
                }
            }
            3 -> {
                if (cardsMap.values.contains(3)) {
                    if (cardsMap.containsKey(Card.JOKER)) {
                        when(cardsMap.getOrDefault(Card.JOKER, 0)) {
                            1 -> type = HandType.FourOfAKind
                            2 -> type = HandType.FiveOfAKind
                            3 -> type = HandType.FourOfAKind
                        }
                    } else {
                        type = HandType.ThreeOfAKind
                    }
                } else {
                    if (cardsMap.containsKey(Card.JOKER)) {
                        when(cardsMap.getOrDefault(Card.JOKER, 0)) {
                            1 -> type = HandType.FullHouse
                            2 -> type = HandType.FourOfAKind
                        }
                    } else {
                        type = HandType.TwoPair
                    }
                }
            }
            4 -> {
                type = if (cardsMap.containsKey(Card.JOKER)) {
                    HandType.ThreeOfAKind
                } else {
                    HandType.OnePair
                }
            }
            5 -> {
                type = if (cardsMap.containsKey(Card.JOKER)) {
                    HandType.OnePair
                } else {
                    HandType.HighCard
                }
            }
        }
    }

    override fun compareTo(other: Hand): Int {
        if (this.type > other.type) {
            return 1
        } else if (this.type < other.type) {
            return -1
        }
        for (i in 0..4) {
            if (this.cards[i] > other.cards[i]) {
                return 1
            }
            if (this.cards[i] < other.cards[i]) {
                return -1
            }
        }
        return 0
    }

}