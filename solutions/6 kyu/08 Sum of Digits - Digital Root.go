package kata

func DigitalRoot(n int) (total int) {
  for n >= 10 { 
    total += n % 10
    n = n / 10
  } 
  total += n
      
  if total >= 10 {
    total = DigitalRoot(total)
  }
  return
}