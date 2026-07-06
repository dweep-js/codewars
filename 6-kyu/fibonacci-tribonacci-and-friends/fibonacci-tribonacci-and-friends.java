public class Kata {
​
    public static int[] xbonacci(int[] signature, int n) {
        if (n == 0) return new int[0];
​
        int[] result = new int[n];
​
        for (int i = 0; i < Math.min(signature.length, n); i++) {
            result[i] = signature[i];
        }
​
        for (int i = signature.length; i < n; i++) {
            for (int j = 1; j <= signature.length; j++) {
                result[i] += result[i - j];
            }
        }
​
        return result;
    }
}