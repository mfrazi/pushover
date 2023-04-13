package pushover

import (
	"io/ioutil"
	"log"
	"net/http"
)

const (
	PushoverMessagesURL        = "https://api.pushover.net/1/messages.json"
	PushoverSoundsURL          = "https://api.pushover.net/1/sounds.json"
	PushoverUsersValidateURL   = "https://api.pushover.net/1/users/validate.json"
	PushoverDevicesValidateURL = "https://api.pushover.net/1/devices/validate.json"
)

// Pushover is a struct that holds the API token
type Pushover struct {
	token string
}

func New(token string) (*Pushover, error) {
	if token == "" {
		return nil, nil
	}

	return &Pushover{
		token: token,
	}, nil
}

func (p *Pushover) SendMessage(rc RecipientsConfig, mc *MessageConfig) error {
	req, err := http.NewRequest(http.MethodPost, PushoverMessagesURL, nil)
	if err != nil {
		return err
	}

	req.PostForm.Add("token", p.token)
	req.PostForm.Add("user", rc.UserKey)
	req.PostForm.Add("message", mc.Message)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println(string(body))

	return nil
}
