# Perceptive

Perceptive is a Go library which implements perceptual hash algorithms for 
comparing images.

Perceptual hash algorithms are a family of comparable hash functions which
generate distinct (but not unique) fingerprints, these fingerprints are then
comparable.

Perceptual hash algorithms are mainly used for detecting duplicates of the same 
files, in a way that standard and cryptographic hashes generally fail.

**Note:** This library can only compute hashes for images, it does not work on 
audio or video files.

Currently, the following perceptual hash algorithms are implemented:

- **Average Hash (Ahash)** - Fast but generates a huge number of false positives.
- **Difference Hash (Dhash)** - Fast and very few false positives.

**Perceptual hash algorithms can give false positives**, but there main aim is to 
give you a sense of similarity between files.

Perceptual hash algorithms tend to return a distance score. When comparing the 
two identical images below, we would receive a distance of 0:

<img src="https://github.com/umahmood/perceptive/blob/master/perceptive_test/test_images/rainbow_flowers.jpg" with="310" height="300"/>
<img src="https://github.com/umahmood/perceptive/blob/master/perceptive_test/test_images/rainbow_flowers.jpg" with="310" height="300"/>

A distance of zero means that the images are **likely** the same.

When comparing the two similar images below we would receive a distance between 
1-10 (depending on the hashing technique used):

<img src="https://github.com/umahmood/perceptive/blob/master/perceptive_test/test_images/toy_story_1.jpg" width="440" height="290"/>
<img src="https://github.com/umahmood/perceptive/blob/master/perceptive_test/test_images/toy_story_2.jpg" width="440" height="290"/>

A distance between 1-10 indicates the images are **likely a variation** of each 
other.

When comparing the two different images below we would receive a distance greater 
than 10:

<img src="https://github.com/umahmood/perceptive/blob/master/perceptive_test/test_images/homer_doh.jpg" width="320" height="240"/>
<img src="https://github.com/umahmood/perceptive/blob/master/perceptive_test/test_images/lena.jpg" width="320" height="240"/>

A distance greater than 10 indicates the images are **likely different**.

Remember perceptual hash algorithms can give false positives. 

# Installation

> go get github.com/umahmood/perceptive
    
> cd $GOPATH/src/github.com/umahmood/perceptive
    
> go test ./...

# Dependencies

- [github.com/umahmood/iter](https://www.github.com/umahmood/iter)
- [github.com/disintegration/imaging](https://www.github.com/disintegration/imaging)

# Usage
  
    package main

    import (
        "log"

        "github.com/disintegration/imaging"
        "github.com/umahmood/perceptive"
    )

    func openImage(filePath string) image.Image {
        img, err := imaging.Open(filePath)
        if err != nil {
            log.Fatalln(err)
        }
        return img
    }

    func main() {
        imgA := openImage("lena.jpg")
        imgB := openImage("lena.jpg")

        distance, err := perceptive.CompareImages(imgA, imgB, perceptive.Difference)

        if distance == 0 {
            // images are likely the same
        } else if distance >= 1 && distance <= 10 {
            // images are potentially a variation
        } else {
            // images are likely different
        }
    }

# To do

- Implement Phash algorithm targeting images.
- Compute perceptual hashes for audio files.
- Compute perceptual hashes for video files.

# Documentation

> http://godoc.org/github.com/umahmood/perceptive

# References

- http://phash.org/
- http://www.hackerfactor.com/blog/index.php?/archives/529-Kind-of-Like-That.html

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
