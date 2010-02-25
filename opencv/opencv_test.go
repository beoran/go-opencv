import "opencv"




func TestLoad() {
  filename := "/group/fgd-ssg/bjmey/rudy/F1_1_untreated/F1--W00061--P00001--Z00000--T00000--FITC.png"
  image    := opencv.LoadImage(filename, 0)
}



