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

// Handle Request for POST `v1/media`
func (h *handler) Upload(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

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

	var fileName string = fmt.Sprintf(imagePath+"%v%s", base64.URLEncoding.EncodeToString(b), filepath.Ext(fileHeader.Filename))

	f, err := os.Create(fileName)

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

	at := h.app.Dao.Attachment() // domain/repositoryの取得

	attachment, err := at.UploadFile(ctx, fileName)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(attachment); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
