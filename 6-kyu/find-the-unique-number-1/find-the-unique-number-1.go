package kata
​
func FindUniq(arr []float32) float32 {
  if arr[0] != arr[1] {
    if arr[0] == arr[2] {
      return arr[1]
    }
    return arr[0]
  }
  common := arr[0]
  for _, v := range arr {
    if v != common {
      return v
    }
  }
  return common
}