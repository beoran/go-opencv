// Created by cgo - DO NOT EDIT
package opencv

import "unsafe"

type _ unsafe.Pointer

type _C_IplImage _Cstruct__IplImage
type _C_int int32
type _C_char int8
type _Cstruct__IplImage struct {
	nSize		_C_int
	ID		_C_int
	nChannels	_C_int
	alphaChannel	_C_int
	depth		_C_int
	colorModel	[4]_C_char
	channelSeq	[4]_C_char
	dataOrder	_C_int
	origin		_C_int
	align		_C_int
	width		_C_int
	height		_C_int
	roi		*_Cstruct__IplROI
	maskROI		*_Cstruct__IplImage
	imageId		unsafe.Pointer
	tileInfo	*[0]byte
	imageSize	_C_int
	imageData	*_C_char
	widthStep	_C_int
	BorderMode	[4]_C_int
	BorderConst	[4]_C_int
	imageDataOrigin	*_C_char
}
type _Cstruct__IplROI struct {
	coi	_C_int
	xOffset	_C_int
	yOffset	_C_int
	width	_C_int
	height	_C_int
}
type _C_void [0]byte

func _C_GoString(*_C_char) string
func _C_cvSetErrMode(_C_int) _C_int
func _C_CString(string) *_C_char
func _C_cvGetErrStatus() _C_int
func _C_cvLoadImage(*_C_char, _C_int) *_C_IplImage
func _C_cvError(_C_int, *_C_char, *_C_char, *_C_char, _C_int)
func _C_cvErrorStr(_C_int) *_C_char
func _C_free(unsafe.Pointer)
func _C_cvSetErrStatus(_C_int)
func _C_cvReleaseImage(**_C_IplImage)
func _C_cvGetErrMode() _C_int
func _C_cvSaveImage(*_C_char, unsafe.Pointer, *_C_int) _C_int
