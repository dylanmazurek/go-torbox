package models

type File struct {
	ID int64 `json:"id"`

	MD5  *string `json:"md5"`
	Hash string  `json:"hash"`

	Size     int64  `json:"size"`
	Zipped   bool   `json:"zipped"`
	S3Path   string `json:"s3_path"`
	Infected bool   `json:"infected"`
	MimeType string `json:"mimetype"`

	AbsolutePath string `json:"absolute_path"`
	Name         string `json:"name"`
	ShortName    string `json:"short_name"`
}
