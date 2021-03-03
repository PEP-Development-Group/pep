package response

import "pep/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
