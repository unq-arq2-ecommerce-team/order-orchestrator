package http

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"io"
	"net/http"
	"time"
)

const contentTypeKey = "Content-Type"

func MakeAndDoRequest(ctx context.Context, logger model.Logger, client *http.Client, httpMethod, url string, body interface{}) (*http.Response, error) {
	log := logger.WithFields(model.LoggerFields{"url": url})
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("http error marshaling body to json")
		return nil, err
	}
	log = logger.WithFields(model.LoggerFields{"requestBody": string(reqBody)})
	req, err := http.NewRequestWithContext(ctx, httpMethod, url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Errorf("error when create request object")
		return nil, err
	}
	req.Header.Set(contentTypeKey, http.DetectContentType(reqBody))
	sw := time.Now()
	res, err := client.Do(req)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("http error do request")
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	log.Debugf("request finished in %s", time.Since(sw))
	return res, nil
}
