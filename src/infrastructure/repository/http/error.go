package http

import "fmt"

func NewUnexpectedError(repository string, statusCode int, url string) error {
	return fmt.Errorf("unexpected error with status code %v from %s with url %s", statusCode, repository, url)
}
