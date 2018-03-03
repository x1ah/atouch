package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

var FileHeadMapping = map[string]string{
    "py": "# coding: utf-8\n\n",
    "c": "# include <stdio.h>\n",
}

func checkError(e error) {
    if e != nil {
        panic(e)
    }
}

func absPath(fileName string) (path string) {
    path, err := filepath.Abs(fileName)

    checkError(err)

    return
}

func aTouchFile(fileName string) {
    f, err := os.Create(fileName)
    checkError(err)

    defer f.Close()

    _fileSplits := strings.Split(fileName, ".")
    if len(_fileSplits) == 1 {
        return
    }

    fileType := _fileSplits[len(_fileSplits)-1]
    str, ok := FileHeadMapping[fileType]
    if ok == false {
        return
    }
    _, err = f.WriteString(str)
    checkError(err)
    f.Sync()
}

func main() {
    files := os.Args[1:]
    for _, fileName := range files {
        r := absPath(fileName)
        if _, err := os.Stat(r); err == nil {
            fmt.Printf("File or directory `%s` aleardy exist.\n", fileName)
        } else {
            aTouchFile(fileName)
        }
    }
}
