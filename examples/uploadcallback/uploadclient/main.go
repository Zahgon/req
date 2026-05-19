package main

import (
	"fmt"
	"io"
	"time"

	"github.com/imroc/req/v3"
)

type SlowReader struct {
	Size int
	n    int
}

func (r *SlowReader) Close() error { _ = "STUB: not implemented"; return nil }

func (r *SlowReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func main() {
	size := 10 * 1024 * 1024
	req.SetFileUpload(req.FileUpload{
		ParamName: "file",
		FileName:  "test.txt",
		GetFileContent: func() (io.ReadCloser, error) {
			return &SlowReader{Size: size}, nil
		},
		FileSize: int64(size),
	}).SetUploadCallbackWithInterval(func(info req.UploadInfo) {
		fmt.Printf("%s: %.2f%%\n", info.FileName, float64(info.UploadedSize)/float64(info.FileSize)*100.0)
	}, 30*time.Millisecond).Post("http://127.0.0.1:8888/upload")
}
