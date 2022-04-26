package response

import "mime/multipart"

type UploadFileQ struct {
	Name string                `json:"name" binding:"required"`
	File *multipart.FileHeader `json:"file" swaggerignore:"true" binding:"required"`
	// swagger 中可能有问题，故采用 swaggerignore
}
