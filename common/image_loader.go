package common

import (
	"bufio"
	"image"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

var imageMap map[string]*ebiten.Image

func GetImage(imgPath string) *ebiten.Image {
	return imageMap[imgPath]
}

func init() {
	imageMap = make(map[string]*ebiten.Image)
	loadPNGImages("assets/imgs")
}

func loadPNGImages(rootDir string) {
	files, err := os.ReadDir(rootDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		path := filepath.Join(rootDir, file.Name())
		if file.IsDir() {
			loadPNGImages(path)
			continue
		}
		lowerCaseFileName := strings.ToLower(file.Name())
		if strings.HasSuffix(lowerCaseFileName, ".png") && !strings.Contains(lowerCaseFileName, "do_not_load.png") {
			image, err := loadImage(path)
			if err != nil {
				panic(err)
			}
			imageMap[path] = image
		}
	}
}

func loadImage(imgPath string) (*ebiten.Image, error) {
	imgFile, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bufio.NewReader(imgFile))
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}
