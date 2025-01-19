## Usage

```go
func BuildZip(pdf, xml FileStruct) ([]byte, error) {
  z := ziputil.Defered()
  z.AddFileBase64(pdf.NameFile, pdf.B64EncodedContents)
  z.AddFileBase64(xml.NameFile, xml.B64EncodedContents)

  return z.Bytes()
}
```