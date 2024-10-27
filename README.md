# Context Cutter

Context Cutter is a command-line tool that splits large files into smaller chunks, making it easier to process or analyze large datasets.

## Table of Contents
- [About](#about)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## About

Context Cutter is designed to help users manage large files by splitting them into smaller, more manageable chunks. This can be particularly useful when dealing with large datasets, logs, or any other sizeable files that need to be processed or analyzed in parts.

## Features

- Split large files into smaller chunks
- Specify custom chunk sizes
- Automatically generate output directories
- Simple command-line interface

## Getting Started

These instructions will help you set up the project on your local machine for development and testing purposes.

### Prerequisites

- Go 1.23.2 or later

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/mzakany23/context-cutter.git
   cd context-cutter
   ```

2. Build the project:
   ```
   go build -o context-cutter cmd/main.go
   ```

## Usage

To use Context Cutter, run the following command:

```
./context-cutter --input <input_file> [--output <output_directory>] [--size <chunk_size>]
```

### Command-line Options:

- `--input` or `-i`: Specify the input file to be split (required)
- `--output` or `-o`: Specify the output directory for chunks (optional, default: auto-generated directory)
- `--size` or `-s`: Specify the chunk size in bytes (optional, default: 1MB)

### Examples:

1. Split a file into 1MB chunks:
   ```
   ./file-cutter -i large_file.txt -s 1048576
   ```

2. Split a file into 10MB chunks and specify an output directory:
   ```
   ./file-cutter -i large_file.txt -s 10485760 -o output_chunks
   ```

3. Split a file into a specific number of files:
   ```
   ./file-cutter -i large_file.txt -f 5
   ```

4. Split a file into 3 files and specify an output directory:
   ```
   ./file-cutter -i large_file.txt -f 3 -o output_chunks
   ```

5. Use short flag names:
   ```
   ./file-cutter -i large_file.txt -s 5242880 -o chunks
   ```

6. Split a file into 1MB chunks (default behavior):
   ```
   ./file-cutter -i large_file.txt
   ```

After running the command, Context Cutter will split the input file into chunks and save them in the specified (or auto-generated) output directory. Each chunk will be named `chunk_XXXX.txt`, where `XXXX` is a zero-padded sequential number.

### Note:

If you don't specify an output directory, Context Cutter will create a directory with a name in the format `cutter-<hash>`, where `<hash>` is a unique identifier based on the current timestamp.
