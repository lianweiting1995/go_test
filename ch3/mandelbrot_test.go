package ch3

import "testing"

func TestMandelBrot(t *testing.T) {
	MandelbrotMain()
}

func TestMandelBrotServer(t *testing.T) {
	HttpServer()
}
