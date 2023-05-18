package http

import (
	"context"
	"encoding/json"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model/exception"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/config"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/logger"
	"net/http"
	"strconv"
	"strings"
)

type customerRepository struct {
	logger      model.Logger
	client      *http.Client
	findByIdUrl string
}

func NewCustomerRepository(baseLogger model.Logger, client *http.Client, customerEndpoints *config.CustomerEndpoints) model.CustomerRepository {
	return customerRepository{
		logger:      baseLogger.WithFields(logger.Fields{"logger": "http.CustomerRepository"}),
		client:      client,
		findByIdUrl: customerEndpoints.FindById,
	}
}

func (repo customerRepository) FindById(ctx context.Context, customerId int64) (*model.Customer, error) {
	log := repo.logger.WithRequestId(ctx).WithFields(logger.Fields{"findByIdUrl": repo.findByIdUrl, "customerId": customerId})
	log.Debugf("http find customer by id")
	url := strings.Replace(repo.findByIdUrl, "{customerId}", strconv.FormatInt(customerId, 10), -1)
	log = repo.logger.WithFields(logger.Fields{"url": url})

	res, rawBody, err := MakeAndDoRequestWithNoBody(ctx, log, repo.client, http.MethodGet, url)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when make and do request http")
		return nil, err
	}
	switch statusCode := res.StatusCode; {
	case IsStatusCode2XX(statusCode):
		log.Debugf("Raw body: %s", string(rawBody))
		var customer model.Customer
		err = json.Unmarshal(rawBody, &customer)
		return &customer, nil
	case statusCode == http.StatusNotFound:
		return nil, exception.CustomerNotFound{Id: customerId}
	default:
		return nil, NewUnexpectedError("customer repository", statusCode, url)
	}
}
