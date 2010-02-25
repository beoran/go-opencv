# Copyright 2010 Beoran. 
# 
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

all: libs test-opencv

libs:
	make -C opencv install

test-opencv: test-opencv.go libs
	$(GC) test-opencv.go
	$(LD) -o $@ test-opencv.$(O)

clean:
	make -C opencv clean
	rm -f -r *.8 *.6 *.o */*.8 */*.6 */*.o */_obj test-opencv
	