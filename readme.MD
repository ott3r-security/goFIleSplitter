# File Splitter Utility

This utility is designed to split a large file into smaller, manageable chunks of a specified size. It's particularly useful for processing large datasets or files that are too large to be opened by standard applications.

## Prerequisites

Before you use this utility, ensure you have the following installed:

- Go (Golang) environment set up on your system. The utility is written in Go and requires the Go runtime to execute. Visit [The Go Programming Language](https://golang.org/doc/install) for installation instructions.

## Installation

Clone this repository to your local machine using:

```
git clone https://github.com/ott3r-security/goSpitter.git
cd file-splitter-utility
```

## Usage

To use the utility, navigate to the project directory in your terminal and run:

```
go run main.go <path-to-file> <chunk-size-in-MB>
```

- `<path-to-file>`: The path to the large file you want to split.
- `<chunk-size-in-MB>`: The size of each chunk in megabytes. For example, to split the file into chunks of 256 MB, you would use `256`.

Example:

```
go run main.go /path/to/large/file.txt 256
```

This command splits `file.txt` into 256 MB chunks and stores them in a newly created directory.

## Troubleshooting

- **Error: "Cannot create a file when that file already exists."** - Ensure the target directory for the file chunks does not already exist or has been properly cleaned up before running the utility.

- **Error: "The system cannot find the path specified."** - Verify that the path to the source file is correct and accessible. Also, check that the utility has permission to read the file and write to the target directory.

For more help or to report issues, please visit submit issue
