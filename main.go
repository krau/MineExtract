package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"
)

type FileInfo struct {
	Hash string `json:"hash"`
	Size int    `json:"size"`
}

type IndexData struct {
	Objects map[string]FileInfo `json:"objects"`
}

func extractFile(hashDir, destPath string) error {
	if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
		return err
	}

	srcFile, err := os.Open(hashDir)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func extractResources(indexPath, objectsDir, outputDir string) error {
	indexFile, err := os.Open(indexPath)
	if err != nil {
		return err
	}
	defer indexFile.Close()

	var indexData IndexData
	if err := json.NewDecoder(indexFile).Decode(&indexData); err != nil {
		return err
	}

	bar := progressbar.Default(int64(len(indexData.Objects)))

	for filePath, fileInfo := range indexData.Objects {
		bar.Add(1)
		hashDir := filepath.Join(objectsDir, fileInfo.Hash[:2], fileInfo.Hash)
		destPath := filepath.Join(outputDir, filePath)

		if err := extractFile(hashDir, destPath); err != nil {
			return fmt.Errorf("failed to extract %s: %w", filePath, err)
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <index.json path> <objects dir> <output dir>")
		return
	}

	indexPath := os.Args[1]
	objectsDir := os.Args[2]
	outputDir := os.Args[3]

	if err := extractResources(indexPath, objectsDir, outputDir); err != nil {
		fmt.Println("Error:", err)
	}
}
