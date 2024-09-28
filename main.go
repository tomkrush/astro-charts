package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

// Main catalog header
type CatalogHeader struct {
	Star0 int32
	Star1 int32
	StarN int32
	StNum int32
	MProp int32
	NMag  int32
	NBEnt int32
}

func (c CatalogHeader) numberOfStars() int {
	return int(math.Abs(float64(c.StarN)))
}

// Each star entry
type StarEntry struct {
	CatalogNumber int32   // Optional
	SRA0          float64 // Right Ascension (radians)
	SDEC0         float64 // Declination (radians)
	ISP           string  // Spectral type (2 characters)
	Magnitude     []int16 // V Magnitude * 100
	XRPM          float32 // R.A. proper motion (radians/year)
	XDPM          float32 // Dec. proper motion (radians/year)
	SVEL          float64 // Radial velocity (km/s)
	ObjectName    string  // Optional, length -StNum
}

func main() {
	// Open the file
	file, err := os.Open("data/BSC5")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// read the first 28 bytes
	header := CatalogHeader{}
	err = binary.Read(file, binary.LittleEndian, &header)

	for i := 0; i < header.numberOfStars(); i++ {
		starBytes := make([]byte, header.NBEnt)
		_, err = file.Read(starBytes)

		// star := StarEntry{}

		fmt.Println(starBytes)

		// reader := bytes.NewReader(starBytes)
		// err = binary.Read(reader, binary.LittleEndian, &star)

		// fmt.Println(star)

		break
	}
}
