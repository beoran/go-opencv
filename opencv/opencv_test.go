import "testing"
import "opencv"



func TestLoad(t *testing.T) {
  filename := "test_input.png"
  image    := opencv.LoadImage(filename, 0)  
}



