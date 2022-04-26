package response

import "mime/multipart"

type UploadFileQ struct {
	Name string                `form:"name"`
	File *multipart.FileHeader `form:"file" swaggerignore:"true"`
	// swagger 中可能有问题，故采用 swaggerignore
}
