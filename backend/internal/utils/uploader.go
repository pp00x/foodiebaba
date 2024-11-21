package utils

import (
    "fmt"
    "io"
    "mime/multipart"
    "os"
    "path/filepath"
    "time"
)

func UploadFile(file *multipart.FileHeader) (string, error) {
    dst := fmt.Sprintf("uploads/%d_%s", time.Now().UnixNano(), file.Filename)
    if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
        return "", err
    }
    if err := saveFile(file, dst); err != nil {
        return "", err
    }
    return "/" + dst, nil // Return the relative URL
}

func saveFile(file *multipart.FileHeader, dst string) error {
    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, src)
    return err
}