package media

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kara9renai/yokattar-go/app/handler/httperror"
)

const (
	MaxUploadSize = 1024 * 1024
	imagePath     = "./uploadimages/"
)

type AddRequest struct {
	ID          int64
	Type        string
	Url         string
	Description string
}

// Handle Request for POST `v1/media`
func (h *handler) Upload(w http.ResponseWriter, r *http.Request) {

	var req AddRequest

	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)

	// ParseMutipartFormはリクエストボディを`multipart/form-data`として解析する関数
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}
	defer file.Close()

	err = os.MkdirAll(imagePath, os.ModePerm)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	// URLの擬似乱数文字列を生成
	c := 40
	b := make([]byte, c)
	rand.Read(b)

	var n string = fmt.Sprintf(imagePath+"%v%s", base64.URLEncoding.EncodeToString(b), filepath.Ext(fileHeader.Filename))

	f, err := os.Create(n)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	req.ID = 123
	req.Type = filepath.Ext(fileHeader.Filename)
	req.Url = n

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(req); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
