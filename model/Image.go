package model

type ImageFile struct {
	ID int `json:"image_id"`
	FileUrl string `json:"file_url"`
	CreateTime int `json:"create_time"`
}

func (ImageFile) TableName() string {
	return "image_file"
}
