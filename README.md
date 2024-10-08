# Word Count Utility

This project is a custom implementation of the classic Unix `wc` command. It was created as part of a coding challenge [here](https://codingchallenges.fyi/challenges/challenge-wc). The utility is designed to perform byte, word, line, and character counts on text data either from a file or standard input.

## Features

- Count bytes in the input text.
- Count words in the input text.
- Count lines in the input text.
- Count characters in the input text.
- Handle input from both files and standard input (piping).
- Efficiently process large files up to 100MB in size.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- You need to have Go installed on your machine (Go 1.15 or later is recommended).
- You can download and install Go from [https://golang.org/dl/](https://golang.org/dl/).

### Installing

Clone the repository to your local machine:

```bash
git clone https://github.com/nullsploit01/cc-wc.git
cd cc-wc
```

### Building

Compile the project using:

```bash
go build -o ccwc
```

### Usage

To run the utility, you can either pass a file name as an argument or pipe text into it via standard input.

#### Count lines in a file:

```bash
./ccwc -l filename.txt
```

#### Count words in a file:

```bash
./ccwc -w filename.txt
```

#### Count characters in a file:

```bash
./ccwc -m filename.txt
```

#### Count bytes in a file:

```bash
./ccwc -c filename.txt
```

#### Using piped input

```bash
cat filename.txt | ./ccwc -l
```

### Running the Tests

```bash
go test ./...
```
