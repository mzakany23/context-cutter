package cutter

import (
	"os"
	"path/filepath"
	"testing"
)

func createTestFile(t *testing.T, content string) (string, func()) {
	tmpfile, err := os.CreateTemp("", "test-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	return tmpfile.Name(), func() {
		os.Remove(tmpfile.Name())
	}
}

func TestSplitFile(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		chunkSize   int64
		expectError bool
	}{
		{
			name:        "Basic split",
			content:     "Hello, World!",
			chunkSize:   5,
			expectError: false,
		},
		{
			name:        "Empty file",
			content:     "",
			chunkSize:   5,
			expectError: false,
		},
		{
			name:        "Large chunk size",
			content:     "Small content",
			chunkSize:   100,
			expectError: false,
		},
		{
			name:        "Zero chunk size",
			content:     "Content",
			chunkSize:   0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test input file
			inputFile, cleanup := createTestFile(t, tt.content)
			defer cleanup()

			// Create temp output directory
			outputDir, err := os.MkdirTemp("", "test-output-*")
			if err != nil {
				t.Fatalf("Failed to create temp output dir: %v", err)
			}
			defer os.RemoveAll(outputDir)

			// Run the split
			err = SplitFile(inputFile, outputDir, tt.chunkSize)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("SplitFile() error = %v, expectError = %v", err, tt.expectError)
				return
			}

			if !tt.expectError {
				// Verify chunks were created
				files, err := os.ReadDir(outputDir)
				if err != nil {
					t.Fatalf("Failed to read output dir: %v", err)
				}

				if len(files) == 0 && len(tt.content) > 0 {
					t.Error("No chunks were created")
				}

				// Verify content
				var totalContent []byte
				for _, file := range files {
					content, err := os.ReadFile(filepath.Join(outputDir, file.Name()))
					if err != nil {
						t.Fatalf("Failed to read chunk file: %v", err)
					}
					totalContent = append(totalContent, content...)
				}

				if string(totalContent) != tt.content {
					t.Errorf("Reconstructed content = %q, want %q", string(totalContent), tt.content)
				}
			}
		})
	}
}

func TestSplitFileByCount(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		fileCount   int
		expectError bool
	}{
		{
			name:        "Split into 3 files",
			content:     "Hello, World! This is a test content.",
			fileCount:   3,
			expectError: false,
		},
		{
			name:        "Empty file",
			content:     "",
			fileCount:   2,
			expectError: true, // Can't split empty file into chunks
		},
		{
			name:        "Single chunk",
			content:     "Small content",
			fileCount:   1,
			expectError: false,
		},
		{
			name:        "Zero file count",
			content:     "Content",
			fileCount:   0,
			expectError: true,
		},
		{
			name:        "More chunks than content",
			content:     "Hi",
			fileCount:   5,
			expectError: true, // It's better to error in this case
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test input file
			inputFile, cleanup := createTestFile(t, tt.content)
			defer cleanup()

			// Create temp output directory
			outputDir, err := os.MkdirTemp("", "test-output-*")
			if err != nil {
				t.Fatalf("Failed to create temp output dir: %v", err)
			}
			defer os.RemoveAll(outputDir)

			// Run the split
			err = SplitFileByCount(inputFile, outputDir, tt.fileCount)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("SplitFileByCount() error = %v, expectError = %v", err, tt.expectError)
				return
			}

			if !tt.expectError {
				// Verify number of chunks
				files, err := os.ReadDir(outputDir)
				if err != nil {
					t.Fatalf("Failed to read output dir: %v", err)
				}

				if len(tt.content) > 0 && len(files) != tt.fileCount {
					t.Errorf("Got %d chunks, want %d chunks", len(files), tt.fileCount)
				}

				// Verify content
				var totalContent []byte
				for _, file := range files {
					content, err := os.ReadFile(filepath.Join(outputDir, file.Name()))
					if err != nil {
						t.Fatalf("Failed to read chunk file: %v", err)
					}
					totalContent = append(totalContent, content...)
				}

				if string(totalContent) != tt.content {
					t.Errorf("Reconstructed content = %q, want %q", string(totalContent), tt.content)
				}
			}
		})
	}
}

func TestInvalidInputFile(t *testing.T) {
	tests := []struct {
		name      string
		inputFile string
		testFunc  func(string, string) error
	}{
		{
			name:      "SplitFile with non-existent file",
			inputFile: "nonexistent.txt",
			testFunc: func(in, out string) error {
				return SplitFile(in, out, 5)
			},
		},
		{
			name:      "SplitFileByCount with non-existent file",
			inputFile: "nonexistent.txt",
			testFunc: func(in, out string) error {
				return SplitFileByCount(in, out, 3)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputDir, err := os.MkdirTemp("", "test-output-*")
			if err != nil {
				t.Fatalf("Failed to create temp output dir: %v", err)
			}
			defer os.RemoveAll(outputDir)

			err = tt.testFunc(tt.inputFile, outputDir)
			if err == nil {
				t.Error("Expected error for non-existent file, got nil")
			}
		})
	}
}
