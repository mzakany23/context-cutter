package main

import (
	cutter "context-cutter/pkg"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var inputFile string
	var outputDir string
	var chunkSize int64
	var fileCount int

	rootCmd := &cobra.Command{
		Use:   "file-cutter",
		Short: "Split large files into smaller chunks",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			if fileCount > 0 {
				err = cutter.SplitFileByCount(inputFile, outputDir, fileCount)
			} else {
				err = cutter.SplitFile(inputFile, outputDir, chunkSize)
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("File successfully split into chunks.")
		},
	}

	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file to split (required)")
	rootCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Output directory for chunks (optional)")
	rootCmd.Flags().Int64VarP(&chunkSize, "size", "s", 1024*1024, "Chunk size in bytes (default 1MB)")
	rootCmd.Flags().IntVarP(&fileCount, "files", "f", 0, "Number of files to split into (optional)")

	rootCmd.MarkFlagRequired("input")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
