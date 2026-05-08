package main

import (
	"errors"
	"io"
	"os"
)

// Напишіть програму, яка здійснює читання даних із файлу в одній горутині та запис цих даних в іншій горутині у файл. Назви файлів мають відрізнятися.


func task1FileCopyGoroutines(srcPath, dstPath string) error {
	if srcPath == dstPath {
		return errors.New("назви файлів мають відрізнятися")
	}

	src, err := os.Open(srcPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if werr := os.WriteFile(srcPath, []byte("Demo input for Lab8 task 1.\n"), 0o644); werr != nil {
				return werr
			}
			src, err = os.Open(srcPath)
		}
		if err != nil {
			return err
		}
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	chunks := make(chan []byte, 8)
	errCh := make(chan error, 2)

	go func() {
		defer close(chunks)
		buf := make([]byte, 32*1024)
		for {
			n, rerr := src.Read(buf)
			if n > 0 {
				b := make([]byte, n)
				copy(b, buf[:n])
				chunks <- b
			}
			if rerr == io.EOF {
				errCh <- nil
				return
			}
			if rerr != nil {
				errCh <- rerr
				return
			}
		}
	}()

	go func() {
		for b := range chunks {
			if _, werr := dst.Write(b); werr != nil {
				errCh <- werr
				return
			}
		}
		errCh <- nil
	}()

	rerr := <-errCh
	werr := <-errCh
	if rerr != nil {
		return rerr
	}
	if werr != nil {
		return werr
	}
	return nil
}
