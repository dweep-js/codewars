package kata
​
import "math"
​
func WallPaper(l, w, h float64) string {
  numbers := []string{
    "zero", "one", "two", "three", "four",
    "five", "six", "seven", "eight", "nine",
    "ten", "eleven", "twelve", "thirteen", "fourteen",
    "fifteen", "sixteen", "seventeen", "eighteen",
    "nineteen", "twenty",
  }
​
  if l == 0 || w == 0 || h == 0 {
    return numbers[0]
  }
​
  wallArea := 2 * h * (l + w)
  requiredArea := wallArea * 1.15
  rollArea := 10.0 * 0.52
​
  rolls := int(math.Ceil(requiredArea / rollArea))
​
  return numbers[rolls]
}