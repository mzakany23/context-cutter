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

1. Split a file into 1MB chunks (default):
   ```
   ./context-cutter --input large_file.txt
   ```

2. Split a file into 5MB chunks:
   ```
   ./context-cutter --input large_file.txt --size 5242880
   ```

3. Split a file and specify an output directory:
   ```
   ./context-cutter --input large_file.txt --output ./chunks
   ```

4. Combine all options:
   ```
   ./context-cutter --input large_file.txt --output ./chunks --size 10485760
   ```

After running the command, Context Cutter will split the input file into chunks and save them in the specified (or auto-generated) output directory. Each chunk will be named `chunk_XXXX.txt`, where `XXXX` is a zero-padded sequential number.

### Note:

If you don't specify an output directory, Context Cutter will create a directory with a name in the format `cutter-<hash>`, where `<hash>` is a unique identifier based on the current timestamp.
