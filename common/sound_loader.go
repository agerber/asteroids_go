package common

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const sampleRate = 44100

var (
	audioContext   = audio.NewContext(sampleRate)
	loopPlayerMap  map[string]*audio.Player
	soundSourceMap map[string][]byte
)

func init() {
	loopPlayerMap = make(map[string]*audio.Player)
	soundSourceMap = make(map[string][]byte)
	loadWAVSounds("assets/sounds")
}

func loadWAVSounds(rootDir string) {
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

			loopPlayerMap[file.Name()] = player
		} else if strings.HasSuffix(lowerCaseFileName, ".wav") {
			data, err := os.Open(path)
			if err != nil {
				panic(err)
			}

			stream, err := wav.DecodeWithoutResampling(data)
			if err != nil {
				panic(err)
			}

			soundSource, err := io.ReadAll(stream)
			if err != nil {
				panic(err)
			}

			soundSourceMap[file.Name()] = soundSource
		}
	}
}

func isLoopSound(fileName string) bool {
	return strings.HasSuffix(strings.ToLower(fileName), "_loop.wav")
}

func PlaySound(fileName string) {
	if isLoopSound(fileName) {
		player, ok := loopPlayerMap[fileName]
		if !ok {
			log.Printf("Loop player %s not found\n", fileName)
			return
		}

		if player.IsPlaying() {
			player.Pause()
		}

		_ = player.Rewind()
		player.Play()
		return
	}

	soundSource, ok := soundSourceMap[fileName]
	if !ok {
		log.Printf("Sound source %s not found\n", fileName)
		return
	}

	player := audioContext.NewPlayerFromBytes(soundSource)
	player.Play()
}

func StopSound(fileName string) {
	if !isLoopSound(fileName) {
		return
	}

	player, ok := loopPlayerMap[fileName]
	if !ok {
		log.Printf("Loop player %s not found\n", fileName)
		return
	}

	if player.IsPlaying() {
		player.Pause()
	}
}

func CloseSound() {
	for _, player := range loopPlayerMap {
		_ = player.Close()
	}
}
