
public class Race {
    private final int time;
    private final int record;
    private int nWaysToBeat;

    public Race(int time, int record) {
        this.time = time;
        this.record = record;
        this.nWaysToBeat = 0;
    }

    public int calculateNWaysToBeatRecord() {
        for (int i = 1; i <= time; i++) {
            var timeAfterReleasing = time - i;
            var distanceInTry = i * timeAfterReleasing;
            if (distanceInTry > record) {
                this.nWaysToBeat++;
            }
        }
        return this.nWaysToBeat;
    }

}
