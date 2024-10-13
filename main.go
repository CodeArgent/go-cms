package main

func main() {
	server := NewAPIService(":3000")
	server.Run()
}
