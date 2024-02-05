package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleIndex)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server is listening on :%s...\n", port)
	http.ListenAndServe(":"+port, nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	// Загрузка содержимого index.html из Cloud Storage
	content, err := loadFileFromCloudStorage(ctx, "your-bucket-name", "index.html")
	if err != nil {
		http.Error(w, "Error loading file from Cloud Storage", http.StatusInternalServerError)
		return
	}
	// Установка заголовков и отправка содержимого
	w.Header().Set("Content-Type", "text/html")
	w.Write(content)
}

func loadFileFromCloudStorage(ctx context.Context, bucket, object string) ([]byte, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	content, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	return content, nil
}
