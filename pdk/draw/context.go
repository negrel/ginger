package draw

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/styles/value"
)

// Context define a drawing context on a Canvas.
type Context interface {
	// Canvas returns the canvas that is tied to this Context.
	Canvas() Canvas

	// Bounds returns the bounds of the underlying Canvas of this Context.
	Bounds() geometry.Rectangle

	// SetFillColor sets the fill color for fill methods.
	SetFillColor(color value.Color)

	// FillColor returns the current fill color.
	FillColor() value.Color

	// FillRectangle draws a geometry.Rectangle that is filled according to the current fill color.
	// Note that this methods overwrite text present on the Canvas.
	FillRectangle(rectangle geometry.Rectangle)

	// FillTextH draws the given text horizontally from the given origin on this Canvas.
	// If the text overflow the Canvas of the context, the overflowing chars are dropped.
	FillTextH(origin geometry.Point, text string)

	// FillTextV draws the given text vertically from the given origin on this Canvas.
	// If the text overflow this canvas, the overflowing chars are dropped.
	FillTextV(origin geometry.Point, text string)

	// FillLine draws a line between the two given geometry.Point.
	FillLine(from, to geometry.Point)

	// Commit applies all the change of the context to the Canvas.
	Commit()
}

var _ Context = &context{}

type context struct {
	canvas    Canvas
	bounds    geometry.Rectangle
	fillColor value.Color
	ops       []func(Canvas)
}

// NewContext returns a new Context with the given canvas and bounds.
func NewContext(canvas Canvas, bounds geometry.Rectangle) Context {
	return newContext(canvas, bounds)
}

func newContext(canvas Canvas, bounds geometry.Rectangle) *context {
	return &context{
		canvas: canvas,
		bounds: bounds,
		ops:    make([]func(Canvas), 0, 8),
	}
}

// Canvas implements the Context interface.
func (c *context) Canvas() Canvas {
	return c.canvas
}

// Bounds implements the Context interface.
func (c *context) Bounds() geometry.Rectangle {
	return c.bounds
}

// SetFillColor implements the Context interface.
func (c *context) SetFillColor(color value.Color) {
	c.fillColor = color
}

// FillColor implements the Context interface.
func (c *context) FillColor() value.Color {
	return c.fillColor
}

// FillRectangle implements the Context interface.
func (c *context) FillRectangle(rectangle geometry.Rectangle) {
	rectangle = c.bounds.Mask(rectangle)
	fillColor := c.fillColor

	c.ops = append(c.ops, func(canvas Canvas) {
		for i := rectangle.Min.X(); i < rectangle.Max.X(); i++ {
			for j := rectangle.Min.Y(); j < rectangle.Max.Y(); j++ {
				pos := geometry.Pt(i, j)
				cell := canvas.Get(pos)

				cell.Style = CellStyle{}
				cell.Style.Background = fillColor
				cell.Style.Foreground = fillColor
				cell.Content = 0

				canvas.Set(pos, cell)
			}
		}
	})
}

// FillTextH implements the Context interface.
func (c *context) FillTextH(origin geometry.Point, text string) {
	c.FillRunesH(origin, []rune(text)...)
}

func (c *context) FillRunesH(origin geometry.Point, runes ...rune) {
	origin = origin.Add(c.bounds.Min)
	rectangle := geometry.Rect(origin.X(), origin.Y(), origin.X()+len(runes), origin.Y()+1)
	rectangle = c.bounds.Mask(rectangle)
	if rectangle.Empty() {
		return
	}

	fillColor := c.fillColor
	c.ops = append(c.ops, func(canvas Canvas) {
		for i := rectangle.Min.X(); i < rectangle.Max.X(); i++ {
			pos := geometry.Pt(i, origin.Y())
			cell := canvas.Get(pos)

			cell.Content = runes[i-rectangle.Min.X()]
			cell.Style.Foreground = fillColor

			canvas.Set(pos, cell)
		}
	})
}

// FillTextV implements the Context interface.
func (c *context) FillTextV(origin geometry.Point, text string) {
	c.FillRunesV(origin, []rune(text)...)
}

func (c *context) FillRunesV(origin geometry.Point, runes ...rune) {
	origin = origin.Add(c.bounds.Min)
	rectangle := geometry.Rect(origin.X(), origin.Y(), origin.X()+1, origin.Y()+len(runes))
	rectangle = c.bounds.Mask(rectangle)
	if rectangle.Empty() {
		return
	}

	fillColor := c.fillColor
	c.ops = append(c.ops, func(canvas Canvas) {
		for i := rectangle.Min.Y(); i < rectangle.Max.Y(); i++ {
			pos := geometry.Pt(origin.X(), i)
			cell := canvas.Get(pos)

			cell.Content = runes[i-rectangle.Min.Y()]
			cell.Style.Foreground = fillColor

			canvas.Set(pos, cell)
		}
	})
}

// FillLine implements the Context interface.
func (c *context) FillLine(from, to geometry.Point) {
	panic("implement me")
}

// Commit implements the Context interface.
func (c *context) Commit() {
	for _, op := range c.ops {
		op(c.canvas)
	}
	c.ops = make([]func(Canvas), 0, 8)
}
