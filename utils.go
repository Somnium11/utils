package utils

func InSlice(a []string, x string) bool {
  for _, n := range a {
     if x == n {
        return true
     }
  }
  return false
}

func SliceInt(a []int, x int) bool {
   for _, n := range a {
      if x == n {
         return true
      }
   }
   return false
 }