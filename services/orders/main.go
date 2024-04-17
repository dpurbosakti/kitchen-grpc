package main

func main() {
	httpServer := newHTTPOrderServer(":8080")
	go httpServer.Run()

	gRPCServer := newGRPCOrderServer(":9000")
	gRPCServer.Run()

}
