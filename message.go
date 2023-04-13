package pushover

type MessageConfig struct {
	Message          string
	Attachment       string
	AttachmentBase64 string
	AttachmentType   string
	Device           string
	HTML             string
	Priority         int8
	Sound            string
	Timestamp        string
	Title            string
	URL              string
	URLTitle         string
}

func NewMesssage(mc MessageConfig) *MessageConfig {
	return &MessageConfig{
		Message:          mc.Message,
		Attachment:       mc.Attachment,
		AttachmentBase64: mc.AttachmentBase64,
		AttachmentType:   mc.AttachmentType,
		Device:           mc.Device,
		HTML:             mc.HTML,
		Priority:         mc.Priority,
		Sound:            mc.Sound,
		Timestamp:        mc.Timestamp,
		Title:            mc.Title,
		URL:              mc.URL,
		URLTitle:         mc.URLTitle,
	}
}
