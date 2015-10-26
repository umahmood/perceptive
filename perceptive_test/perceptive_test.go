package perceptive_test

import (
	"image"
	"log"
	"testing"

	"github.com/disintegration/imaging"
	"github.com/umahmood/perceptive"

	_ "image/jpeg"
)

const testDir = "test_images/"

func openImage(fileName string) image.Image {
	img, err := imaging.Open(testDir + fileName)
	if err != nil {
		log.Fatalln(err)
	}
	return img
}

var testsAverage = []struct {
	imgA image.Image
	imgB image.Image
	dist int
}{
	{ // test: identical images
		imgA: openImage("lena.jpg"),
		imgB: openImage("lena.jpg"),
		dist: 0,
	},
	{ // test: similar images
		imgA: openImage("lena.jpg"),
		imgB: openImage("lena_pink.jpg"),
		dist: 8,
	},
	{ // test: different images
		imgA: openImage("lena.jpg"),
		imgB: openImage("rainbow_flowers.jpg"),
		dist: 27,
	},
}

var testsDifference = []struct {
	imgA image.Image
	imgB image.Image
	dist int
}{
	{ // test: identical images
		imgA: openImage("lena.jpg"),
		imgB: openImage("lena.jpg"),
		dist: 0,
	},
	{ // test: similar images
		imgA: openImage("toy_story_1.jpg"),
		imgB: openImage("toy_story_2.jpg"),
		dist: 0,
	},
	{ // test: different images
		imgA: openImage("lena.jpg"),
		imgB: openImage("homer_doh.jpg"),
		dist: 31,
	},
}

func TestCompareImagesWithAverageHash(t *testing.T) {
	for _, x := range testsAverage {
		got, err := perceptive.CompareImages(x.imgA, x.imgB, perceptive.Average)
		if err != nil {
			t.Errorf("CompareImages: %s", err.Error())
		}

		if got != x.dist {
			t.Errorf("CompareImages: distance got %d want %d", got, x.dist)
		}
	}
}

func TestCompareImagesWithDifferenceHash(t *testing.T) {
	for _, x := range testsDifference {
		got, err := perceptive.CompareImages(x.imgA, x.imgB, perceptive.Difference)
		if err != nil {
			t.Errorf("CompareImages: %s", err.Error())
		}

		if got != x.dist {
			t.Errorf("CompareImages: distance got %d want %d", got, x.dist)
		}
	}
}

func TestAhash(t *testing.T) {
	img := openImage("rainbow_flowers.jpg")

	var want uint64 = 4069139135820476416

	got, err := perceptive.Ahash(img)

	if err != nil {
		t.Errorf("Ahash: error %s", err.Error())
	}

	if got != want {
		t.Errorf("Ahash: got %d want %d", got, want)
	}
}

func TestDhash(t *testing.T) {
	img := openImage("rainbow_flowers.jpg")

	var want uint64 = 17352528693468127320

	got, err := perceptive.Dhash(img)

	if err != nil {
		t.Errorf("Dhash: error %s", err.Error())
	}

	if got != want {
		t.Errorf("Dhash: got %d want %d", got, want)
	}
}

var testsHamming = []struct {
	a    uint64
	b    uint64
	dist int
}{
	{
		a:    171798691,
		b:    536870911,
		dist: 15,
	},
	{
		a:    171798691,
		b:    42,
		dist: 13,
	},
	{
		a:    6,
		b:    21,
		dist: 3,
	},
	{
		a:    42,
		b:    42,
		dist: 0,
	},
	{
		a:    13280149282866979154,
		b:    17352528693468127320,
		dist: 18,
	},
}

func TestHammingDistance(t *testing.T) {
	for _, x := range testsHamming {
		got := perceptive.HammingDistance(x.a, x.b)

		if got != x.dist {
			t.Errorf("HammingDistance: got %d want %d", got, x.dist)
		}
	}
}
