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
  
  ctx := context.Background()
  clusterName := "test-source-cluster"
  groupID := "6155d1572d865910d4137d3a"

	cluster, _, err := client.Clusters.Get(ctx, groupID, clusterName,)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%#v", cluster.Name)
}
