# README

## Authors

- [agoncalv](https://zone01normandie.org/git/agoncalv)
- [ncourtel](https://zone01normandie.org/git/ncourtel)
- [ngenty](https://zone01normandie.org/git/ngenty)

## Description

Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII.

## Usage

```sh
go run . "Hello World!"
```

## Tests

Run:

```sh
bash ./tests.sh
```

You can refer to the ([audit page](https://github.com/01-edu/public/blob/master/subjects/ascii-art/audit/README.md)) in the public git of Zone01 to verify of the results are corrects.  


***

<!-- draft

### ascii

[link](https://github.com/01-edu/public/tree/master/subjects/ascii-art)

Usage: go run . [STRING]

Ex:

```sh
go run . "Hello\n"
```

### color (ascii-art-color)

[link](https://github.com/01-edu/public/blob/master/subjects/ascii-art/color/README.md)

Usage: go run . [OPTION] [STRING]

Ex:

```sh
go run . --color=<color> <substring to be colored> "something"
```

### fs (ascii-art-fs)

[link](https://github.com/01-edu/public/blob/master/subjects/ascii-art/fs/README.md)

Usage: go run . [STRING] [BANNER]

EX:

```sh
go run . something standard
```

### justify (ascii-art-justify)

[link](https://github.com/01-edu/public/blob/master/subjects/ascii-art/justify/README.md)

Usage: go run . [OPTION] [STRING] [BANNER]

Example:

```sh
go run . --align=right something standard
```

### output (ascii-art-output)

[link](https://github.com/01-edu/public/blob/master/subjects/ascii-art/output/README.md)

Usage: go run . [OPTION] [STRING] [BANNER]

EX:

```sh
go run . --output=<fileName.txt> something standard
```

### reverse (ascii-reverse)

[link](https://github.com/01-edu/public/blob/master/subjects/ascii-art/reverse/README.md)

Usage: go run . [OPTION]

EX:

```sh
go run . --reverse=<fileName>
```
-->