package keyboard

import (
	"bytes"
	"encoding/binary"
	"os"
	"sync"
)

type KeyBoardState struct {
	sync.Mutex
	Keystates map[string]bool
}

var KeysState KeyBoardState

func createKeyBoardState() *KeyBoardState {
	kbs := new(KeyBoardState)
	kbs.Keystates = map[string]bool{}
	return kbs
}

func StartWatcher() {
	// KeysState.Keystates = make(map[string]bool)
	KeysState = *createKeyBoardState()

	f, err := os.Open("/dev/input/event16") //TODO: make this dynamic
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := make([]byte, 24)
	for {
		f.Read(b)
		var value int32
		// typ := binary.LittleEndian.Uint16(b[16:18])
		code := binary.LittleEndian.Uint16(b[18:20])
		binary.Read(bytes.NewReader(b[20:]), binary.LittleEndian, &value)
		if code == 57 && value == 1 {
			KeysState.updateKey("space", true)
		} else if code == 57 && value == 0 {
			KeysState.updateKey("space", false)
		} else if code == 30 && value == 1 {
			KeysState.updateKey("A", true)
		} else if code == 30 && value == 0 {
			KeysState.updateKey("A", false)
		} else if code == 32 && value == 1 {
			KeysState.updateKey("D", true)
		} else if code == 32 && value == 0 {
			KeysState.updateKey("D", false)
		} else if code == 1 && value == 1 {
			KeysState.updateKey("Esc", true)
		}
	}
}

func (k *KeyBoardState) updateKey(key string, flag bool) {
	k.Lock()
	defer k.Unlock()
	k.Keystates[key] = flag
}

func (k *KeyBoardState) GetKey(key string) bool {
	k.Lock()
	defer k.Unlock()
	return k.Keystates[key]
}
