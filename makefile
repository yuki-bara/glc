# SPDX-License-Identifier: 0BSD
# Author: Makkhawan Sardlah

CMD = go
ACT = build
NAM = glc
VERSION=1.1

build :
	mkdir -p bin
	$(CMD) $(ACT) -ldflags="-X 'main.version=$(VERSION)'" -o bin/$(NAM) src/main.go

build_test :
	mkdir -p tests
	$(CMD) $(ACT) -ldflags="-X 'main.version=$(VERSION)'" -o tests/$(NAM) src/main.go

clean :
	rm -r bin