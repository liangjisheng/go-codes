package watcher

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

// Watcher ...
type Watcher struct {
	Events chan uint32
	fd     int
}

// NewWatcher ...
func NewWatcher() (*Watcher, error) {
	fd, err := unix.InotifyInit()
	if err != nil {
		return nil, err
	}

	watcher := &Watcher{
		Events: make(chan uint32),
		fd:     fd,
	}
	watcher.getEvents()
	return watcher, nil
}

// AddWatcher ...
func (w *Watcher) AddWatcher(file string, mask uint32) error {
	_, err := unix.InotifyAddWatch(w.fd, file, mask)
	if err != nil {
		return err
	}
	return nil
}

func (w *Watcher) getEvents() {
	fmt.Println("unix.SizeofInotifyEvent", unix.SizeofInotifyEvent)
	go func() {
		var buf [unix.SizeofInotifyEvent * 4096]byte
		for {
			n, err := unix.Read(w.fd, buf[:])
			if err != nil {
				n = 0
				continue
			}

			var offset uint32
			fmt.Println("n:", n)
			fmt.Println("offset:", offset)
			fmt.Println(uint32(n - unix.SizeofInotifyEvent))

			for offset <= uint32(n-unix.SizeofInotifyEvent) {
				raw := (*unix.InotifyEvent)(unsafe.Pointer(&buf[offset]))
				mask := uint32(raw.Mask)
				nameLen := uint32(raw.Len)

				w.Events <- mask
				offset += unix.SizeofInotifyEvent + nameLen
			}
		}
	}()
}
