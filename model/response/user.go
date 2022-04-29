package response

type CommonA struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type LoginQ struct {
	Name     string `json:"name" binding:"min=3,max=100,required"`
	Password string `json:"password" binding:"gte=6,required"`
}

type RegisterQ struct {
	Name     string `json:"name" binding:"min=3,max=100,required"`
	Password string `json:"password" binding:"gte=6,required"`
}

type LoginA struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
	ID      uint64 `json:"id"`
}

type UploadFileA struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Path    string `json:"path"`
}

type GetUserA struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Name    string `json:"name"`
}
