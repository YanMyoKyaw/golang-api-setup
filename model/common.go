package model

type DBProperties struct {
	Host     string
	Port     string
	UserName string
	Password string
	Database string
	DBType   string
}

type FileUploadInput struct {
	Path string
}
