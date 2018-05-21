package main

import (
	"cloud.google.com/go/storage"
	"cloud.google.com/go/vision/apiv1"
	"fmt"
	"golang.org/x/net/context"
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

func TestCanPerformTextDetection(t *testing.T) {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(context.Background())

	file := "gs://postit/IMG_7237.JPG"
	image := vision.NewImageFromURI(file)
	annotations, err := client.DetectDocumentText(ctx, image, nil)

	if err != nil {
		t.Errorf("%v", err)
	}

	if annotations == nil {
		t.Errorf("No text found")
	} else {
		for _, page := range annotations.Pages {
			for _, block := range page.Blocks {
				for _, p := range block.Paragraphs {
					fmt.Printf("\n\n")
					for _, w := range p.Words {
						for _, s := range w.Symbols {
							fmt.Printf("%v", s.Text)
						}
                                        fmt.Printf(" ")
                                        }
				}
			}
		}
	}
}
