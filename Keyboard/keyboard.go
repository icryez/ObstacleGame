package keyboard

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"sync"
)

type devices []string

func (d *devices) hasDevice(str string) bool {
	for _, device := range *d {
		if strings.Contains(str, device) {
			return true
		}
	}

	return false
}

var restrictedDevices = devices{"mouse", "mice", "Glorious", "glorious"}
var allowedDevices = devices{"keyboard", "Firefly", "logitech mx keys", "Keyboard", "firefly"}

type KeyBoardState struct {
	sync.Mutex
	Keystates map[string]bool
}

var KeysState KeyBoardState

var eventPaths []string

func CreateKeyBoardState() *KeyBoardState {
	kbs := new(KeyBoardState)
	kbs.Keystates = map[string]bool{}
	eventPaths = FindAllKeyboardDevices()
	return kbs
}
func FindAllKeyboardDevices() []string {
	path := "/sys/class/input/event%d/device/name"
	resolved := "/dev/input/event%d"

	valid := make([]string, 0)

	for i := 0; i < 255; i++ {
		buff, err := os.ReadFile(fmt.Sprintf(path, i))

		// prevent from checking non-existant files
		if os.IsNotExist(err) {
			break
		}
		if err != nil {
			continue
		}

		deviceName := strings.ToLower(string(buff))

		if restrictedDevices.hasDevice(deviceName) {
			continue
		} else if allowedDevices.hasDevice(deviceName) {
			valid = append(valid, fmt.Sprintf(resolved, i))
		}
	}
	return valid
}

func StartWatcher() {
	// KeysState.Keystates = make(map[string]bool)
	// KeysState = *createKeyBoardState()

	for _, v := range eventPaths {
		go func(path string) {

			f, err := os.Open(path)
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
				//TODO:use switch case??
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
				} else if code == 46 && value == 1 {
					KeysState.updateKey("C",true)
				}
			}
		}(v)
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
