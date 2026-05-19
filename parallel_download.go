package req

import (
	"context"
	"io"
	"os"
	"sync"
)

type ParallelDownload struct {
	url          string
	client       *Client
	concurrency  int
	output       io.Writer
	filename     string
	segmentSize  int64
	perm         os.FileMode
	tempRootDir  string
	tempDir      string
	taskCh       chan *downloadTask
	doneCh       chan struct{}
	wgDoneCh     chan struct{}
	errCh        chan error
	wg           sync.WaitGroup
	taskMap      map[int]*downloadTask
	taskNotifyCh chan *downloadTask
	mu           sync.Mutex
	lastIndex    int
}

func (pd *ParallelDownload) completeTask(task *downloadTask) { _ = "STUB: not implemented"; return }

func (pd *ParallelDownload) popTask(index int) *downloadTask { _ = "STUB: not implemented"; return nil }

func md5Sum(s string) string { _ = "STUB: not implemented"; return "" }

func (pd *ParallelDownload) ensure() error { _ = "STUB: not implemented"; return nil }

// 10MB

func (pd *ParallelDownload) SetSegmentSize(segmentSize int64) *ParallelDownload {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ParallelDownload) SetTempRootDir(tempRootDir string) *ParallelDownload {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ParallelDownload) SetFileMode(perm os.FileMode) *ParallelDownload {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ParallelDownload) SetConcurrency(concurrency int) *ParallelDownload {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ParallelDownload) SetOutput(output io.Writer) *ParallelDownload {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ParallelDownload) SetOutputFile(filename string) *ParallelDownload {
	_ = "STUB: not implemented"
	return nil
}

func getRangeTempFile(rangeStart, rangeEnd int64, workerDir string) string {
	_ = "STUB: not implemented"
	return ""
}

type downloadTask struct {
	index                int
	rangeStart, rangeEnd int64
	tempFilename         string
	tempFile             *os.File
}

func (pd *ParallelDownload) handleTask(t *downloadTask, ctx ...context.Context) {
	_ = "STUB: not implemented"
	return
}

func (pd *ParallelDownload) startWorker(ctx ...context.Context) { _ = "STUB: not implemented"; return }

func (pd *ParallelDownload) mergeFile() { _ = "STUB: not implemented"; return }

func (pd *ParallelDownload) Do(ctx ...context.Context) error { _ = "STUB: not implemented"; return nil }

func (pd *ParallelDownload) getOutputFile() (io.Writer, error) {
	_ = "STUB: not implemented"
	return *new(io.Writer), nil
}
