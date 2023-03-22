package media

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/app/config"
	"github.com/kara9renai/yokattar-go/app/domain/object"
	"github.com/kara9renai/yokattar-go/app/server/handler/httperror"
)

// Handle Request for POST `v1/media`
func (h *handler) Upload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	attachment := new(object.Attachment)
	r.Body = http.MaxBytesReader(w, r.Body, config.MaxUploadSize)
	// ParseMutipartFormはリクエストボディを`multipart/form-data`として解析する関数
	if err := r.ParseMultipartForm(config.MaxUploadSize); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}
	defer file.Close()
	fileName := attachment.CreateFileName(fileHeader)
	err = attachment.CopyFile(file, fileName)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	at := h.app.Dao.Attachment() // domain/repositoryの取得
	attachment, err = at.Save(ctx, fileName)
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
