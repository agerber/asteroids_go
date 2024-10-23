package common

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const sampleRate = 44100

var soundMap map[string]*audio.Player

func init() {
	soundMap = make(map[string]*audio.Player)
	loadWAVSounds("assets/sounds")
}

func loadWAVSounds(rootDir string) {
	audioContext := audio.NewContext(sampleRate)

	files, err := os.ReadDir(rootDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		path := filepath.Join(rootDir, file.Name())
		if file.IsDir() {
			loadWAVSounds(path)
			continue
		}
		lowerCaseFileName := strings.ToLower(file.Name())
		if strings.HasSuffix(lowerCaseFileName, "_loop.wav") {
			data, err := os.Open(path)
			if err != nil {
				panic(err)
			}

			stream, err := wav.DecodeWithoutResampling(data)
			if err != nil {
				panic(err)
			}

			loop := audio.NewInfiniteLoop(stream, stream.Length())
			player, err := audioContext.NewPlayer(loop)
			if err != nil {
				panic(err)
			}

			soundMap[file.Name()] = player
		} else if strings.HasSuffix(lowerCaseFileName, ".wav") {
			data, err := os.Open(path)
			if err != nil {
				panic(err)
			}

			stream, err := wav.DecodeWithoutResampling(data)
			if err != nil {
				panic(err)
			}

			player, err := audioContext.NewPlayer(stream)
			if err != nil {
				panic(err)
			}

			soundMap[file.Name()] = player
		}
	}
}

func PlaySound(fileName string) {
	player, ok := soundMap[fileName]
	if !ok {
		log.Printf("Sound %s not found\n", fileName)
		return
	}

	if player.IsPlaying() {
		return
	}

	_ = player.Rewind()
	player.Play()
}

func StopSound(fileName string) {
	player, ok := soundMap[fileName]
	if !ok {
		log.Printf("Sound %s not found\n", fileName)
		return
	}

	player.Pause()
}

func CloseSound() {
	for _, player := range soundMap {
		_ = player.Close()
	}
}
