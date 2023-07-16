package mywidgets

import (

	// "image/color"
	// "log"
	// "math"
	// "math/rand"
	// "time"

	// "github.com/ebitenui/ebitenui"
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	// "github.com/golang/freetype/truetype"
	// "github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "github.com/hajimehoshi/ebiten/v2/inpututil"
	// "github.com/hajimehoshi/ebiten/v2/vector"
	// "golang.org/x/image/font"
	// "golang.org/x/image/font/gofont/goregular"
)

func GetContainer() *widget.Container {

	// load images for button states: idle, hover, and pressed
	// buttonImage, _ := loadButtonImage()

	// load button text font
	// face, _ := loadFont(20)

	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),

		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	// // construct a button
	// buttonReset := widget.NewButton(
	// 	// set general widget options
	// 	widget.ButtonOpts.WidgetOpts(
	// 		// instruct the container's anchor layout to center the button both horizontally and vertically
	// 		widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
	// 			HorizontalPosition: widget.AnchorLayoutPositionCenter,
	// 			VerticalPosition:   widget.AnchorLayoutPositionCenter,
	// 		}),
	// 	),

	// 	// specify the images to use
	// 	widget.ButtonOpts.Image(buttonImage),

	// 	// specify the button's text, the font face, and the color
	// 	widget.ButtonOpts.Text("Reset", face, &widget.ButtonTextColor{
	// 		Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
	// 	}),

	// 	// specify that the button's text needs some padding for correct display
	// 	widget.ButtonOpts.TextPadding(widget.Insets{
	// 		Left:   30,
	// 		Right:  30,
	// 		Top:    5,
	// 		Bottom: 5,
	// 	}),

	// 	// add a handler that reacts to clicking the button
	// 	widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
	// 		println("button clicked")
	// 	}),
	// )

	// rootContainer.AddChild(buttonReset)

	// buttonNextFrame := widget.NewButton(
	// 	// set general widget options
	// 	widget.ButtonOpts.WidgetOpts(
	// 		// instruct the container's anchor layout to center the button both horizontally and vertically
	// 		widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
	// 			HorizontalPosition: widget.AnchorLayoutPositionCenter,
	// 			VerticalPosition:   widget.AnchorLayoutPositionCenter,
	// 		}),
	// 	),

	// 	// specify the images to use
	// 	widget.ButtonOpts.Image(buttonImage),

	// 	// specify the button's text, the font face, and the color
	// 	widget.ButtonOpts.Text("Next Frame", face, &widget.ButtonTextColor{
	// 		Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
	// 	}),

	// 	// specify that the button's text needs some padding for correct display
	// 	widget.ButtonOpts.TextPadding(widget.Insets{
	// 		Left:   30,
	// 		Right:  30,
	// 		Top:    5,
	// 		Bottom: 5,
	// 	}),

	// 	// add a handler that reacts to clicking the button
	// 	widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
	// 		println("button clicked")
	// 	}),
	// )

	// rootContainer.AddChild(buttonNextFrame)

	// // construct a slider
	// slider := widget.NewSlider(
	// 	// Set the slider orientation - n/s vs e/w
	// 	widget.SliderOpts.Direction(widget.DirectionHorizontal),
	// 	// Set the minimum and maximum value for the slider
	// 	widget.SliderOpts.MinMax(0, 10),

	// 	widget.SliderOpts.WidgetOpts(
	// 		// Set the Widget to layout in the center on the screen
	// 		widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
	// 			HorizontalPosition: widget.AnchorLayoutPositionCenter,
	// 			VerticalPosition:   widget.AnchorLayoutPositionCenter,
	// 		}),
	// 		// Set the widget's dimensions
	// 		widget.WidgetOpts.MinSize(200, 6),
	// 	),
	// 	widget.SliderOpts.Images(
	// 		// Set the track images
	// 		&widget.SliderTrackImage{
	// 			Idle:  image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
	// 			Hover: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
	// 		},
	// 		// Set the handle images
	// 		&widget.ButtonImage{
	// 			Idle:    image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
	// 			Hover:   image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
	// 			Pressed: image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
	// 		},
	// 	),
	// 	// Set the size of the handle
	// 	widget.SliderOpts.FixedHandleSize(6),
	// 	// Set the offset to display the track
	// 	widget.SliderOpts.TrackOffset(0),
	// 	// Set the size to move the handle
	// 	widget.SliderOpts.PageSizeFunc(func() int {
	// 		return 1
	// 	}),
	// 	// Set the callback to call when the slider value is changed
	// 	widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
	// 		fmt.Println(args.Current)
	// 	}),
	// )
	// // Set the current value of the slider
	// slider.Current = 5
	// // add the slider as a child of the container
	// rootContainer.AddChild(slider)

	return rootContainer
}

func loadButtonImage() (*widget.ButtonImage, error) {
	idle := image.NewNineSliceColor(color.NRGBA{170, 170, 180, 255})
	hover := image.NewNineSliceColor(color.NRGBA{130, 130, 150, 255})
	pressed := image.NewNineSliceColor(color.NRGBA{100, 100, 120, 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}

func loadFont(size float64) (font.Face, error) {
	ttfFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}
