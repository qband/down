package main

import (
	"os"

	// cli
	"github.com/codegangsta/cli"

	// gRPC
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	pb "helloworld"
)

func main() {
	app := cli.NewApp()
	app.Name = "down-cli"
	app.Usage = "a cli to add download task"

	// Flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "127.0.0.1",
			Usage: "downloader ssh [host]:[port]",
		},
	}
	app.Action = func(c *cli.Context) {
		cmd := "default"
		if len(c.Args()) > 0 {
			cmd = c.Args()[0]
		}

		// check flags
		if c.String("host") == "test" {
			println("--host", "test")
		}

		// check main option
		if cmd == "default" {
			println("cmd:", cmd)
			testRPC()
		} else {
			println("command opts not found")
		}
	}

	// Subcommands
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) {
				println("added task: ", c.Args().First())
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) {
				println("completed task: ", c.Args().First())
			},
		},
		{
			Name:    "template",
			Aliases: []string{"r"},
			Usage:   "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) {
						println("new task template: ", c.Args().First())
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) {
						println("removed task template: ", c.Args().First())
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func testRPC() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
