package stream

import (
	"image"
  "image/jpeg"
	"sync"
  "os"
)

func SaveImageToFile(img image.Image, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    return jpeg.Encode(file, img, &jpeg.Options{Quality: 90})

}

type ImageStream struct { 
  currentFrame image.Image
  mutex sync.RWMutex
}

func NewImageStream() *ImageStream {
  return &ImageStream{}
}

func (s *ImageStream) UpdateFrame(img image.Image) {
  s.mutex.Lock()
  defer s.mutex.Unlock()
  s.currentFrame = img
}

func (s *ImageStream) GetCurrentFrame() image.Image {
  s.mutex.RLock()
  defer s.mutex.RUnlock()
  return s.currentFrame
}
