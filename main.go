package main

import (
	"flag"
	"fmt"
	"os"
	"net/http"
	"io"
	"log"
	"github.com/fogleman/gg"
)

func fetchImage(link string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("URL Could not be accessed!")
	}
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Could not read the body!")
	}

	file, err := os.Create("image.jpg")
	if err != nil {
		fmt.Println("Could not copy the image!")
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Could not copy the image!")
	}
}

func editImage(text string) {
    im, err := gg.LoadImage("image.jpg")
    if err != nil {
        log.Fatal(err)
    }
	dc := gg.NewContextForImage(im)

	w := float64(dc.Width())
	h := float64(dc.Height())
	boxHeight := 60.0 
	boxY := h/2 - boxHeight/2

	dc.SetRGBA(0, 0, 0, 0.5)
	dc.DrawRectangle(0, boxY, w, boxHeight)
	dc.Fill()

	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace("./Fonts/Arial.ttf", 24); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored(text, w/2, h/2, 0.5, 0.5)
	dc.SavePNG("out.png")
	err = os.Remove("image.jpg")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
}

func main() {
	meme := flag.NewFlagSet("meme", flag.ExitOnError)
	link := meme.String("link", "", "link")
	text := meme.String("text", "", "text")
	help := meme.Bool("help", false, "Show help message")

	if len(os.Args) < 2 {
		fmt.Println("expected meme command")
	}

	switch os.Args[1] {
	case "meme":
		meme.Parse(os.Args[2:])

		if *help {
			fmt.Println("Usage: meme --link=<image_url> --text=<meme_text>")
			fmt.Println("Example: meme --link=https://example.com/image.jpg --text=Hello World")
			return
		}

		fetchImage(*link)
		editImage(*text)

		fmt.Println(*text)

		fmt.Println("\033[32mMeme created! Look at out.png\033[0m")
	}
}
