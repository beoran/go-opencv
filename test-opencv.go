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
  filename := "/group/fgd-ssg/bjmey/rudy/F1_1_untreated/F1--W00061--P00001--Z00000--T00000--FITC.png"
  image    := opencv.LoadImage(filename, 0)
  image.Release()
}

func TestSave() {
  filename := "/group/fgd-ssg/bjmey/rudy/F1_1_untreated/F1--W00061--P00001--Z00000--T00000--FITC.png"
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
