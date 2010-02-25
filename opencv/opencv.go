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

// Constants for image.Convert
const (
  CVTIMG_FLIP 		= 1
  CVTIMG_SWAP_RB 	= 2
)

// Concert converts one image to another with an optional vertical flip.
func (self * Image) Convert(destination * Image, flags int) {
  C.cvConvertImage(unsafe.Pointer(self.cimage), unsafe.Pointer(destination.cimage), C.int(flags))  
} 

type Trackbar struct {
  handle int;
  value int;

}  

// CreateTrackbar creates a trackbar and attaches it to the specified window.
// Does not support callbacks yet.
func CreateTrackbar(name string, window string, value int, max int) * Trackbar {
  cname 	:= cstr(name) 	; defer cname.free()
  cwindow	:= cstr(window)	; defer cwindow.free()
  cmax	       	:= C.int(max)
  trackbar     	:= &Trackbar{0, value}
  cvalue       	:= (* C.int)(unsafe.Pointer(&trackbar.value))
  chandle      	:= C.cvCreateTrackbar(cname, cwindow, cvalue, cmax, nil)
  trackbar.handle = int(chandle)
  return trackbar
}

// DestroyAllWindows() destroys all of the opened HighGUI windows.
func DestroyAllWindows() {
  C.cvDestroyAllWindows()
}


type Window struct {
  name string
}

func (self * Window) Destroy() {
  cname := cstr(self.name) ; defer cname.free()
  C.cvDestroyWindow(cname)
}



/*
DestroyWindow¶

void cvDestroyWindow(const char* name)¶

    Destroys a window.
    Parameter:  name – Name of the window to be destroyed.

    The function cvDestroyWindow() destroys the window with the given name.

GetTrackbarPos¶

int cvGetTrackbarPos(const char* trackbarName, const char* windowName)¶

    Returns the trackbar position.
    Parameters: 

        * trackbarName – Name of the trackbar.
        * windowName – Name of the window which is the parent of the trackbar.

    The function cvGetTrackbarPos() returns the current position of the specified trackbar.

GetWindowHandle¶

void* cvGetWindowHandle(const char* name)¶

    Gets the window’s handle by its name.
    Parameter:  name – Name of the window.

    The function cvGetWindowHandle() returns the native window handle (HWND in case of Win32 and GtkWidget in case of GTK+).

GetWindowName¶

const char* cvGetWindowName(void* windowHandle)¶

    Gets the window’s name by its handle.
    Parameter:  windowHandle – Handle of the window.

    The function cvGetWindowName() returns the name of the window given its native handle (HWND in case of Win32 and GtkWidget in case of GTK+).

InitSystem¶

int cvInitSystem(int argc, char** argv)¶

    Initializes HighGUI.
    Parameters: 

        * argc – Number of command line arguments
        * argv – Array of command line arguments

    The function cvInitSystem() initializes HighGUI. If it wasn’t called explicitly by the user before the first window was created, it is called implicitly then with argc=0, argv=NULL. Under Win32 there is no need to call it explicitly. Under X Window the arguments may be used to customize a look of HighGUI windows and controls.

MoveWindow¶

void cvMoveWindow(const char* name, int x, int y)¶

    Sets the position of the window.
    Parameters: 

        * name – Name of the window to be moved.
        * x – New x coordinate of the top-left corner
        * y – New y coordinate of the top-left corner

    The function cvMoveWindow() changes the position of the window.

NamedWindow¶

int cvNamedWindow(const char* name, int flags)¶

    Creates a window.
    Parameters: 

        * name – Name of the window in the window caption that may be used as a window identifier.
        * flags – Flags of the window. Currently the only supported flag is CV_WINDOW_AUTOSIZE. If this is set, window size is automatically adjusted to fit the displayed image (see ShowImage), and the user can not change the window size manually.

    The function cvNamedWindow() creates a window which can be used as a placeholder for images and trackbars. Created windows are referred to by their names.

    If a window with the same name already exists, the function does nothing.

ResizeWindow¶

void cvResizeWindow(const char* name, int width, int height)¶

    Sets the window size.
    Parameters: 

        * name – Name of the window to be resized.
        * width – New width
        * height – New height

    The function cvResizeWindow() changes the size of the window.

SetMouseCallback¶

void cvSetMouseCallback(const char* windowName, CvMouseCallback onMouse, void* param=NULL)¶

    Assigns callback for mouse events.

        #define CV_EVENT_MOUSEMOVE      0
        #define CV_EVENT_LBUTTONDOWN    1
        #define CV_EVENT_RBUTTONDOWN    2
        #define CV_EVENT_MBUTTONDOWN    3
        #define CV_EVENT_LBUTTONUP      4
        #define CV_EVENT_RBUTTONUP      5
        #define CV_EVENT_MBUTTONUP      6
        #define CV_EVENT_LBUTTONDBLCLK  7
        #define CV_EVENT_RBUTTONDBLCLK  8
        #define CV_EVENT_MBUTTONDBLCLK  9

        #define CV_EVENT_FLAG_LBUTTON   1
        #define CV_EVENT_FLAG_RBUTTON   2
        #define CV_EVENT_FLAG_MBUTTON   4
        #define CV_EVENT_FLAG_CTRLKEY   8
        #define CV_EVENT_FLAG_SHIFTKEY  16
        #define CV_EVENT_FLAG_ALTKEY    32

        CV_EXTERN_C_FUNCPTR( void (*CvMouseCallback )(int event,
                                                    int x,
                                                    int y,
                                                    int flags,
                                                    void* param) );

    Parameter:  windowName – Name of the window.

param onMouse:  Pointer to the function to be called every time a mouse event occurs in the specified window. This function should be prototyped as

void Foo(int event, int x, int y, int flags, void* param)¶

where event is one of CV_EVENT_*, x and y are the coordinates of the mouse pointer in image coordinates (not window coordinates), flags is a combination of CV_EVENT_FLAG, and param is a user-defined parameter passed to the cvSetMouseCallback() function call. :param param: User-defined parameter to be passed to the callback function.

The function cvSetMouseCallback() sets the callback function for mouse events occuring within the specified window. To see how it works, look at

http://opencvlibrary.sourceforge.net/../../samples/c/ffilldemo.c|opencv/samples/c/ffilldemo.c
SetTrackbarPos¶

void cvSetTrackbarPos(const char* trackbarName, const char* windowName, int pos)¶

    Sets the trackbar position.
    Parameters: 

        * trackbarName – Name of the trackbar.
        * windowName – Name of the window which is the parent of trackbar.
        * pos – New position.

    The function cvSetTrackbarPos() sets the position of the specified trackbar.

ShowImage¶

void cvShowImage(const char* name, const CvArr* image)¶

    Displays the image in the specified window
    Parameters: 

        * name – Name of the window.
        * image – Image to be shown.

    The function cvShowImage() displays the image in the specified window. If the window was created with the CV_WINDOW_AUTOSIZE flag then the image is shown with its original size, otherwise the image is scaled to fit in the window. The function may scale the image, depending on its depth:

        * If the image is 8-bit unsigned, it is displayed as is.
        * If the image is 16-bit unsigned or 32-bit integer, the pixels are divided by 256. That is, the value range [0,255*256] is mapped to [0,255].
        * If the image is 32-bit floating-point, the pixel values are multiplied by 255. That is, the value range [0,1] is mapped to [0,255].

WaitKey¶

int cvWaitKey(int delay=0)¶

    Waits for a pressed key.
    Parameter:  delay – Delay in milliseconds.

    The function cvWaitKey() waits for key event infinitely ($ \texttt{delay} <= 0$) or for delay milliseconds. Returns the code of the pressed key or -1 if no key was pressed before the specified time had elapsed.

    Note: This function is the only method in HighGUI that can fetch and handle events, so it needs to be called periodically for normal event processing, unless HighGUI is used within some environment that takes care of event processing.



void cvAdaptiveThreshold(const CvArr* src, CvArr* dst, double maxValue, int adaptive_method=CV_ADAPTIVE_THRESH_MEAN_C, int thresholdType=CV_THRESH_BINARY, int blockSize=3, double param1=5)¶

    Applies an adaptive threshold to an array.
    Parameters: 

        * src – Source image
        * dst – Destination image
        * maxValue – Maximum value that is used with CV_THRESH_BINARY and CV_THRESH_BINARY_INV
        * adaptive_method – Adaptive thresholding algorithm to use: CV_ADAPTIVE_THRESH_MEAN_C or CV_ADAPTIVE_THRESH_GAUSSIAN_C (see the discussion)
        * thresholdType –

          Thresholding type; must be one of
              o CV_THRESH_BINARY - xxx
              o CV_THRESH_BINARY_INV - xxx
        * blockSize – The size of a pixel neighborhood that is used to calculate a threshold value for the pixel: 3, 5, 7, and so on
        * param1 – The method-dependent parameter. For the methods CV_ADAPTIVE_THRESH_MEAN_C and CV_ADAPTIVE_THRESH_GAUSSIAN_C it is a constant subtracted from the mean or weighted mean (see the discussion), though it may be negative

    The function transforms a grayscale image to a binary image according to the formulas:

            *

              CV_THRESH_BINARY -

              dst(x,y) = \fork {\texttt{maxValue}}{if $src(x,y) > T(x,y)$}{0}{otherwise}
            *

              CV_THRESH_BINARY_INV -

              dst(x,y) = \fork {0}{if $src(x,y) > T(x,y)$}{\texttt{maxValue}}{otherwise}

    where $T(x,y)$ is a threshold calculated individually for each pixel.

    For the method CV_ADAPTIVE_THRESH_MEAN_C it is the mean of a $\texttt{blockSize} \times \texttt{blockSize}$ pixel neighborhood, minus param1.

    For the method CV_ADAPTIVE_THRESH_GAUSSIAN_C it is the weighted sum (gaussian) of a $\texttt{blockSize} \times \texttt{blockSize}$ pixel neighborhood, minus param1.

CvtColor¶

void cvCvtColor(const CvArr* src, CvArr* dst, int code)¶

    Converts an image from one color space to another.
    Parameters: 

        * src – The source 8-bit (8u), 16-bit (16u) or single-precision floating-point (32f) image
        * dst – The destination image of the same data type as the source. The number of channels may be different
        * code – Color conversion operation that can be specifed using CV_ *src_color_space* 2 *dst_color_space* constants (see below)

    The function converts the input image from one color space to another. The function ignores the colorModel and channelSeq fields of the IplImage header, so the source image color space should be specified correctly (including order of the channels in the case of RGB space. For example, BGR means 24-bit format with $B_0, G_0, R_0, B_1, G_1, R_1, ...$ layout whereas RGB means 24-format with $R_0, G_0, B_0, R_1, G_1, B_1, ...$ layout).

    The conventional range for R,G,B channel values is:

        * 0 to 255 for 8-bit images
        * 0 to 65535 for 16-bit images and
        * 0 to 1 for floating-point images.

    Of course, in the case of linear transformations the range can be specific, but in order to get correct results in the case of non-linear transformations, the input image should be scaled.

    The function can do the following transformations:

        * Transformations within RGB space like adding/removing the alpha channel, reversing the channel order, conversion to/from 16-bit RGB color (R5:G6:B5 or R5:G5:B5), as well as conversion to/from grayscale using:

    \text {RGB[A] to Gray:} Y \leftarrow 0.299 \cdot R + 0.587 \cdot G + 0.114 \cdot B

    and

    \text {Gray to RGB[A]:} R \leftarrow Y, G \leftarrow Y, B \leftarrow Y, A \leftarrow 0

    The conversion from a RGB image to gray is done with:

        cvCvtColor(src ,bwsrc, CV_RGB2GRAY)

        * RGB $\leftrightarrow $ CIE XYZ.Rec 709 with D65 white point (CV_BGR2XYZ, CV_RGB2XYZ, CV_XYZ2BGR, CV_XYZ2RGB):

    \begin{bmatrix} X \\ Y \\ Z \end{bmatrix} \leftarrow \begin{bmatrix} 0.412453 & 0.357580 & 0.180423 \\ 0.212671 & 0.715160 & 0.072169 \\ 0.019334 & 0.119193 & 0.950227 \end{bmatrix} \cdot \begin{bmatrix} R \\ G \\ B \end{bmatrix}

    \begin{bmatrix} R \\ G \\ B \end{bmatrix} \leftarrow \begin{bmatrix} 3.240479 & -1.53715 & -0.498535 \\ -0.969256 & 1.875991 & 0.041556 \\ 0.055648 & -0.204043 & 1.057311 \end{bmatrix} \cdot \begin{bmatrix} X \\ Y \\ Z \end{bmatrix}

    $X$, $Y$ and $Z$ cover the whole value range (in the case of floating-point images $Z$ may exceed 1).

        * RGB $\leftrightarrow $ YCrCb JPEG (a.k.a. YCC) (CV_BGR2YCrCb, CV_RGB2YCrCb, CV_YCrCb2BGR, CV_YCrCb2RGB)

    Y \leftarrow 0.299 \cdot R + 0.587 \cdot G + 0.114 \cdot B

    Cr \leftarrow (R-Y) \cdot 0.713 + delta

    Cb \leftarrow (B-Y) \cdot 0.564 + delta

    R \leftarrow Y + 1.403 \cdot (Cr - delta)

    G \leftarrow Y - 0.344 \cdot (Cr - delta) - 0.714 \cdot (Cb - delta)

    B \leftarrow Y + 1.773 \cdot (Cb - delta)

    where

    delta = \left\{ \begin{array}{l l} 128 & \mbox{for 8-bit images}\\ 32768 & \mbox{for 16-bit images}\\ 0.5 & \mbox{for floating-point images} \end{array} \right.

    Y, Cr and Cb cover the whole value range.

        * RGB $\leftrightarrow $ HSV (CV_BGR2HSV, CV_RGB2HSV, CV_HSV2BGR, CV_HSV2RGB) in the case of 8-bit and 16-bit images R, G and B are converted to floating-point format and scaled to fit the 0 to 1 range

    V \leftarrow max(R,G,B)

    S \leftarrow \fork {\frac{V-min(R,G,B)}{V}}{if $V \neq 0$}{0}{otherwise}

    H \leftarrow \forkthree {{60(G - B)}/{S}}{if $V=R$} {{120+60(B - R)}/{S}}{if $V=G$} {{240+60(R - G)}/{S}}{if $V=B$}

    if $H<0$ then $H \leftarrow H+360$

    On output $0 \leq V \leq 1$, $0 \leq S \leq 1$, $0 \leq H \leq 360$.

    The values are then converted to the destination data type:

        *

          8-bit images *

              V \leftarrow 255 V, S \leftarrow 255 S, H \leftarrow H/2 \text {(to fit to 0 to 255)}

        *

          16-bit images (currently not supported) *

              V <- 65535 V, S <- 65535 S, H <- H

        *

          32-bit images *

              H, S, V are left as is

        *

          RGB $\leftrightarrow $ HLS (CV_BGR2HLS, CV_RGB2HLS, CV_HLS2BGR, CV_HLS2RGB). in the case of 8-bit and 16-bit images R, G and B are converted to floating-point format and scaled to fit the 0 to 1 range.

    V_{max} \leftarrow {max}(R,G,B)

    V_{min} \leftarrow {min}(R,G,B)

    L \leftarrow \frac{V_{max} - V_{min}}{2}

    S \leftarrow \fork {\frac{V_{max} - V_{min}}{V_{max} + V_{min}}}{if $L < 0.5$} {\frac{V_{max} - V_{min}}{2 - (V_{max} + V_{min})}}{if $L \ge 0.5$}

    H \leftarrow \forkthree {{60(G - B)}/{S}}{if $V_{max}=R$} {{120+60(B - R)}/{S}}{if $V_{max}=G$} {{240+60(R - G)}/{S}}{if $V_{max}=B$}

    if $H<0$ then $H \leftarrow H+360$ On output $0 \leq V \leq 1$, $0 \leq S \leq 1$, $0 \leq H \leq 360$.

    The values are then converted to the destination data type:

        *

          8-bit images *

              V \leftarrow 255 V, S \leftarrow 255 S, H \leftarrow H/2 \text {(to fit to 0 to 255)}

        *

          16-bit images (currently not supported) *

              V <- 65535 V, S <- 65535 S, H <- H

        *

          32-bit images *

              H, S, V are left as is

        *

          RGB $\leftrightarrow $ CIE L*a*b* (CV_BGR2Lab, CV_RGB2Lab, CV_Lab2BGR, CV_Lab2RGB) in the case of 8-bit and 16-bit images R, G and B are converted to floating-point format and scaled to fit the 0 to 1 range

    \vecthree {X}{Y}{Z} \leftarrow \vecthreethree {0.412453}{0.357580}{0.180423} {0.212671}{0.715160}{0.072169} {0.019334}{0.119193}{0.950227} \cdot \vecthree {R}{G}{B}

    X \leftarrow X/X_ n, \text {where} X_ n = 0.950456

    Z \leftarrow Z/Z_ n, \text {where} Z_ n = 1.088754

    L \leftarrow \fork {116*Y^{1/3}-16}{for $Y>0.008856$} {903.3*Y}{for $Y \le 0.008856$}

    a \leftarrow 500 (f(X)-f(Y)) + delta

    b \leftarrow 200 (f(Y)-f(Z)) + delta

    where

    f(t)=\fork {t^{1/3}}{for $t>0.008856$} {7.787 t+16/116}{for $t<=0.008856$}

    and

    delta = \fork {128}{for 8-bit images}{0}{for floating-point images}

    On output $0 \leq L \leq 100$, $-127 \leq a \leq 127$, $-127 \leq b \leq 127$

    The values are then converted to the destination data type:

        *

          8-bit images *

              L \leftarrow L*255/100, a \leftarrow a + 128, b \leftarrow b + 128

        *

          16-bit images *

              currently not supported

        *

          32-bit images *

              L, a, b are left as is

        *

          RGB $\leftrightarrow $ CIE L*u*v* (CV_BGR2Luv, CV_RGB2Luv, CV_Luv2BGR, CV_Luv2RGB) in the case of 8-bit and 16-bit images R, G and B are converted to floating-point format and scaled to fit 0 to 1 range

    \vecthree {X}{Y}{Z} \leftarrow \vecthreethree {0.412453}{0.357580}{0.180423} {0.212671}{0.715160}{0.072169} {0.019334}{0.119193}{0.950227} \cdot \vecthree {R}{G}{B}

    L \leftarrow \fork {116 Y^{1/3}}{for $Y>0.008856$} {903.3 Y}{for $Y<=0.008856$}

    u' \leftarrow 4*X/(X + 15*Y + 3 Z)

    v' \leftarrow 9*Y/(X + 15*Y + 3 Z)

    u \leftarrow 13*L*(u' - u_ n) \quad \text {where} \quad u_ n=0.19793943

    v \leftarrow 13*L*(v' - v_ n) \quad \text {where} \quad v_ n=0.46831096

    On output $0 \leq L \leq 100$, $-134 \leq u \leq 220$, $-140 \leq v \leq 122$.

    The values are then converted to the destination data type:

        *

          8-bit images *

              L \leftarrow 255/100 L, u \leftarrow 255/354 (u + 134), v \leftarrow 255/256 (v + 140)

        *

          16-bit images *

              currently not supported

        *

          32-bit images *

              L, u, v are left as is

    The above formulas for converting RGB to/from various color spaces have been taken from multiple sources on Web, primarily from the Ford98 at the Charles Poynton site.

        * Bayer $\rightarrow $ RGB (CV_BayerBG2BGR, CV_BayerGB2BGR, CV_BayerRG2BGR, CV_BayerGR2BGR, CV_BayerBG2RGB, CV_BayerGB2RGB, CV_BayerRG2RGB, CV_BayerGR2RGB) The Bayer pattern is widely used in CCD and CMOS cameras. It allows one to get color pictures from a single plane where R,G and B pixels (sensors of a particular component) are interleaved like this:

    \definecolor {BackGray}{rgb}{0.8,0.8,0.8} \begin{array}{ c c c c c } \color {red}R& \color {green}G& \color {red}R& \color {green}G& \color {red}R\\ \color {green}G& \colorbox {BackGray}{\color {blue}B}& \colorbox {BackGray}{\color {green}G}& \color {blue}B& \color {green}G\\ \color {red}R& \color {green}G& \color {red}R& \color {green}G& \color {red}R\\ \color {green}G& \color {blue}B& \color {green}G& \color {blue}B& \color {green}G\\ \color {red}R& \color {green}G& \color {red}R& \color {green}G& \color {red}R\end{array}

    The output RGB components of a pixel are interpolated from 1, 2 or 4 neighbors of the pixel having the same color. There are several modifications of the above pattern that can be achieved by shifting the pattern one pixel left and/or one pixel up. The two letters $C_1$ and $C_2$ in the conversion constants CV_Bayer $ C_1 C_2 $ 2BGR and CV_Bayer $ C_1 C_2 $ 2RGB indicate the particular pattern type - these are components from the second row, second and third columns, respectively. For example, the above pattern has very popular “BG” type.

DistTransform¶

void cvDistTransform(const CvArr* src, CvArr* dst, int distance_type=CV_DIST_L2, int mask_size=3, const float* mask=NULL, CvArr* labels=NULL)¶

    Calculates the distance to the closest zero pixel for all non-zero pixels of the source image.
    Parameters: 

        * src – 8-bit, single-channel (binary) source image
        * dst – Output image with calculated distances (32-bit floating-point, single-channel)
        * distance_type – Type of distance; can be CV_DIST_L1, CV_DIST_L2, CV_DIST_C or CV_DIST_USER
        * mask_size – Size of the distance transform mask; can be 3 or 5. in the case of CV_DIST_L1 or CV_DIST_C the parameter is forced to 3, because a $3\times 3$ mask gives the same result as a $5\times 5 $ yet it is faster
        * mask – User-defined mask in the case of a user-defined distance, it consists of 2 numbers (horizontal/vertical shift cost, diagonal shift cost) in the case ofa $3\times 3$ mask and 3 numbers (horizontal/vertical shift cost, diagonal shift cost, knight’s move cost) in the case of a $5\times 5$ mask
        * labels – The optional output 2d array of integer type labels, the same size as src and dst

    The function calculates the approximated distance from every binary image pixel to the nearest zero pixel. For zero pixels the function sets the zero distance, for others it finds the shortest path consisting of basic shifts: horizontal, vertical, diagonal or knight’s move (the latest is available for a $5\times 5$ mask). The overall distance is calculated as a sum of these basic distances. Because the distance function should be symmetric, all of the horizontal and vertical shifts must have the same cost (that is denoted as a), all the diagonal shifts must have the same cost (denoted b), and all knight’s moves must have the same cost (denoted c). For CV_DIST_C and CV_DIST_L1 types the distance is calculated precisely, whereas for CV_DIST_L2 (Euclidian distance) the distance can be calculated only with some relative error (a $5\times 5$ mask gives more accurate results), OpenCV uses the values suggested in :
    CV_DIST_C   $(3\times 3)$   a = 1, b = 1
    CV_DIST_L1  $(3\times 3)$   a = 1, b = 2
    CV_DIST_L2  $(3\times 3)$   a=0.955, b=1.3693
    CV_DIST_L2  $(5\times 5)$   a=1, b=1.4, c=2.1969

    And below are samples of the distance field (black (0) pixel is in the middle of white square) in the case of a user-defined distance:

    User-defined $3 \times 3$ mask (a=1, b=1.5)
    4.5   4   3.5   3   3.5   4   4.5
    4   3   2.5   2   2.5   3   4
    3.5   2.5   1.5   1   1.5   2.5   3.5
    3   2   1       1   2   3
    3.5   2.5   1.5   1   1.5   2.5   3.5
    4   3   2.5   2   2.5   3   4
    4.5   4   3.5   3   3.5   4   4.5

    User-defined $5 \times 5$ mask (a=1, b=1.5, c=2)
    4.5   3.5   3   3   3   3.5   4.5
    3.5   3   2   2   2   3   3.5
    3   2   1.5   1   1.5   2   3
    3   2   1       1   2   3
    3   2   1.5   1   1.5   2   3
    3.5   3   2   2   2   3   3.5
    4   3.5   3   3   3   3.5   4

    Typically, for a fast, coarse distance estimation CV_DIST_L2, a $3\times 3$ mask is used, and for a more accurate distance estimation CV_DIST_L2, a $5\times 5$ mask is used.

    When the output parameter labels is not NULL, for every non-zero pixel the function also finds the nearest connected component consisting of zero pixels. The connected components themselves are found as contours in the beginning of the function.

    In this mode the processing time is still O(N), where N is the number of pixels. Thus, the function provides a very fast way to compute approximate Voronoi diagram for the binary image.

FloodFill¶

void cvFloodFill(CvArr* image, CvPoint seed_point, CvScalar new_val, CvScalar lo_diff=cvScalarAll(0), CvScalar up_diff=cvScalarAll(0), CvConnectedComp* comp=NULL, int flags=4, CvArr* mask=NULL)¶

    Fills a connected component with the given color.

        typedef struct CvConnectedComp
        {
            double area;     area of the segmented component
            CvScalar value;  average color of the connected component 
            CvRect rect;     ROI of the segmented component 
            CvSeq* contour;  optional component boundary
                              (the contour might have child contours corresponding to the holes) 
        } CvConnectedComp;

        #define CV_FLOODFILL_FIXED_RANGE (1 << 16)
        #define CV_FLOODFILL_MASK_ONLY   (1 << 17)

    Parameters: 

        * image – Input 1- or 3-channel, 8-bit or floating-point image. It is modified by the function unless the CV_FLOODFILL_MASK_ONLY flag is set (see below)
        * seed_point – The starting point
        * new_val – New value of the repainted domain pixels
        * lo_diff – Maximal lower brightness/color difference between the currently observed pixel and one of its neighbors belonging to the component, or a seed pixel being added to the component. In the case of 8-bit color images it is a packed value
        * up_diff – Maximal upper brightness/color difference between the currently observed pixel and one of its neighbors belonging to the component, or a seed pixel being added to the component. In the case of 8-bit color images it is a packed value
        * comp – Pointer to the structure that the function fills with the information about the repainted domain
        * flags –

          The operation flags. Lower bits contain connectivity value, 4 (by default) or 8, used within the function. Connectivity determines which neighbors of a pixel are considered. Upper bits can be 0 or a combination of the following flags:
              o CV_FLOODFILL_FIXED_RANGE - if set, the difference between the current pixel and seed pixel is considered, otherwise the difference between neighbor pixels is considered (the range is floating)
              o CV_FLOODFILL_MASK_ONLY - if set, the function does not fill the image (new_val is ignored), but fills the mask (that must be non-NULL in this case)
        * mask – Operation mask, should be a single-channel 8-bit image, 2 pixels wider and 2 pixels taller than image. If not NULL, the function uses and updates the mask, so the user takes responsibility of initializing the mask content. Floodfilling can’t go across non-zero pixels in the mask, for example, an edge detector output can be used as a mask to stop filling at edges. It is possible to use the same mask in multiple calls to the function to make sure the filled area do not overlap. Note: because the mask is larger than the filled image, a pixel in mask that corresponds to $(x,y)$ pixel in image will have coordinates $(x+1,y+1)$

    The function fills a connected component starting from the seed point with the specified color. The connectivity is determined by the closeness of pixel values. The pixel at $(x,y)$ is considered to belong to the repainted domain if:

        *

          grayscale image, floating range *

              src(x',y')-\texttt{lo\_ diff} <= src(x,y) <= src(x',y')+\texttt{up\_ diff}

        *

          grayscale image, fixed range *

              src(seed.x,seed.y)-\texttt{lo\_ diff}<=src(x,y)<=src(seed.x,seed.y)+\texttt{up\_ diff}

        *

          color image, floating range *

              src(x',y')_ r-\texttt{lo\_ diff}_ r<=src(x,y)_ r<=src(x',y')_ r+\texttt{up\_ diff}_ r

              src(x',y')_ g-\texttt{lo\_ diff}_ g<=src(x,y)_ g<=src(x',y')_ g+\texttt{up\_ diff}_ g

              src(x',y')_ b-\texttt{lo\_ diff}_ b<=src(x,y)_ b<=src(x',y')_ b+\texttt{up\_ diff}_ b

        *

          color image, fixed range *

              src(seed.x,seed.y)_ r-\texttt{lo\_ diff}_ r<=src(x,y)_ r<=src(seed.x,seed.y)_ r+\texttt{up\_ diff}_ r

              src(seed.x,seed.y)_ g-\texttt{lo\_ diff}_ g<=src(x,y)_ g<=src(seed.x,seed.y)_ g+\texttt{up\_ diff}_ g

              src(seed.x,seed.y)_ b-\texttt{lo\_ diff}_ b<=src(x,y)_ b<=src(seed.x,seed.y)_ b+\texttt{up\_ diff}_ b

    where $src(x’,y’)$ is the value of one of pixel neighbors. That is, to be added to the connected component, a pixel’s color/brightness should be close enough to the:

        * color/brightness of one of its neighbors that are already referred to the connected component in the case of floating range
        * color/brightness of the seed point in the case of fixed range.

Inpaint¶

void cvInpaint(const CvArr* src, const CvArr* mask, CvArr* dst, double inpaintRadius, int flags)¶

    Inpaints the selected region in the image.
    Parameters: 

        * src – The input 8-bit 1-channel or 3-channel image.
        * mask – The inpainting mask, 8-bit 1-channel image. Non-zero pixels indicate the area that needs to be inpainted.
        * dst – The output image of the same format and the same size as input.
        * inpaintRadius – The radius of circlular neighborhood of each point inpainted that is considered by the algorithm.
        * flags –

          The inpainting method, one of the following:
              o CV_INPAINT_NS - Navier-Stokes based method.
              o CV_INPAINT_TELEA - The method by Alexandru Telea bgroup({# Telea04})bgroup({[Telea04]})

    The function reconstructs the selected image area from the pixel near the area boundary. The function may be used to remove dust and scratches from a scanned photo, or to remove undesirable objects from still images or video.

Integral¶

void cvIntegral(const CvArr* image, CvArr* sum, CvArr* sqsum=NULL, CvArr* tiltedSum=NULL)¶

    Calculates the integral of an image.
    Parameters: 

        * image – The source image, $W\times H$, 8-bit or floating-point (32f or 64f)
        * sum – The integral image, $(W+1)\times (H+1)$, 32-bit integer or double precision floating-point (64f)
        * sqsum – The integral image for squared pixel values, $(W+1)\times (H+1)$, double precision floating-point (64f)
        * tiltedSum – The integral for the image rotated by 45 degrees, $(W+1)\times (H+1)$, the same data type as sum

    The function calculates one or more integral images for the source image as following:

    \texttt{sum}(X,Y) = \sum _{x<X,y<Y} \texttt{image}(x,y)

    \texttt{sqsum}(X,Y) = \sum _{x<X,y<Y} \texttt{image}(x,y)^2

    \texttt{tiltedSum}(X,Y) = \sum _{y<Y,abs(x-X)<y} \texttt{image}(x,y)

    Using these integral images, one may calculate sum, mean and standard deviation over a specific up-right or rotated rectangular region of the image in a constant time, for example:

    \sum _{x_1<=x<x_2, \, y_1<=y<y_2} = \texttt{sum}(x_2,y_2)-\texttt{sum}(x_1,y_2)-\texttt{sum}(x_2,y_1)+\texttt{sum}(x_1,x_1)

    It makes possible to do a fast blurring or fast block correlation with variable window size, for example. In the case of multi-channel images, sums for each channel are accumulated independently.

PyrMeanShiftFiltering¶

void cvPyrMeanShiftFiltering(const CvArr* src, CvArr* dst, double sp, double sr, int max_level=1, CvTermCriteria termcrit=cvTermCriteria(CV_TERMCRIT_ITER+CV_TERMCRIT_EPS, 5, 1))¶

    Does meanshift image segmentation
    Parameters: 

        * src – The source 8-bit, 3-channel image.
        * dst – The destination image of the same format and the same size as the source.
        * sp – The spatial window radius.
        * sr – The color window radius.
        * max_level – Maximum level of the pyramid for the segmentation.
        * termcrit – Termination criteria: when to stop meanshift iterations.

    The function implements the filtering stage of meanshift segmentation, that is, the output of the function is the filtered “posterized” image with color gradients and fine-grain texture flattened. At every pixel $(X,Y)$ of the input image (or down-sized input image, see below) the function executes meanshift iterations, that is, the pixel $(X,Y)$ neighborhood in the joint space-color hyperspace is considered:

    (x,y): X-\texttt{sp} \le x \le X+\texttt{sp} , Y-\texttt{sp} \le y \le Y+\texttt{sp} , ||(R,G,B)-(r,g,b)|| \le \texttt{sr}

    where (R,G,B) and (r,g,b) are the vectors of color components at (X,Y) and (x,y), respectively (though, the algorithm does not depend on the color space used, so any 3-component color space can be used instead). Over the neighborhood the average spatial value (X',Y') and average color vector (R',G',B') are found and they act as the neighborhood center on the next iteration:

    $(X,Y)~ (X’,Y’), (R,G,B)~ (R’,G’,B’).$

    After the iterations over, the color components of the initial pixel (that is, the pixel from where the iterations started) are set to the final value (average color at the last iteration):

    $I(X,Y) <- (R*,G*,B*)$

    Then $\texttt{max\_ level}>0$ , the gaussian pyramid of $\texttt{max\_ level}+1$ levels is built, and the above procedure is run on the smallest layer. After that, the results are propagated to the larger layer and the iterations are run again only on those pixels where the layer colors differ much ( $>\texttt{sr}$ ) from the lower-resolution layer, that is, the boundaries of the color regions are clarified. Note, that the results will be actually different from the ones obtained by running the meanshift procedure on the whole original image (i.e. when $\texttt{max\_ level}==0$ ).

PyrSegmentation¶

void cvPyrSegmentation(IplImage* src, IplImage* dst, CvMemStorage* storage, CvSeq** comp, int level, double threshold1, double threshold2)¶

    Implements image segmentation by pyramids.
    Parameters: 

        * src – The source image
        * dst – The destination image
        * storage – Storage; stores the resulting sequence of connected components
        * comp – Pointer to the output sequence of the segmented components
        * level – Maximum level of the pyramid for the segmentation
        * threshold1 – Error threshold for establishing the links
        * threshold2 – Error threshold for the segments clustering

    The function implements image segmentation by pyramids. The pyramid builds up to the level level. The links between any pixel a on level i and its candidate father pixel b on the adjacent level are established if $p(c(a),c(b))<threshold1$. After the connected components are defined, they are joined into several clusters. Any two segments A and B belong to the same cluster, if $p(c(A),c(B))<threshold2$. If the input image has only one channel, then $p(c^1,c^2)=|c^1-c^2|$. If the input image has three channels (red, green and blue), then

    p(c^1,c^2) = 0.30 (c^1_ r - c^2_ r) + 0.59 (c^1_ g - c^2_ g) + 0.11 (c^1_ b - c^2_ b).

    There may be more than one connected component per a cluster. The images src and dst should be 8-bit single-channel or 3-channel images or equal size.

Threshold¶

double cvThreshold(const CvArr* src, CvArr* dst, double threshold, double maxValue, int thresholdType)¶
*/




