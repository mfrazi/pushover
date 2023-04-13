package pushover

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
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
	form := url.Values{}
	form.Add("token", p.token)
	form.Add("user", rc.UserKey)
	form.Add("message", mc.Message)

	req, err := http.NewRequest(http.MethodPost, PushoverMessagesURL, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

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
