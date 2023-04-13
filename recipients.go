package pushover

type RecipientsConfig struct {
	UserKey string
}

func NewRecipients(rc RecipientsConfig) *RecipientsConfig {
	return &RecipientsConfig{
		UserKey: rc.UserKey,
	}
}
