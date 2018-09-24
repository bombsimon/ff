<h1 align="center">
  <img src="https://github.com/bombsimon/ff/blob/master/img/logo.png" alt="FF" width="200">
  <br>
  Find Files
  <br>
</h1>

<h4 align="center">Find files with glob syntax.</h4>

<p align="center">
  <a href="https://forthebadge.com/">
      <img src="https://forthebadge.com/images/badges/its-not-a-lie-if-you-believe-it.svg">
  </a>
  <a href="https://forthebadge.com/">
      <img src="https://forthebadge.com/images/badges/fuck-it-ship-it.svg">
  </a>
</p>

## Description

Find Files is a quick way to combine powerful features from `filepath` and `os` to
find files in your project based on patterns.

The idea came from when I was build a linter and wanted a quick way to parse
all my arguments to use as input for files to lint. I took a look at most of
the [awesome-go-linters](https://github.com/golangci/awesome-go-linters) at the
time and it was obvious that there was no common way to handle user input to
match files.

## Matches

Find Files will take any pattern which may or may not describe hierarchical
name and return a list of files matching the pattern.

By supporting an optional boolean flag to parse recursive files can be found
with slight difference to the traditional way to find files.

[Reference and inspiration](https://github.com/begin/globbing).
