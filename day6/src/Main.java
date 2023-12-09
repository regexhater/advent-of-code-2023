import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;

public class Main {
    public static void main(String[] args) {
        System.out.println("day 6!");
        var input = readInput();
        var nWaysToBeatMul = 1;
        for (var race : input) {
            nWaysToBeatMul = nWaysToBeatMul * race.calculateNWaysToBeatRecord();
        }
        System.out.println("The multiplication o number o ways to beat is: " + nWaysToBeatMul);
    }

    public static List<Race> readInput() {
        var path = Paths.get("src/input.txt");
        var times = new ArrayList<Integer>();
        var distances = new ArrayList<Integer>();
        var races = new ArrayList<Race>();
        List<String> lines;
        try {
            lines = Files.readAllLines(path, StandardCharsets.UTF_8);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
        if (lines != null) {
            var time = lines.getFirst();
            time = time.split(": ")[1].trim();
            for (var s : time.split(" ")) {
                if (s.isBlank()) {
                    continue;
                }
                times.add(Integer.parseInt(s));
            }
            var distance = lines.getLast();
            distance = distance.split(": ")[1].trim();
            for (var s: distance.split(" ")) {
                if (s.isBlank()) {
                    continue;
                }
                distances.add(Integer.parseInt(s));
            }
        }
        if (times.size() != distances.size()) {
            throw new RuntimeException();
        }
        for (int i = 0; i < times.size(); i++) {
            races.add(new Race(times.get(i), distances.get(i)));
        }
        return races;
    }
}