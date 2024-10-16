package controller

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/resample"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Sound struct {
	Buffer *beep.Buffer
	Format beep.Format
}

var (
	loopedClipsMap     map[string]*Sound
	loopedClipsMutex   sync.Mutex
	loopedStreamsMap   map[string]*beep.Ctrl
	loopedStreamsMutex sync.Mutex
	soundSemaphore     chan struct{}
	mixer              *beep.Mixer
)

func init() {
	loopedClipsMap = make(map[string]*Sound)
	loopedStreamsMap = make(map[string]*beep.Ctrl)
	soundSemaphore = make(chan struct{}, 5)
	sampleRate := beep.SampleRate(44100)
	speaker.Init(sampleRate, sampleRate.N(time.Second/10))
	mixer = &beep.Mixer{}
	speaker.Play(mixer)

	rootDirectory := "src/main/resources/sounds"

	err := filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}
		if !info.IsDir() && loopedCondition(path) {
			err := loadLoopedSound(path)
			if err != nil {
				fmt.Printf("Error loading sound %s: %v\n", path, err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", rootDirectory, err)
	}
}

func loopedCondition(str string) bool {
	str = strings.ToLower(str)
	return strings.HasSuffix(str, "_loop.wav")
}

func loadLoopedSound(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	streamer, format, err := wav.Decode(f)
	if err != nil {
		return err
	}
	defer streamer.Close()

	if format.SampleRate != speaker.SampleRate {
		streamer = resample.Resample(4, format.SampleRate, speaker.SampleRate, streamer)
		format.SampleRate = speaker.SampleRate
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)

	filename := filepath.Base(path)
	loopedClipsMutex.Lock()
	loopedClipsMap[filename] = &Sound{Buffer: buffer, Format: format}
	loopedClipsMutex.Unlock()

	return nil
}

func playSound(strPath string) {
	if loopedCondition(strPath) {
		loopedClipsMutex.Lock()
		sound, ok := loopedClipsMap[strPath]
		loopedClipsMutex.Unlock()
		if ok {
			loopedStreamsMutex.Lock()
			_, exists := loopedStreamsMap[strPath]
			if exists {
				loopedStreamsMutex.Unlock()
				return
			}
			streamer := sound.Buffer.Streamer(0, sound.Buffer.Len())
			loop := beep.Loop(-1, streamer)
			ctrl := &beep.Ctrl{Streamer: loop, Paused: false}
			loopedStreamsMap[strPath] = ctrl
			loopedStreamsMutex.Unlock()

			mixer.Add(ctrl)
		} else {
			fmt.Printf("Looped sound %s not found\n", strPath)
		}
	} else {
		go func() {
			soundSemaphore <- struct{}{}
			defer func() { <-soundSemaphore }()

			err := playNonLoopedSound(strPath)
			if err != nil {
				fmt.Printf("Error playing sound %s: %v\n", strPath, err)
			}
		}()
	}
}

func stopSound(strPath string) {
	if !loopedCondition(strPath) {
		return
	}
	loopedStreamsMutex.Lock()
	ctrl, ok := loopedStreamsMap[strPath]
	if ok {
		ctrl.Paused = true
		ctrl.Streamer = nil
		delete(loopedStreamsMap, strPath)
	}
	loopedStreamsMutex.Unlock()
}

func playNonLoopedSound(strPath string) error {
	f, err := os.Open("src/main/resources/sounds/" + strPath)
	if err != nil {
		return err
	}
	defer f.Close()

	streamer, format, err := wav.Decode(f)
	if err != nil {
		return err
	}
	defer streamer.Close()

	if format.SampleRate != speaker.SampleRate {
		streamer = resample.Resample(4, format.SampleRate, speaker.SampleRate, streamer)
		format.SampleRate = speaker.SampleRate
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done

	return nil
}

func main() {
	playSound("background_loop.wav")
	time.Sleep(5 * time.Second)
	stopSound("background_loop.wav")

	playSound("effect.wav")
	playSound("effect.wav")
	playSound("effect.wav")
	playSound("effect.wav")
	playSound("effect.wav")
	playSound("effect.wav") // This will wait if there are already 5 sounds playing
	time.Sleep(5 * time.Second)
}
