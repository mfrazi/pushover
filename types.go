package pushover

type DefaultResponse struct {
	Status  int    `json:"status"`
	Request string `json:"request"`
}

type SoundListResponse struct {
	Sounds map[string]string `json:"sounds"`
	DefaultResponse
}
