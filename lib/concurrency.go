package lib

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
	defer close(fg.errCh)
	fg.dispatch()

	for fg.n < fg.counter {
		err := <-fg.errCh
		if err != nil {
			return err
		}
		fg.n += 1
	}

	return nil
}
