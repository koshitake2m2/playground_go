package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	fakeGCSEndpoint := "http://localhost:4443/storage/v1/"

	client, err := storage.NewClient(ctx,
		option.WithEndpoint(fakeGCSEndpoint),
	)
	if err != nil {
		log.Fatal(err)
	}

	bucketName := "my-bucket"
	bkt := client.Bucket(bucketName)

	listObjects(ctx, bkt)

	objectName := "index.html"
	getObjectAttrs(ctx, bkt, objectName)

	if err := createObject(ctx, bkt); err != nil {
		log.Fatal(err)
	}

	listObjects(ctx, bkt)

}

func listObjects(ctx context.Context, bkt *storage.BucketHandle) error {
	query := &storage.Query{Prefix: ""}

	var objs []string
	it := bkt.Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		objs = append(objs, attrs.Name)
	}

	fmt.Println("Objects:", objs)
	return nil
}

func getObjectAttrs(ctx context.Context, bkt *storage.BucketHandle, objectName string) error {
	obj := bkt.Object(objectName)
	r, err := obj.NewReader(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Attrs:", r.Attrs)
	return nil
}

func createObject(ctx context.Context, bkt *storage.BucketHandle) error {
	uuid, _ := uuid.NewV7()
	obj := bkt.Object(uuid.String())
	w := obj.NewWriter(ctx)
	if _, err := io.WriteString(w, "<h1>hello world</h1>"); err != nil {
		return err
	}
	// if _, err := w.Write([]byte("<h1>hello world</h1>")); err != nil {
	// 	return err
	// }
	if err := w.Close(); err != nil {
		return err
	}
	return nil
}
