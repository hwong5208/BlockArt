/*

A trivial application to illustrate how the blockartlib library can be
used from an application in project 1 for UBC CS 416 2017W2.

Usage:
go run art-app.go
*/

package main

// Expects blockartlib.go to be in the ./blockartlib/ dir, relative to
// this art-app.go file



import (
	//"crypto/x509"
	"./blockartlib"
	"fmt"
	"os"

	"crypto/x509"
	"encoding/gob"
	"net"
	"crypto/elliptic"
)
//import "./blockartlib"
func main() {
	args := os.Args[1:]
	gob.Register(&net.TCPAddr{})
	gob.Register(&elliptic.CurveParams{})
	minerAddr := args[0]
	keyfile := args[1]


	buf := make([]byte, 200)
	file, err := os.OpenFile("priv" + keyfile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Seek(0, 0)
	size,err := file.Read(buf)
	fmt.Println(err)
	buf = buf[:size]
	privKey,err := x509.ParseECPrivateKey(buf)
	fmt.Println(err)


	// Open a canvas.
	canvas, settings, err := blockartlib.OpenCanvas(minerAddr, *privKey)
	fmt.Println(canvas,settings)
	if checkError(err) != nil {
		return
	}

    validateNum := 2


    starSVG := "M 100 100 l 50 100"

	// Add a line.
	shapeHash, blockHash, ink, err := canvas.AddShape(uint8(validateNum), blockartlib.PATH, starSVG, "transparent", "yellow")
	fmt.Println("result of add shape 1")
	fmt.Println(shapeHash,blockHash,ink,err)

	// Add another line. Expect receive
	shapeHash2, blockHash2, ink2, err := canvas.AddShape(uint8(validateNum), blockartlib.PATH, "M 250 300 h 101", "transparent", "yellow")
	fmt.Println("result of add shape 2")
	fmt.Println(shapeHash2, blockHash2, ink2, err)



   shapes,err := canvas.GetShapes(blockHash)
   fmt.Println("result of get shapes")
   fmt.Println(shapes)

	// Delete the first line.


	//ink4, err := canvas.DeleteShape(uint8(validateNum), shapeHash)
	//fmt.Println("result of delete shape 1")
	//fmt.Println(ink4,err)
	//ink4, err = canvas.DeleteShape(uint8(validateNum), shapeHash2)
	//fmt.Println("result of delete shape 2")
	//fmt.Println(ink4,err)



	// assert ink3 > ink2
	svgstring, err := canvas.GetSvgString(shapeHash)
	fmt.Println("result of get svg string")
	fmt.Println(svgstring)
	// Close the canvas.
	ink5, err := canvas.CloseCanvas()
	fmt.Println("result of close canvas")
	fmt.Println(ink5,err)

}

// If error is non-nil, print it out and return it.
func checkError(err error) error {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ", err.Error())
		return err
	}
	return nil
}
