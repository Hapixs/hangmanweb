# Hangman classic

Sumary :  
1) The description of your project
2) The available features
3) How to install the project on your computer
4) How to launch the project
5) Examples

## Description

This is a simple hangman game that runs in the command line. It is written in Go.

## Features

- Choose the difficulty of the game
- Choose the file of words
- Choose the ascii art file
- Choose the size of the hangman ascii art
- Choose if you want to auto clear the terminal after each actions
- Choose if you want to auto save the game after each actions
- Choose if you want to use ascii art for letters

## Installation

To install the project, you need to have Go 1.19 installed on your computer. Then, you can run the following command to install the library and lauch the interface version of the game:
```bash
go get github.com/rivo/tview
```

## Launch

To run the game, simply go to the main directory:

```bash
cd main
```
Then run the program with the following command:
```bash
go run main.go [file of words*]
```

If you want to have more fun, you can run the following command:
```bash
go run main.go [file of words*] --hard
```

*The file of words is mandatory, it can be `words.txt` or `words2.txt` or `words3.txt`.

If you want to use the command line version of the game without big letter and big hangman, you can run the following command:
```bash
go run main.go [file of words*] --low
```

Either if you want to use the command line version of the game with big letter and big hangman, you can run the following command:
```bash
go run main.go [file of words*] --high
```

We also added a lot of flags to customize the game. You can see all the flags with the following command:
```bash
go run main.go --help
```

## Examples

Here are some examples of the game:

```bash
go run main.go words.txt
```
To have the game with the default settings.

```bash
go run main.go words.txt --hard
```
To have the game with the default settings and the hard mode.

```bash
go run main.go words.txt --low
```
To have the game with the default settings and the low quality.

```bash
go run main.go words.txt --high
```
To have the game with the default settings and the high quality.

```bash
go run main.go words.txt -ac
```
To have the game with the default settings and the auto clear of the terminal.

## **Enjoy**
![Pendu](http://images1.memedroid.com/images/UPLOADED19/50984a8283be2.jpeg)