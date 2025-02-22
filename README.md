# Meme Generator

A simple command-line meme generator written in Go. This is my first project learning Go. It fetches an image from a URL, overlays text on it, and saves the result as `out.png`.

## Features
- Fetches images from a URL.
- Adds text to the image with a semi-transparent background.
- Outputs a new meme image.

## Installation

### Prerequisites
Make sure you have Go installed. If not, download and install it from [Go's official website](https://go.dev/dl/).

### Clone the Repository
```sh
git clone https://github.com/AvyuktBallari/MemeGenerator.git
cd meme-generator
```

### Install Dependencies
```sh
go mod init meme-generator
go get github.com/fogleman/gg
```

### Ensure Fonts Exist
Create a `Fonts` directory and add an `Arial.ttf` font file:
```sh
mkdir Fonts
wget -O Fonts/Arial.ttf https://github.com/google/fonts/blob/main/apache/arial/Arial.ttf?raw=true
```

## Usage
### Compile the Program
```sh
go build -o meme
```

### Run the Meme Generator
```sh
./meme meme --link="https://example.com/image.jpg" --text="Hello World"
```

or without compiling:
```sh
go run main.go meme --link="https://example.com/image.jpg" --text="Hello World"
```

## Example Output
- Input: `--link=https://example.com/image.jpg --text="Hello World"`
- Output: `out.png` (saved in the project directory)

## Troubleshooting
- **URL Could not be accessed!**: Ensure the image URL is correct and accessible.
- **Could not copy the image!**: Check your network connection.
- **panic: open ./Fonts/Arial.ttf: no such file or directory**: Ensure the font file exists in the `Fonts/` directory.
- **Command not found**: Ensure the binary is executable (`chmod +x meme` on macOS/Linux).

## License
This project is open-source and available under the MIT License.

## Author
Created by [Avyukt Ballari](https://github.com/AvyuktBallari). Contributions welcome!

