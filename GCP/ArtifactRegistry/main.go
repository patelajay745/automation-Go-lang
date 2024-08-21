package main

import (
	"context"
	"fmt"
	"log"

	artifactregistry "cloud.google.com/go/artifactregistry/apiv1"
	"cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"google.golang.org/api/option"
)

func main() {

	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx, option.WithCredentialsFile("../../keys/artifact-key.json"))
	if err != nil {
		panic(err)
	}
	defer c.Close()
	getDockerImage(c, ctx)

}

func getDockerImage(c *artifactregistry.Client, ctx context.Context) {
	req := &artifactregistrypb.ListDockerImagesRequest{Parent: "projects/personal-project-429201/locations/us-east1/repositories/my-repo", OrderBy: "UPLOAD_TIME desc"}

	it := c.ListDockerImages(ctx, req)
	for {
		resp, err := it.Next()
		if err != nil {
			if err.Error() == "no more items in iterator" {
				break
			}
			log.Fatalf("Failed to list docker images: %v", err)
		}
		// fullName := resp.GetName()
		// parts := strings.Split(fullName, "/")
		// imageName := parts[len(parts)-1]

		fmt.Println("Docker Image Name: ", resp.GetTags())
	}
}

func listAllRepository(c *artifactregistry.Client, ctx context.Context) {
	req := &artifactregistrypb.ListRepositoriesRequest{Parent: "projects/personal-project-429201/locations/us-east1"}

	resp := c.ListRepositories(ctx, req)

	for {
		repo, err := resp.Next()
		if err != nil {
			if err.Error() == "no more items in iterator" {
				break
			}
			log.Fatalf("Failed to list repositories: %v", err)
		}
		fmt.Printf("Repository: %s\n", repo.Name)
	}
}
