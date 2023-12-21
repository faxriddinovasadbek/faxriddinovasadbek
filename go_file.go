func main(){
  
  var n , fn int

  print("n = ")
  fmt.Scan(&n)

  number , err := os.Create("nember.txt")

  defer number.Close()
  if err != nil{
    panic(err)
  }

  for i := 0; i < n; i++ {
    print(i+1, " chi soni kiriting: ")
    fmt.Scan(&fn)

    a := strconv.Itoa(fn)
    a += "\n"
    
    number.WriteString(a)
  }
  
  nums , err := os.ReadFile("nember.txt")

  if err != nil{
    panic(err)
  }

  lst := strings.Split(string(nums) , "\n")

  
  summa := 0

  for _, l := range lst {
    ls , _ := strconv.Atoi(l)
    summa += ls
  }  

  sum := strconv.Itoa(summa)

  number.WriteString("\nnatija: ")
  number.WriteString(sum)

}
