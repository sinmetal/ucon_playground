package main

import (
	"context"
	"net/http"

	"github.com/favclip/ucon/v3"
	"github.com/favclip/ucon/v3/swagger"
)

func Setup() {
	var info *swagger.HandlerInfo

	info = swagger.NewHandlerInfo(ImageHandler)
	ucon.Handle("GET", "/api/image/", info)
	info.Description, info.Tags = "", []string{}
}

type ImageHandlerReq struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

type ImageHandlerResp struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

func ImageHandler(ctx context.Context, form *ImageHandlerReq, w http.ResponseWriter, r *http.Request) (*ImageHandlerResp, error) {
	resp := &ImageHandlerResp{
		ID:   form.ID,
		Size: form.Size,
	}
	return resp, nil
}
