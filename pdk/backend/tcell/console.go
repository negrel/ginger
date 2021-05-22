package tcell

import (
	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/backend"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/events"
)

var _ backend.Console = &Console{}

// Console is a wrapper around https://www.github.com/gdamore/tcell Screen
// that satisfy the backend.Console interface.
type Console struct {
	tcell.Screen

	eventChannel chan<- events.Event
	done         chan struct{}
}

// NewConsole returns a new Console object configured with the
// given options.
func NewConsole(options ...Option) (*Console, error) {
	console := &Console{}

	var err error
	for _, option := range options {
		err = option(console)
		if err != nil {
			return nil, err
		}
	}

	if console.Screen == nil {
		console.Screen, err = tcell.NewScreen()
		if err != nil {
			return nil, err
		}
	}

	return console, nil
}

// Bounds implements the draw.Canvas interface.
func (c *Console) Bounds() geometry.Rectangle {
	w, h := c.Screen.Size()
	return geometry.Rect(0, 0, w, h)
}

// Get implements the draw.Canvas interface.
func (c *Console) Get(pos geometry.Point) draw.Cell {
	return fromTcell(c.Screen.GetContent(pos.X(), pos.Y()))
}

// Set implements the draw.Canvas interface.
func (c *Console) Set(pos geometry.Point, cell draw.Cell) {
	mainc, combc, style := toTcell(cell)
	c.Screen.SetContent(pos.X(), pos.Y(), mainc, combc, style)
}

func (c *Console) NewContext(bounds geometry.Rectangle) draw.Context {
	return draw.NewContext(c, bounds)
}

// Clear implements the backend.Console interface.
func (c *Console) Clear() {
	c.Screen.Clear()
}

// Flush implements the backend.Console interface.
func (c *Console) Flush() {
	c.Screen.Show()
}

// Start implements the backend.Console interface.
func (c *Console) Start() error {
	err := c.Screen.Init()
	if err != nil {
		return err
	}

	c.done = make(chan struct{})

	if c.eventChannel != nil {
		go eventLoop(c.done, c.eventChannel, c.Screen.PollEvent)
	}

	return nil
}

func (c *Console) Stop() {
	c.Screen.Fini()
	c.done <- struct{}{}
	close(c.done)
}

func adaptEvent(event tcell.Event) events.Event {
	switch ev := event.(type) {
	case *tcell.EventError:
		_ = ev
		return nil

	default:
		return nil
	}
}

func eventLoop(done <-chan struct{}, eventChannel chan<- events.Event, pollEvent func() tcell.Event) {
	ch := make(chan events.Event)

	go func(ch chan<- events.Event) {
		ch <- adaptEvent(pollEvent())
	}(ch)

loop:
	for {
		select {
		case ev := <-ch:
			eventChannel <- ev

		case <-done:
			close(ch)
			break loop
		}
	}
}
