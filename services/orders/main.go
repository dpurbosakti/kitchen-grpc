package main

func main() {
	httpServer := NewHTTPServer(":8080")
	go httpServer.Run()

	gRPCServer := NewGRPCServer(":9000")
	gRPCServer.Run()

}
