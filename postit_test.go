package main

import (
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
        "cloud.google.com/go/vision/apiv1"
        "fmt"
	"io"
	"os"
	"testing"
)

func _TestCanUploadToGoogleCloudStorage(t *testing.T) {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		// TODO: handle error.
		t.Errorf("%v", err)
	}
	_ = client // Use the client.

	//bkt := client.Bucket("postit")

	f, err := os.Open("image.png")
	if err != nil {
		t.Errorf("%v", err)
	}
	defer f.Close()

	wc := client.Bucket("postit").Object("test_file_upload").NewWriter(context.Background())
	if _, err = io.Copy(wc, f); err != nil {
		t.Errorf("%v", err)
	}
	if err := wc.Close(); err != nil {
		t.Errorf("%v", err)
	}
}

func _TestCanPerformTextDetection(t *testing.T) {
        ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(context.Background())

        file := "gs://postit/IMG_7237.JPG"
	image := vision.NewImageFromURI(file)
        annotations, err := client.DetectTexts(ctx, image, nil, 10)

	if err != nil {
		t.Errorf("%v", err)
	}

        if len(annotations) == 0 {
                t.Errorf("No text found") 
        } else {
                fmt.Println("Text:")
                for _, annotation := range annotations {
                        fmt.Printf("%v\n", annotation.Description)
                }
        }
}
