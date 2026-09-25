package gotils

import (
	"bytes"
	"errors"
	"io"
	"os"
)

func CopyFile(srcPath string, dstPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}

	defer src.Close()

	out, err := os.Create(dstPath)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}

	return out.Sync()
}

func ReplaceFileBytesExact(filePath string, toReplace []byte, replaceWith []byte) error {
	if len(toReplace) != len(replaceWith) {
		return errors.New("data to replace and replacement data have to have the same length")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	offset := bytes.Index(data, toReplace)
	if offset == -1 {
		return errors.New("failed to find the data to be replaced")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteAt(replaceWith, int64(offset))
	return err
}

func ReplaceFileBytesFromLeading(filePath string, leadingBytes []byte, replaceWith []byte) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	offset := bytes.Index(data, leadingBytes)
	if offset == -1 {
		return errors.New("failed to find leading bytes")
	}

	if len(replaceWith) > len(data)-offset {
		return errors.New("length of replacement data is greater than file data")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteAt(replaceWith, int64(offset))
	return err
}
