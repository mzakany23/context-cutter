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
