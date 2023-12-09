import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class AlmanacMap {
    private final List<Range> ranges;

    public AlmanacMap(List<Range> ranges) {
        this.ranges = Objects.requireNonNullElseGet(ranges, ArrayList::new);
    }

    public void addRange(Range range) {
        ranges.add(range);
    }

    public long translate(long source) {
       for (var range : ranges) {
           if (source >= range.sourceStart() && source <= range.sourceStart() + range.shift()) {
               var difference = source - range.sourceStart();
               return range.destinationStart() + difference;
           }
       }
       return source;
    }
}
