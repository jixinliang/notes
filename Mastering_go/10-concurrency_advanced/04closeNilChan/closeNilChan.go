package main

func main() {
	var c chan string
	close(c)
}

/*
trying to read or write from a nil channel will block. This property of channels can be
very useful when you want to disable a branch of a select statement by assigning
the nil value to a channel variable
*/