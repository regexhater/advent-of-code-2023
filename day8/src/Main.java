import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.HashMap;

public class Main {
    public static void main(String[] args) {
        System.out.println("day 8!");
        var input = readInput();
        var node = "AAA";
        var stepsNeeded = getSteps(input, node, "ZZZ");
        System.out.printf("There needs to be %d steps to the end\n", stepsNeeded);

        // Part 2
        // We calculate min steps for each start separately then find The Least Common Multiple
        // also known as the Least Common Denominator. If we take a look at scores when trying to calculate
        // all at the same time we can notice that there is a cycle for each of them meaning that if one of them
        // reaches the end in s1 steps it will reach it again in s2 steps where s2 = 2s1 and so on.
        // This leads us to just find a place when those cycles sync up thus the LCM method.
        var minStepsForMultipleStarts = input
                .mapping()
                .keySet()
                .stream()
                .filter(x -> x.endsWith("A"))
                .map(x -> getSteps(input, x, "Z"))
                .reduce(Main::lcm)
                .orElse(-1L);

        System.out.printf("There needs to be %d simulations steps to the end\n", minStepsForMultipleStarts);
    }


    public static Input readInput() {
        BufferedReader reader;
        var directions = "";
        var mapping = new HashMap<String, String[]>();
        try {
            reader = new BufferedReader(new FileReader("src/input.txt"));
            // First line contains directions;
            directions = reader.readLine();
            var line = reader.readLine();
            while (line != null) {
                if (line.isBlank()) {
                    line = reader.readLine();
                    continue;
                }
                var split = line.split(" = ");
                var paths = split[1]
                        .replace("(", "")
                        .replace(")", "")
                        .split(", ");
                mapping.put(split[0], paths);
                line = reader.readLine();
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

        return new Input(directions, mapping);
    }

    public static long getSteps(Input input, String current, String endsWith) {
        long steps = 0;
        int i = 0;
        while (!current.endsWith(endsWith)) {
            if (input.directions().charAt(i) == 'L') {
                current = input.mapping().get(current)[0];
            } else {
                current = input.mapping().get(current)[1];
            }
            i++;
            steps++;
            if (i >= input.directions().length()) {
                i = 0;
            }
        }

        return steps;
    }

    public static long lcm(long a, long b) {
        var ma = a;
        var mb = b;
        long remainder;

        while (mb != 0L) {
            remainder = ma % mb;
            ma = mb;
            mb = remainder;
        }

        return a * b / ma;
    }
}

