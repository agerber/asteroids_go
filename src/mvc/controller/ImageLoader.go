package util

import (
	"image"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

var IMAGE_MAP map[string]*ebiten.Image

// LoadPngImages recursively loads all PNG images from the specified root directory and its subdirectories.
func LoadPngImages(rootDirectory string) (map[string]*ebiten.Image, error) {
	images := make(map[string]*ebiten.Image)

	err := filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(strings.ToLower(info.Name()), ".png") &&
			!strings.Contains(strings.ToLower(info.Name()), "do_not_load.png") {

			imgFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer imgFile.Close()

			img, _, err := image.Decode(imgFile)
			if err != nil {
				return err
			}

			ebitenImg := ebiten.NewImageFromImage(img)

			// Get the relative path, ignoring case, and store the image in IMAGE_MAP
			relPath := strings.ToLower(strings.TrimPrefix(path, rootDirectory+"/"))
			images[relPath] = ebitenImg

			log.Printf("Loaded image: %s", relPath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return images, nil
}

// GetImage retrieves the image by its file name, searching within the loaded IMAGE_MAP.
func GetImage(fileName string) *ebiten.Image {
	fileName = strings.ToLower(fileName)
	for path, img := range IMAGE_MAP {
		if strings.HasSuffix(path, fileName) {
			log.Printf("Found image for %s at path: %s", fileName, path)
			return img
		}
	}
	log.Printf("Image not found: %s", fileName)
	return nil
}

func getSourceDirectory() string {
	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get source directory")
	}
	return filepath.Dir(filePath) 
}

func init() {
	var err error

	sourceDirectory := getSourceDirectory()

	imageDirectory := filepath.Join(sourceDirectory, "..", "..", "resources", "imgs")

	log.Print("Loading images from directory: ", imageDirectory)

	// Load all images into the global IMAGE_MAP
	IMAGE_MAP, err = LoadPngImages(imageDirectory)
	if err != nil {
		log.Fatalf("Failed to load images: %v", err)
	}
}
