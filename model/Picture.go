package model

type Picture struct {
	UUID        string   `json:"uuid"`
	Path        string   `json:"path"`
	AdderName   string   `json:"adderName"`
	Fingerprint string   `json:"fingerprint"`
	Like        []string `json:"like"`
	UUIDEvent   string   `json:"uuidEvent"`
}

type SimplePicture struct {
	UUID string `json:"uuid"`
	Like int    `json:"like"`
}
