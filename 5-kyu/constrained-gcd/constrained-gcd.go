package kata
​
func Solve(n, k int) []int {
  m := k * (k + 1) / 2
  if n < m {
    return []int{}
  }
  b, g := 1, 1
  for g*g <= n {
    if n%g == 0 {
      if g <= n/m && g > b {
        b = g
      }
      if o := n / g; o <= n/m && o > b {
        b = o
      }
    }
    g++
  }
  t := n / b
  r := make([]int, k)
  s := 0
  for i := 0; i < k-1; i++ {
    r[i] = (i + 1) * b
    s += i + 1
  }
  r[k-1] = (t - s) * b
  return r
}