package perceptive

import (
	"errors"
	"image"
	"image/color"

	"github.com/disintegration/imaging"
	"github.com/umahmood/iter"
)

// Errors
var (
	ErrNilImage    = errors.New("perceptive: nil image")
	ErrInvalidHash = errors.New("perceptive: invalid perceptual hash")
)

// PerceptualHash perceptual hash algorithms
type PerceptualHash uint8

// Valid perceptual hashes
const (
	Average    PerceptualHash = iota // Average hash (Ahash)
	Difference                       // Difference hash (Dhash)
)

// grayscale produces the grayscale version of the image
func grayscale(img image.Image) image.Image {
	return imaging.Grayscale(img)
}

// resize resizes an image to the specified width and height using the
// Lanczos filter.
func resize(img image.Image, width, height int) image.Image {
	return imaging.Resize(img, width, height, imaging.Lanczos)
}

// sumPixels sums the values of three 1 byte values
func sumPixels(r, g, b uint8) uint64 {
	return uint64(r) + uint64(g) + uint64(b)
}

// Ahash implements the average hash algorithm.
func Ahash(img image.Image) (uint64, error) {
	if img == nil {
		return 0, ErrNilImage
	}

	const hashSize = 8

	pImg := grayscale(img)
	pImg = resize(pImg, hashSize, hashSize)

	// calc. average of the colors.
	var total uint64
	for row := range iter.N(hashSize) {
		for col := range iter.N(hashSize) {
			pixelVal := pImg.At(col, row).(color.NRGBA)
			total = total + sumPixels(pixelVal.R, pixelVal.G, pixelVal.B)
		}
	}

	avg := (total / (hashSize * hashSize))

	var hash uint64
	var pos uint64

	for row := range iter.N(hashSize) {
		for col := range iter.N(hashSize) {
			pixelVal := pImg.At(col, row).(color.NRGBA)
			sum := sumPixels(pixelVal.R, pixelVal.G, pixelVal.B)
			if sum > avg {
				hash |= (1 << pos)
			}
			pos++
		}
	}

	return hash, nil
}

// Dhash implements the difference hash algorithm.
func Dhash(img image.Image) (uint64, error) {
	if img == nil {
		return 0, ErrNilImage
	}

	const hashSize = 8

	pImg := grayscale(img)
	pImg = resize(pImg, hashSize+1, hashSize)

	// comparePixels compares the values of two pixels, returns true if px1 is
	// greater than px2.
	comparePixels := func(px1, px2 color.NRGBA) bool {
		x := sumPixels(px1.R, px1.G, px1.B)
		y := sumPixels(px2.R, px2.G, px2.B)
		if x > y {
			return true
		}
		return false
	}

	var hash uint64
	var pos uint64

	for row := range iter.N(hashSize) {
		for col := range iter.N(hashSize) {
			pixelLeft := pImg.At(col, row).(color.NRGBA)
			pixelRight := pImg.At(col+1, row).(color.NRGBA)
			if comparePixels(pixelLeft, pixelRight) {
				hash |= (1 << pos)
			}
			pos++
		}
	}

	return hash, nil
}

// HammingDistance computes the Hamming distance of two integers.
func HammingDistance(x, y uint64) int {
	dist := 0
	val := x ^ y
	for val != 0 {
		dist++
		val = val & (val - 1)
	}
	return dist
}

// CompareImages compares two images using the given perceptual hash algorithm.
// The function returns a distance value indicating how similar the two images
// are. The distance value differs depending on the perceptual hash algorithm
// used. If perceptual hash type is Average or Difference:
//
// - A distance of 0 indicates the same hash and likely a similar image.
//
// - A distance between 1 and 10 the image is potentially a variation.
//
// - A distance greater than 10 the image is likely a different image.
func CompareImages(img1, img2 image.Image, hash PerceptualHash) (int, error) {
	var a uint64
	var b uint64
	var err error

	switch hash {
	case Average:
		a, err = Ahash(img1)
		b, err = Ahash(img2)
	case Difference:
		a, err = Dhash(img1)
		b, err = Dhash(img2)
	default:
		return -1, ErrInvalidHash
	}

	if err != nil {
		return -1, err
	}

	return HammingDistance(a, b), nil
}
