package model

type User struct {
	UUID                  string `json:"uuid"`
	BrowserFingerPrinting string `json:"browser_fingerprinting"`
	Ip                    string `json:"ip"`
}
