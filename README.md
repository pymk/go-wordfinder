# go-wordfinder

A CLI tool to search for words in text files using regex and configurable matching options.

## Usage

```bash
go run . -file=path/to/file.txt -term=searchterm
go-wordfinder -file=path/to/file.txt -term=searchterm
```

## Options

```bash
-file            Path to the file
-term            Term of interest
-case-sensitive  Match with case-sensitivity (default: true)
-whole-word      Match whole words only (default: true)
```

## Examples

Search for "Marika" in a file, case-sensitive:

```bash
go-wordfinder -file=elden-ring.txt -term=Marika
go-wordfinder -file=elden-ring.txt -term=marika -case-sensitive=false -whole-word=true
```
