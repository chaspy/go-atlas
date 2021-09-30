package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mongodb-forks/digest"
	"go.mongodb.org/atlas/mongodbatlas"
)

func main() {
	publicKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
	privateKey := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")
	t := digest.NewTransport(publicKey, privateKey)
	tc, err := t.Client()
	if err != nil {
		log.Fatalf(err.Error())
	}

	client := mongodbatlas.NewClient(tc)
	projects, _, err := client.Projects.GetAllProjects(context.Background(), nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%#v", projects)
//	fmt.Printf("%#v", res)
}
