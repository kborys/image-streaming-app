package main

import (
  "fmt"
  "image"
  "log"
  "net/http"
  "time"

  "image-streaming-app/internal/stream"
)

func handleCapture(w http.ResponseWriter, r *http.Request){
  img := image.NewRGBA(image.Rect(0, 0, 640, 480))


  filename := fmt.Sprintf("capture_%d.jpg", time.Now().Unix())
  err := stream.SaveImageToFile(img, filename)
  if err != nil {
    http.Error(w, "Failed to save image", http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, "Image save as %s", filename)
}

func main() {
  fmt.Println("Starting streaming server on :8080")

  http.HandleFunc("/stream", handleStream)
  http.HandleFunc("/capture", handleCapture)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleStream(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Streaming endpoint")
}
