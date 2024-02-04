package lib

import "fmt"

type FgConc struct {
	counter  int
	n        int
	errCh    chan error
	handlers []FgHandler
}

type FgHandler func() error

// New instantiate new fg concurrency
func New() *FgConc {
	return &FgConc{
		counter: 0,
		n:       0,
		errCh:   make(chan error),
	}
}

// Dispatch
func (fgc *FgConc) Go(fn FgHandler) {
	fgc.handlers = append(fgc.handlers, fn)
	fgc.counter += 1
}

func (fgc *FgConc) dispatch() {
	for _, fn := range fgc.handlers {
		go func(fn FgHandler) {
			err := fn()
			fmt.Println("err: ", err)
			if err != nil {
				fgc.errCh <- err
				return
			}
			fgc.errCh <- nil
		}(fn)
	}
}

// WhenDone
func (fg *FgConc) WhenDone() error {
	fg.dispatch()

	fmt.Println("sampe sini")
	for fg.n < fg.counter {
		err := <-fg.errCh
		if err != nil {
			close(fg.errCh)
			return err
		}
		fg.n += 1
	}
	close(fg.errCh)

	return nil
}
