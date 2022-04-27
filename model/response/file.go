package response

import "mime/multipart"

type UploadFileQ struct {
	Name string                `json:"name" form:"name" `
	File *multipart.FileHeader `json:"file" form:"file" swaggerignore:"true" `
	// swagger 中可能有问题，故采用 swaggerignore
}
