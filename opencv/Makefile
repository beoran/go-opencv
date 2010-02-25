# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=opencv

# GOFILES:=constants.$(O).go

CGOFILES:=opencv.go

CGO_CFLAGS:=-I/usr/local/include/opencv -I/usr/include/opencv

CGO_LDFLAGS:=-L/usr/local/lib -lcxcore -lcv -lhighgui -lcallback

CLEANFILES+=opencv

include $(GOROOT)/src/Make.pkg

constants.$(O).go: constants.c
	godefs -g opencv constants.c > constants.$(O).go
	gofmt -w constants.$(O).go