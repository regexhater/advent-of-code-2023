import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;

public class Main {
    public static void main(String[] args) {
        System.out.println("day5!");
        var input = readInput();
        var smallestLocation = Long.MAX_VALUE;
        for (var seed : input.seeds()) {
            var entry = seed;
            for (var mapping : input.mappings()) {
                entry = mapping.translate(entry);
            }
            if (entry < smallestLocation) {
                smallestLocation = entry;
            }
        }
        System.out.println("The smallest location number is: " + smallestLocation);
    }

    public static Input readInput() {
        var mappings = new ArrayList<AlmanacMap>();
        var seeds = new ArrayList<Long>();
        BufferedReader reader;

        try {
            reader = new BufferedReader(new FileReader("src/input.txt"));
            var line = reader.readLine();
            var map = new AlmanacMap(new ArrayList<>());
            while (line != null) {
                if (line.isBlank()) {
                    mappings.add(map);
                    map = new AlmanacMap(new ArrayList<>());
                    line = reader.readLine();
                    continue;
                }
                if (!line.contains("map")){
                    if (line.contains("seeds")) {
                        var seedsLineSplit = line.split(": ");
                        var seedsSplit = seedsLineSplit[1].split(" ");
                        for (var seed : seedsSplit) {
                            seeds.add(Long.parseLong(seed));
                        }
                    } else {
                        var mapSplit = line.split(" ");
                        map.addRange(new Range(
                                Long.parseLong(mapSplit[0]),
                                Long.parseLong(mapSplit[1]),
                                Long.parseLong(mapSplit[2]))
                        );
                    }
                }
                line = reader.readLine();
            }
            mappings.add(map);

        } catch (IOException e) {
            throw new RuntimeException(e);
        }
        // remove first element which is empty
        mappings.removeFirst();
        return new Input(seeds, mappings);
    }
}

