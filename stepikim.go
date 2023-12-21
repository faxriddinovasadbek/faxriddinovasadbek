func main() {
    
    sum := 0
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
        i, _ := strconv.Atoi(scanner.Text())
        sum += i
    }
    
    io.WriteString(os.Stdout, strconv.Itoa(sum))
}
