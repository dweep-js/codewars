public class Square {    
    public static boolean isSquare(int n) {
      if(n<0)
        return false;
      if(Math.sqrt(n)%1==0)
        return true; 
      else
        return false;
    }
}