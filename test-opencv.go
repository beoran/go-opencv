package main

import "opencv"
import "fmt"
import "os"


func error(error string) {
  fmt.Fprintln(os.Stderr, error);
}

func assert(cond bool, err string)  {
  if cond { 
    return 
  }
  error(err) 
}

func TestLoadRelease() {
  filename := "test_input.png"
  image    := opencv.LoadImage(filename, 0)
  image.Release()
}

func TestSave() {
  filename := "test_input.png"
  image    := opencv.LoadImage(filename, 0)
  defer       image.Release()
  result   := image.Save("test_out.jpg")
  assert(result != nil, "Result of save should not be nil.")
}

func frobnicate (fake * opencv.Image) int {
  return 1
}

func main() {
  opencv.Init()
  TestLoadRelease()
  TestSave()  
}
