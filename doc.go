/*
Package perceptive implements perceptual hash algorithms for comparing images.

Perceptual hash algorithms are a family of comparable hash functions which
generate distinct (but not unique) fingerprints, these fingerprints are then
comparable.

Perceptual hash algorithms are mainly used for detecting duplicates of the same
files, in a way that standard and cryptographic hashes generally fail.

The following perceptual hash algorithms are implemented:

- Average Hash (Ahash) - Fast but generates a huge number of false positives.

- Difference Hash (Dhash) - Fast and very few false positives.

Below are some examples on how to use the library:

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

You can also use the perceptual hash algorithms directly, this is good if you
want to store the hashes in a database or some look up table:

    func main() {
        imgA := openImage("lena.jpg")
        imgB := openImage("lena.jpg")

        hash1, err := perceptive.Dhash(imgA)

        if err != nil {
            // handle error
        }

        hash2, err := perceptive.Dhash(imgB)

        if err != nil {
            // handle error
        }

        // hash1 and hash2 can be stored (in a database, etc...) for later use

        // hash1 and hash2 can be compared directly
        distance := perceptive.HammingDistance(hash1, hash2)

        ...
    }

When performing a Hamming distance on two hashes from Ahash or Dhash, the
distance output has the following meaning:

- A distance of 0 means that the images are likely the same.

- A distance between 1-10 indicates the images are likely a variation of each
other.

- A distance greater than 10 indicates the images are likely different.
*/
package perceptive
