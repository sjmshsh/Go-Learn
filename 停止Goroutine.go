func main() {
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:

			}
		}
	}()
	quit <- true
}
