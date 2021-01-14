package main

import {
	"fmt"
}

/*
#cgo LDFLAGS: -L./ -lArrowToolsStaticLib -lATLS
#include "ATSL.h"
*/

import "C"


func parseFixedFileToParquet() {
	C.parseFixedFileToParquet()
	return
}
