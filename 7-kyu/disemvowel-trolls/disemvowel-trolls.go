package kata
​
import "strings"
​
func Disemvowel(comment string) string {
  comment = strings.ReplaceAll(comment, "a", "")
  comment = strings.ReplaceAll(comment, "e", "")
  comment = strings.ReplaceAll(comment, "i", "")
  comment = strings.ReplaceAll(comment, "o", "")
  comment = strings.ReplaceAll(comment, "u", "")
  comment = strings.ReplaceAll(comment, "A", "")
  comment = strings.ReplaceAll(comment, "E", "")
  comment = strings.ReplaceAll(comment, "I", "")
  comment = strings.ReplaceAll(comment, "O", "")
  comment = strings.ReplaceAll(comment, "U", "")
​
  return comment
}