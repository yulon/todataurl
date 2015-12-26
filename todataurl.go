package main

import (
	"os"
	"io"
	"encoding/base64"
	"mime"
	"path/filepath"
	"net/http"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		return
	}
	defer f.Close()

	fh := make([]byte, 512)
	f.Read(fh)

	contType := mime.TypeByExtension(filepath.Ext(os.Args[1]))
	if contType == "" {
		contType = http.DetectContentType(fh)
	}
	fmt.Printf("data:%s;base64,", contType)

	enc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	enc.Write(fh)
	io.Copy(enc, f)
	enc.Close()

	println("")
}
