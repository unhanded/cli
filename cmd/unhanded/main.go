package main

func welcomeMessage() string {
	return "----------------\nUNHANDED CLI\n----------------"
}

func main() {
	println(welcomeMessage())
	panic("You're too early (not implemented)")
}
