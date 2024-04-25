package client

import (
	"encoding/base64"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetFile(url string) (string, int, error) {

	httpresponse, err := http.Get(url)
	if err != nil {

		// Is there an Invalid http Response Code?
		if httpresponse == nil {
			return "", 0, err
		}
		return "", httpresponse.StatusCode, err
	}

	defer func() {
		_ = httpresponse.Body.Close()
	}()

	body, err := io.ReadAll(httpresponse.Body)
	if err != nil {
		return "", httpresponse.StatusCode, err
	}

	log.Info("File resource retrieved")

	// Encode to Base64
	encoded := base64.StdEncoding.EncodeToString(body)

	return encoded, httpresponse.StatusCode, nil
}
