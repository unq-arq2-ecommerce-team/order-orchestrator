package http

import (
	"context"
	"encoding/json"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model/exception"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/config"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/dto"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type orderRepository struct {
	logger       model.Logger
	client       *http.Client
	findByIdUrl  string
	createUrl    string
	confirmUrl   string
	deliveredUrl string
}

func NewOrderRepository(baseLogger model.Logger, client *http.Client, orderEndpoints *config.OrderEndpoints) model.OrderRepository {
	return orderRepository{
		logger:       baseLogger.WithFields(model.LoggerFields{"logger": "http.OrderRepository"}),
		client:       client,
		findByIdUrl:  orderEndpoints.FindById,
		createUrl:    orderEndpoints.Create,
		confirmUrl:   orderEndpoints.Confirm,
		deliveredUrl: orderEndpoints.Delivered,
	}
}

func (repo orderRepository) FindById(ctx context.Context, orderId int64) (*model.Order, error) {
	log := repo.logger.WithFields(model.LoggerFields{"findByIdUrl": repo.findByIdUrl, "orderId": orderId})
	log.Debugf("http find order by id")
	url := strings.Replace(repo.findByIdUrl, "{orderId}", strconv.FormatInt(orderId, 10), -1)
	log = repo.logger.WithFields(model.LoggerFields{"url": url})

	res, err := MakeAndDoRequest(ctx, log, repo.client, http.MethodGet, url, nil)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when make and do request http")
		return nil, err
	}

	switch statusCode := res.StatusCode; {
	case IsStatusCode2XX(statusCode):
		rawBody, _ := io.ReadAll(res.Body)
		log = log.WithFields(model.LoggerFields{"bodyRaw": rawBody})
		var order model.Order
		err = json.Unmarshal(rawBody, &order)
		return &order, nil
	case statusCode == http.StatusNotFound:
		return nil, exception.OrderNotFound{Id: orderId}
	default:
		return nil, NewUnexpectedError("order repository", statusCode, url)
	}
}

func (repo orderRepository) Create(ctx context.Context, order model.Order) (int64, error) {
	url := repo.createUrl
	log := repo.logger.WithFields(model.LoggerFields{"createOrderUrl": url, "order": order})
	log.Debugf("http create order")
	log = repo.logger.WithFields(model.LoggerFields{"url": url})
	orderDTO := dto.NewOrderDTO(order)
	log.Debugf("orderDTO: %s", orderDTO)

	res, err := MakeAndDoRequest(ctx, log, repo.client, http.MethodPost, url, orderDTO)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when make and do request http")
		return 0, err
	}

	switch statusCode := res.StatusCode; {
	case IsStatusCode2XX(statusCode):
		rawBody, _ := io.ReadAll(res.Body)
		log = log.WithFields(model.LoggerFields{"bodyRaw": rawBody})
		var orderIdRes dto.IdRes
		err = json.Unmarshal(rawBody, &orderIdRes)
		return orderIdRes.Id, nil
	case statusCode == http.StatusNotFound:
		return 0, exception.ProductNotFound{Id: order.Product.Id}
	case statusCode == http.StatusNotAcceptable:
		return 0, exception.ProductWithNoStock{Id: order.Product.Id}
	default:
		return 0, NewUnexpectedError("order repository", statusCode, url)
	}
}

func (repo orderRepository) Confirm(ctx context.Context, orderId int64) error {
	log := repo.logger.WithFields(model.LoggerFields{"confirmUrl": repo.confirmUrl, "orderId": orderId})
	log.Debugf("http confirm order by id")
	url := strings.Replace(repo.confirmUrl, "{orderId}", strconv.FormatInt(orderId, 10), -1)
	log = repo.logger.WithFields(model.LoggerFields{"url": url})

	res, err := MakeAndDoRequest(ctx, log, repo.client, http.MethodPost, url, nil)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when make and do request http")
		return err
	}

	switch statusCode := res.StatusCode; {
	case IsStatusCode2XX(statusCode):
		rawBody, _ := io.ReadAll(res.Body)
		log = log.WithFields(model.LoggerFields{"bodyRaw": rawBody})
		var order model.Order
		err = json.Unmarshal(rawBody, &order)
		return nil
	case statusCode == http.StatusNotFound:
		return exception.OrderNotFound{Id: orderId}
	case statusCode == http.StatusNotAcceptable:
		return exception.OrderInvalidTransitionState{Id: orderId}
	default:
		return NewUnexpectedError("order repository", statusCode, url)
	}
}

func (repo orderRepository) Delivered(ctx context.Context, orderId int64) error {
	log := repo.logger.WithFields(model.LoggerFields{"deliveredUrl": repo.deliveredUrl, "orderId": orderId})
	log.Debugf("http delivered order by id")
	url := strings.Replace(repo.deliveredUrl, "{orderId}", strconv.FormatInt(orderId, 10), -1)
	log = repo.logger.WithFields(model.LoggerFields{"url": url})

	res, err := MakeAndDoRequest(ctx, log, repo.client, http.MethodPost, url, nil)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when make and do request http")
		return err
	}

	switch statusCode := res.StatusCode; {
	case IsStatusCode2XX(statusCode):
		rawBody, _ := io.ReadAll(res.Body)
		log = log.WithFields(model.LoggerFields{"bodyRaw": rawBody})
		var order model.Order
		err = json.Unmarshal(rawBody, &order)
		return nil
	case statusCode == http.StatusNotFound:
		return exception.OrderNotFound{Id: orderId}
	case statusCode == http.StatusNotAcceptable:
		return exception.OrderInvalidTransitionState{Id: orderId}
	default:
		return NewUnexpectedError("order repository", statusCode, url)
	}
}
