# Wordcounter

A simple wordcounter written in go.  
Accepts either filepaths in args or piped input of filepaths ex: output from <-find->

**Outputs total line, word, and bytecount**

## Usage

```bash
go install github.com/be-nice/wordcounter

wordcount file1.txt ./relative/file2

or

find . -name "*.go" | wordcount
```
