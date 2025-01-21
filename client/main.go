package main
import(
	proto "www.github.com/amanraj8848/gRPC-golang/proto"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PORT = ":2500"
)



func main() {
	conn, err := grpc.Dial("localhost" + PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(" failed to connect: %v", err)
	}	
	defer conn.Close() 

	client := proto.NewGreetServiceClient(conn) // registring the client with the server
	// req := &proto.HelloRequest{
	// 	name: "Aman",
	// }	
	callSayHello(client)
}