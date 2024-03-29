package response

import (
	"pep/config"
)

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}

type BoardResponse struct {
	Brs []BoardRecord
}

type BoardRecord struct {
	Msg string `json:"msg"`
}
