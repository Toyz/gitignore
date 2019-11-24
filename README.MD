[![Actions Status](https://github.com/toyz/gitignore/workflows/Go/badge.svg)](https://github.com/toyz/gitignore/actions)


# Gitignore

A simple tool to let you download gitignore files from https://github.com/github/gitignore and also can download license file as well

## What is this?
This is a port of https://github.com/msfeldstein/gitignore so you don't need nodejs to use it

## Installation 
```
$ go get -u github.com/toyz/gitignore
```

## Usage
```
$ gitignore

Usage:
gitignore --types [types | license] -- List supported gitignore types
gitignore --license [file] -- License file to download
gitignore [file] -- Gitignore file to download
```

## Exmaple
```
$ gitignore go 
> .gitignore has been created... Your good to go!

$ gitignore --license mit
> LICENSE has been created... Your good to go!
```