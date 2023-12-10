import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.HashMap;

public class Main {
    public static void main(String[] args) {
        System.out.println("day 8!");
        var input = readInput();
        var node = "AAA";
        var i = 0;
        var counter = 0;
        while (!node.equals("ZZZ")) {
            if (input.directions().charAt(i) == 'L') {
                node = input.mapping().get(node)[0];
            } else {
                node = input.mapping().get(node)[1];
            }
            i++;
            counter++;
            if (i >= input.directions().length()) {
                i = 0;
            }
        }
        System.out.printf("There needs to be %d steps to the end\n", counter);
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
}