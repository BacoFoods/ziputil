package ziputil

import (
	"archive/zip"
	"bytes"
)

type ZipDefered interface {
	Bytes() ([]byte, error)
	AddFile(name string, contents []byte)
	AddFileBase64(name string, b64EncodedContents string)
}

type zipDefered struct {
	zipFile

	err error
}

func Defered(filename string) ZipDefered {
	buff := new(bytes.Buffer)
	return &zipDefered{
		zipFile: zipFile{
			Filename: filename,
			buff:     buff,
			zWriter:  zip.NewWriter(buff),
		},
	}
}

// Bytes returns the zip file as a byte slice.
// This method should be called after all files have been added to the zip file.
// It will close the internal zip writer before returning the byte slice.
//
// If an error occurred while adding files, it will be returned here.
// If an error is returned, it is guaranteed that the zip writer has been closed and the returned error will be the last error that occurred.
// There is no guarantee that the zip file is valid and that a there are no more errors in the process.
func (z *zipDefered) Bytes() ([]byte, error) {
	defer z.zWriter.Close()

	if z.err != nil {
		return nil, z.err
	}

	return z.buff.Bytes(), nil
}

// AddFile adds a file to the zip archive.
func (z *zipDefered) AddFile(name string, contents []byte) {
	z.err = z.zipFile.AddFile(name, contents)
}

// AddFileBase64 adds a file to the zip archive from base64 encoded contents.
// The contents are decoded before adding the file to the zip archive using {base64.StdEncoding}.
func (z *zipDefered) AddFileBase64(name string, b64EncodedContents string) {
	z.err = z.zipFile.AddFileBase64(name, b64EncodedContents)
}
