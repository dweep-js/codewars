public class Maskify {
    public static String maskify(String str) {
        if (str.length() <= 4) {
            return str;
        }
​
        StringBuilder result = new StringBuilder();
​
        for (int i = 0; i < str.length() - 4; i++) {
            result.append('#');
        }
​
        result.append(str.substring(str.length() - 4));
​
        return result.toString();
    }
}