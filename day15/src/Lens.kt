class Lens(val label: String, val focalLength: Int = 0) {
    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as Lens

        return label == other.label
    }

    override fun hashCode(): Int {
        return label.hashCode()
    }
}