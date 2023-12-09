
public class Race {
    private final long time;
    private final long record;
    private int nWaysToBeat;

    public Race(long time, long record) {
        this.time = time;
        this.record = record;
        this.nWaysToBeat = 0;
    }

    public int calculateNWaysToBeatRecord() {
        for (long i = 1; i <= time; i++) {
            var timeAfterReleasing = time - i;
            var distanceInTry = i * timeAfterReleasing;
            if (distanceInTry > record) {
                this.nWaysToBeat++;
            }
        }
        return this.nWaysToBeat;
    }

}
