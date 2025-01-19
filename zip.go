package ziputil

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
)

type Zip interface {
	Bytes() []byte
	Close() error
	AddFile(name string, contents []byte) error
	AddFileBase64(name string, b64EncodedContents string) error
}

type zipFile struct {
	buff    *bytes.Buffer
	zWriter *zip.Writer
}

func New() Zip {
	buff := new(bytes.Buffer)
	return &zipFile{
		buff:    buff,
		zWriter: zip.NewWriter(buff),
	}
}

// Bytes returns the zip file as a byte slice.
// This method should be called after all files have been added to the zip file.
// It will close the internal zip writer before returning the byte slice.
func (z *zipFile) Bytes() []byte {
	z.zWriter.Close()

	return z.buff.Bytes()
}

// Close closes the internal zip writer.
func (z *zipFile) Close() error {
	return z.zWriter.Close()
}

// AddFile adds a file to the zip archive.
func (z *zipFile) AddFile(name string, contents []byte) error {
	file, err := z.zWriter.Create(name)
	if err != nil {
		return err
	}

	_, err = file.Write(contents)
	if err != nil {
		return err
	}

	return nil
}

// AddFileBase64 adds a file to the zip archive.
// The contents are decoded using {base64.StdEncoding} and the bytes are added to the zip file.
func (z *zipFile) AddFileBase64(name string, b64EncodedContents string) error {
	contents, err := base64.StdEncoding.DecodeString(b64EncodedContents)
	if err != nil {
		return err
	}

	return z.AddFile(name, contents)
}
