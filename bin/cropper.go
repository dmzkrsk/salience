/*
  PUBLIC DOMAIN STATEMENT
  To the extent possible under law, Ian Davis has waived all copyright
  and related or neighboring rights to this Source Code file.
  This work is published from the United Kingdom.
*/

package main

import (
	salience ".."
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"strconv"
)

// A simple command line program for finding the most interesting section of an image
func main() {
	if len(os.Args) < 4 {
		println("Please supply input image filename, output filename and output image width as arguments")
		os.Exit(1)
	}
	finName := os.Args[1]
	foutName := os.Args[2]
	widthStr := os.Args[3]

	width, err := strconv.ParseInt(widthStr, 10, 0)
	if err != nil {
		fmt.Printf("Error parsing image width argument: %s\n", err.Error())
		os.Exit(1)
	}

	fin, err := os.OpenFile(finName, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("Error reading input image: %s\n", err.Error())
		os.Exit(1)
	}

	img, _, err := image.Decode(fin)
	if err != nil {
		fmt.Printf("Error decoding input image: %s\n", err.Error())
		os.Exit(1)
	}

	fout, err := os.OpenFile(foutName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Error writing output image: %s\n", err.Error())
		os.Exit(1)
	}

	imgOut := salience.Crop(img, int(width), int(width))

	if err = png.Encode(fout, imgOut); err != nil {
		fmt.Printf("Error encoding output image: %s\n", err.Error())
		os.Exit(1)
	}
}
