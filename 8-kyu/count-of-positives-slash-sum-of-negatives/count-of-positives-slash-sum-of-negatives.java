public class Kata {
    public static int[] countPositivesSumNegatives(int[] input) {
        if (input == null || input.length == 0) {
            return new int[0];
        }
​
        int[] n = new int[2];
​
        for (int i = 0; i < input.length; i++) {
            if (input[i] > 0) {
                n[0]++;          // Count positives
            } else if (input[i] < 0) {
                n[1] += input[i]; // Sum negatives
            }
        }
​
        return n;
    }
}