package model

type Picture struct {
	Path        string `json:"path"`
	AdderName   string `json:"adderName"`
	Fingerprint string `json:"fingerprint"`
}
