package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
)

func compressImage(inputPath, outputPath string, quality int) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	opts := &jpeg.Options{Quality: quality}
	err = jpeg.Encode(outFile, img, opts)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	dir := filepath.Dir(exePath)
	imgDir := filepath.Join(dir, "img")

	quality := 75

	// 指定ディレクトリ内のすべてのファイルとディレクトリをリストアップ
	files, err := os.ReadDir(imgDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, f := range files {
		filename := f.Name()

		// .jpgファイルのみを対象とする
		if strings.HasSuffix(strings.ToLower(filename), ".jpg") {
			inputPath := filepath.Join(imgDir, filename)
			// 例えば、input.jpgの場合、出力はoutput-input.jpgとする
			outputPath := filepath.Join(imgDir, "output-"+filename)

			err := compressImage(inputPath, outputPath, quality)
			if err != nil {
				fmt.Printf("Error compressing %s: %v\n", filename, err)
			} else {
				fmt.Printf("Compression of %s successful!\n", filename)
			}
		}
	}
}
