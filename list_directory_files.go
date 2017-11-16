package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    dirPath := "."  // Get current dir

    // Save our files in a list (like an array)
    var fileList []string

    // Walk walks the file tree rooted at root
    filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            //fileList = append(fileList, path)
            fileList = append(fileList, info.Name())
        }
        return nil
    })

    // Print the files
    // uses of '_' because we don't need this var
    for _, file := range fileList {
        fmt.Println(file)
    }
}
