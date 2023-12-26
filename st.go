func removeDuplicates(inputStream chan string, outputStream chan string) {
	defer close(outputStream)
	var t string
	for {
	  val, ok := <-inputStream
	  if !ok {
		break
	  }
	  if val != t {
		outputStream <- val
		t = val
	  }
	}
}
