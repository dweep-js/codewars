import java.util.StringTokenizer;
​
public class JadenCase {
​
  public String toJadenCase(String phrase) {
    if (phrase == null || phrase.isEmpty()) return null;
​
    String s = "";
    StringTokenizer words = new StringTokenizer(phrase);    
    while (words.hasMoreTokens()) {
      String str = words.nextToken();
      str = Character.toUpperCase(str.charAt(0)) + str.substring(1);
      s = s + str + (words.hasMoreTokens() ? " " : ""); 
    }
    
    return s;
  }
​
}
​