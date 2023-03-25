package object

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/kara9renai/yokattar-go/pkg/config"
)

type (
	Attachment struct {
		ID          int64  `json:"id"`
		Type        string `json:"type"`
		URL         string `json:"url"`
		Description string `json:"description"`
	}
)

func (at *Attachment) CopyFile(file multipart.File, fileName string) error {
	err := os.MkdirAll(config.ImagePath, os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return err
	}
	return nil
}

func (at *Attachment) CreateFileName(h *multipart.FileHeader) string {
	// URLの擬似乱数文字列を生成
	c := 40
	b := make([]byte, c)
	rand.Read(b)
	s := base64.URLEncoding.EncodeToString(b)
	ext := filepath.Ext(h.Filename)
	var fileName string = fmt.Sprintf(config.ImagePath+"%v%s", s, ext)
	return fileName
}
