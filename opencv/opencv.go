/*
Go Language wrappers around Open CV
*/
package opencv

// #include <opencv/cv.h>
// #include <opencv/cvaux.h>
// #include <opencv/highgui.h>
import "C"
import "unsafe" 
import "fmt"

type Image struct { 
  cimage * C.IplImage
}


// free is a method on C char * strings to method to free the associated memory 
func (self * C.char) free() {
  C.free(unsafe.Pointer(self))
}

// cstring converts a string th a C string. This alloactes memory, so don't 
// forget a defer s.free 
func cstr(self string) (* C.char) {
  return C.CString(self)
}

// SetErrStatus sets the error status.
// The function sets the error status to the specified value. 
// Mostly, the function is used to reset the error status (set to it StsOk) 
// to recover after an error. In other cases it is more natural to call Error.
func SetErrStatus(status int) {
  C.cvSetErrStatus(C.int(status))
}


// constants for SetErrMode
const (
  /* Print error and exit program */
  ErrModeLeaf     =  0   
  /* Print error and continue */
  ErrModeParent   =  1   
  /* Don't print and continue */
  ErrModeSilent   =  2   
)

// SetErrMode sets the specified error mode. For descriptions of different 
// error modes, see the beginning of the error section.
func SetErrMode(mode int) {
  C.cvSetErrMode(C.int(mode))
}


// GetErrStatus returns the current error status.
// The function returns the current error status - the value set with the last 
// SetErrStatus call. Note that in Leaf mode, the program terminates immediately 
// after an error occurs, so to always gain control after the function call, 
// one should call SetErrMode and set the Parent or Silent error mode.
func GetErrStatus() int {
  return int(C.cvGetErrStatus());
}

// GetErrMode returns the current error mode.
// The function returns the current error mode - the value set with the last 
// SetErrMode call.
func GetErrMode() int {
  return int(C.cvGetErrMode());
}

// Error raises an error.
// Parameters: 
//        * status – The error status
//        * func_name – Name of the function where the error occured
//        * err_msg – Additional information/diagnostics about the error
//        * filename – Name of the file where the error occured
//        * line – Line number, where the error occured
// The function sets the error status to the specified value (via SetErrStatus) 
// and, if the error mode is not Silent, calls the error handler.

func Error(status int, func_name, err_msg, filename string, line int) {
  cfunc   := cstr(func_name)  ; defer cfunc.free()
  cerr    := cstr(err_msg)    ; defer cerr.free()
  cfile   := cstr(filename)   ; defer cfile.free()
  cstatus := C.int(status)
  cline   := C.int(line)
  C.cvError(cstatus, cfunc, cerr, cfile, cline)
}

//ErrorStr returns textual description of an error status code.
//Parameter:  status – The error status

//The function returns the textual description for the specified error status 
// code. In the case of unknown status, the function returns a NULL pointer.
func ErrorStr(status int) string {
  cstr    := C.cvErrorStr(C.int(status))
  return C.GoString(cstr)
}

// RedirectError and other error callbacks not supported 



// Initializes opencv, particularily it's error handling
func Init() {
  SetErrMode(ErrModeParent)
}


/*
func (Image self) destroy() {
  C.IplImageFree(self)
}
*/

func WrapImage(cimage * C.IplImage) * Image {
  if cimage == nil {
    return nil
  }
  return &Image{cimage}
} 

func LoadImage(filename string, iscolor int) * Image {
  cfile   := C.CString(filename)
  defer   cfile.free()
  ccolor  := C.int(iscolor)  
  cimage  := C.cvLoadImage(cfile, ccolor)
  
  if cimage == nil {
    return nil
  }
  return WrapImage(cimage)
}

func (self *C.IplImage) releaseimage() {
  C.cvReleaseImage(&self)
}

//Constantd declarations for SaveEX
const (
  IMWRITE_JPEG_QUALITY      = 1
  IMWRITE_PNG_COMPRESSION   = 16
  CV_IMWRITE_PXM_BINARY     = 32
)

//SaveEx saves the image to the named file name, with extra quality parameters.
//Returns true on success or false on failiure.
func (self * Image) SaveEx(filename string, quality int) (* Image) {
  cfile   := C.CString(filename)
  defer   cfile.free()
  if self.cimage == nil { return nil }
  cimage  := unsafe.Pointer(self.cimage)
  params  := make([]int, 1)
  params[0]= quality
  cparams := unsafe.Pointer(&params[0])
  res     := C.cvSaveImage(cfile, cimage, (*C.int)(cparams))
  if int(res) > 0  {  return self }
  return nil; 
}  

//Save saves the image to the named file name. 
//Returns true on success or false on failiure.
func (self * Image) Save(filename string) (* Image) {
  return self.SaveEx(filename, 0)
}
  
  

// Release releases the memory associated with the block
func (self * Image) Release() {
  if self.cimage != nil {
    self.cimage.releaseimage()
  }  
  self.cimage = nil
}

// Destructor for GC. Doesn't work yet.
func (self * Image) destroy() {
  self.Release()
  fmt.Println("Released!")
}






