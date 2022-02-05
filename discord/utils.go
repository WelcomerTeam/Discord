package discord

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"golang.org/x/xerrors"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func bytesToBase64Data(b []byte) (data string, err error) {
	mime, err := getImageMimeType(b)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	w := base64.NewEncoder(base64.StdEncoding, &out)

	_, err = w.Write(b)
	if err != nil {
		return "", xerrors.Errorf("Failed to base64 bytes: %v", err)
	}

	defer w.Close()

	return "data:" + mime + ";base64," + out.String(), nil
}

func getImageMimeType(b []byte) (mimeType string, err error) {
	switch {
	case bytes.Equal(b[0:8], []byte{137, 80, 78, 71, 13, 10, 26, 10}):
		return "image/png", nil
	case bytes.Equal(b[0:3], []byte{255, 216, 255}) ||
		bytes.Equal(b[6:10], []byte("JFIF")) ||
		bytes.Equal(b[6:10], []byte("Exif")):
		return "image/jpeg", nil
	case bytes.Equal(b[0:6], []byte{71, 73, 70, 56, 55, 97}) ||
		bytes.Equal(b[0:6], []byte{71, 73, 70, 56, 57, 97}):
		return "image/gif", nil
	case bytes.Equal(b[0:4], []byte("RIFF")) && bytes.Equal(b[8:12], []byte("WEBP")):
		return "image/webp", nil
	default:
		return "", ErrUnsupportedImageType
	}
}

func multipartBodyWithJSON(data interface{}, files []*File) (contentType string, body []byte, err error) {
	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)

	payload, err := jsoniter.Marshal(data)
	if err != nil {
		return "", nil, xerrors.Errorf("Failed to marshal payload: %v", err)
	}

	var part io.Writer

	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", `form-data; name="payload_json"`)
	header.Set("Content-Type", "application/json")

	part, err = writer.CreatePart(header)
	if err != nil {
		return "", nil, xerrors.Errorf("Failed to create part: %v", err)
	}

	_, err = part.Write(payload)
	if err != nil {
		return "", nil, xerrors.Errorf("Failed to write payload: %v", err)
	}

	for i, file := range files {
		header := make(textproto.MIMEHeader)
		header.Set(
			"Content-Disposition",
			fmt.Sprintf(
				`form-data; name="file%d"; filename="%s"`,
				i, quoteEscaper.Replace(file.Name),
			),
		)

		fileContentType := file.ContentType
		if fileContentType == "" {
			fileContentType = "application/octet-stream"
		}

		header.Set("Content-Type", fileContentType)

		part, err = writer.CreatePart(header)
		if err != nil {
			return "", nil, xerrors.Errorf("Failed to create part: %v", err)
		}

		_, err = io.Copy(part, file.Reader)
		if err != nil {
			return "", nil, xerrors.Errorf("Failed to copy file: %v", err)
		}
	}

	err = writer.Close()
	if err != nil {
		return "", nil, xerrors.Errorf("Failed to close writer")
	}

	return writer.FormDataContentType(), requestBody.Bytes(), nil
}
