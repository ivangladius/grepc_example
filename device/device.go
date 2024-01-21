package main

import (
	"context"
	"log"
	pb "myrpc/bigboss/student"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

const (
	address       string  = "localhost:9999"
	defaultName   string  = "hans"
	defaultAge    int     = 99
	defaultWeight float64 = 0.001
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Dial failed: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewNetworkClient(conn)

	name := defaultName
	// age := defaultAge
	weight := defaultWeight
	if len(os.Args) >= 4 {
		name = os.Args[1]
		// age, err = strconv.Atoi(os.Args[2])
		// if err != nil {
		// 	panic(err)
		// }
		weight, err = strconv.ParseFloat(os.Args[3], 32)
		if err != nil {
			panic(err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for i := 0; i < 100; i++ {
		r, err := c.WelcomeStudent(ctx,
			&pb.Student{
				Name:   name,
				Age:    int32(i),
				Weight: float32(weight)})
		if err != nil {
			log.Fatalf("could not be welcomed: %v\n", err)
		}
		log.Printf("server send: %s\n", r.GetMessage())
	}

}
