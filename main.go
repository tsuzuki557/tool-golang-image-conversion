package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
)

func compressImage(inputPath, outputPath string, quality int) error {
	// 画像ファイルを開く
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 画像をデコードする
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// 出力先のファイルを作成する
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// 画像をJPEG形式でエンコードして、品質を設定して保存する
	opts := &jpeg.Options{Quality: quality}
	err = jpeg.Encode(outFile, img, opts)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// 実行ファイルの絶対パスを取得
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 実行ファイルのディレクトリを取得
	dir := filepath.Dir(exePath)

	// 画像へのパスを組み立てる
	inputPath := filepath.Join(dir, "img", "input.jpg")
	outputPath := filepath.Join(dir, "img", "output.jpg")

	quality := 75 // 0から100までの品質。75は一般的な圧縮率です。

	err = compressImage(inputPath, outputPath, quality)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Compression successful!")
	}
}
