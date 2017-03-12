package watcher

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/go-xiaohei/pugo/app/helper/printer"
	"gopkg.in/fsnotify.v1"
)

var (
	// ErrWatchNoFunction means watching function is not found
	ErrWatchNoFunction = errors.New("watching function is missing")
)

// Watcher is a watcher to directories
type Watcher struct {
	locker       sync.Mutex
	w            *fsnotify.Watcher
	scheduleTime int64
	watchExt     []string
	watchDirs    map[string]bool
	done         chan struct{}
}

// New creates new Watcher,
// it does not start
func New() (*Watcher, error) {
	var (
		w = &Watcher{
			scheduleTime: 0,
			watchDirs:    make(map[string]bool),
			watchExt:     []string{".md", ".toml", ".html", ".css", ".js", ".jpg", ".png", ".gif"},
			done:         make(chan struct{}),
		}
		err error
	)
	w.w, err = fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return w, nil
}

// SetExt sets watching files' extension before Start.
func (w *Watcher) SetExt(ext ...string) {
	w.watchExt = ext
}

// Start starts watching and doing the fn if notice
func (w *Watcher) Start(fn func()) error {
	if fn == nil {
		return ErrWatchNoFunction
	}
	go func() {
		c := time.Tick(1 * time.Second)
		for {
			t := <-c
			if w.scheduleTime > 0 && t.UnixNano() > w.scheduleTime {
				w.locker.Lock()
				w.scheduleTime = 0
				w.locker.Unlock()
				// do again
				fn()
			}
		}
	}()
	go func() {
		for {
			select {
			case event := <-w.w.Events:
				printer.Logf(event.String())
				ext := path.Ext(event.Name)
				for _, e := range w.watchExt {
					if e == ext {
						if event.Op != fsnotify.Chmod {
							printer.Info("watch changes %v", event.String())
							w.locker.Lock()
							w.scheduleTime = time.Now().Add(time.Second).UnixNano()
							w.locker.Unlock()
						}
						break
					}
				}
			case err := <-w.w.Errors:
				printer.Error("watch notify error : %v", err)
			}
		}
	}()
	<-w.done
	return nil
}

// Add adds dir to Watcher
func (w *Watcher) Add(dirs ...string) error {
	for _, dir := range dirs {
		if err := w.w.Add(dir); err != nil {
			return err
		}
		printer.Trace("watch dir %v", dir)
		err := filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				if dir == p {
					return nil
				}
				if err = w.w.Add(p); err != nil {
					return err
				}
				printer.Trace("watch dir %v", p)
				w.locker.Lock()
				w.watchDirs[p] = true
				w.locker.Unlock()
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// Close closes Watcher
func (w *Watcher) Close() error {
	w.done <- struct{}{}
	return w.w.Close()
}
