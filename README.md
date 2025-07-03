# Wordcounter

A simple wordcounter written in go.  
Accepts either filepaths in args or piped input of filepaths ex: output from <-find->

**Outputs total line, word, and bytecount**

## Instructions

### Requierments

**Go 1.16 or newer**

### Installing

```bash
go install github.com/be-nice/wordcounter@latest
```

### Usage

```bash
wordcount file1.txt ./relative/file2

or

find . -name "*.go" | wordcount
```
