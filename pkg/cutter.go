package cutter

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func SplitFile(inputFile, outputDir string, chunkSize int64) error {
	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	// Create output directory if it doesn't exist
	if outputDir == "" {
		hash := md5.Sum([]byte(time.Now().String()))
		outputDir = fmt.Sprintf("cutter-%x", hash[:8])
	}
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create a buffer to read chunks
	buffer := make([]byte, chunkSize)

	// Split the file into chunks
	chunkNum := 0
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("error reading input file: %w", err)
		}
		if bytesRead == 0 {
			break
		}

		// Create a new chunk file
		chunkFileName := filepath.Join(outputDir, fmt.Sprintf("chunk_%04d.txt", chunkNum))
		chunkFile, err := os.Create(chunkFileName)
		if err != nil {
			return fmt.Errorf("failed to create chunk file: %w", err)
		}

		// Write the chunk
		_, err = chunkFile.Write(buffer[:bytesRead])
		chunkFile.Close()
		if err != nil {
			return fmt.Errorf("failed to write chunk: %w", err)
		}

		chunkNum++
	}

	return nil
}

func SplitFileByCount(inputFile, outputDir string, fileCount int) error {
	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}
	fileSize := fileInfo.Size()

	// Calculate chunk size
	chunkSize := fileSize / int64(fileCount)
	lastChunkSize := fileSize - (chunkSize * int64(fileCount-1))

	// Create output directory if it doesn't exist
	if outputDir == "" {
		hash := md5.Sum([]byte(time.Now().String()))
		outputDir = fmt.Sprintf("cutter-%x", hash[:8])
	}
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Split the file into chunks
	for i := 0; i < fileCount; i++ {
		// Determine the size of the current chunk
		currentChunkSize := chunkSize
		if i == fileCount-1 {
			currentChunkSize = lastChunkSize
		}

		// Create a buffer to read the chunk
		buffer := make([]byte, currentChunkSize)

		// Read the chunk
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("error reading input file: %w", err)
		}
		if bytesRead == 0 {
			break
		}

		// Create a new chunk file
		chunkFileName := filepath.Join(outputDir, fmt.Sprintf("chunk_%04d.txt", i))
		chunkFile, err := os.Create(chunkFileName)
		if err != nil {
			return fmt.Errorf("failed to create chunk file: %w", err)
		}

		// Write the chunk
		_, err = chunkFile.Write(buffer[:bytesRead])
		chunkFile.Close()
		if err != nil {
			return fmt.Errorf("failed to write chunk: %w", err)
		}
	}

	return nil
}
