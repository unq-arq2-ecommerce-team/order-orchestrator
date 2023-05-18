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

const (
	contentTypeKey  = "Content-Type"
	contentTypeJson = "application/json"
)

func MakeAndDoRequestWithBody(ctx context.Context, logger model.Logger, client *http.Client, httpMethod, url string, contentTypeHeader string, body interface{}) (*http.Response, []byte, error) {
	log := logger.WithFields(model.LoggerFields{"url": url})
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("http error marshaling body to json")
		return nil, nil, err
	}
	log = logger.WithFields(model.LoggerFields{"reqBodyRaw": string(reqBody)})
	req, err := http.NewRequestWithContext(ctx, httpMethod, url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Errorf("error when create request object")
		return nil, nil, err
	}
	req.Header.Set(contentTypeKey, contentTypeHeader)
	return doRequest(logger, client, req)
}

// MakeAndDoRequestWithNoBody : return http.Response, response body as []byte, error
func MakeAndDoRequestWithNoBody(ctx context.Context, logger model.Logger, client *http.Client, httpMethod, url string) (*http.Response, []byte, error) {
	log := logger.WithFields(model.LoggerFields{"url": url})
	req, err := http.NewRequestWithContext(ctx, httpMethod, url, nil)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Errorf("error when create request object")
		return nil, nil, err
	}
	return doRequest(logger, client, req)
}

func doRequest(log model.Logger, client *http.Client, req *http.Request) (*http.Response, []byte, error) {
	sw := time.Now()
	res, err := client.Do(req)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("http error do request")
		return nil, nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	log.Debugf("request finished in %s", time.Since(sw))

	var rawBody []byte
	if rawBodyFromRes, err := io.ReadAll(res.Body); err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Errorf("not returning body because reading failed")
	} else {
		rawBody = rawBodyFromRes
	}
	log.Debug(string(rawBody))

	return res, rawBody, nil
}

func IsStatusCode2XX(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}
