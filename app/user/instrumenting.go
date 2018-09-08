package user

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	createdId      metrics.Histogram
	next           UserService
}

func NewInstrumentingMiddleware(svc UserService) UserService {
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Name: "request_count",
		Help: "Number of requests recieved.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Name: "request_latency_microseconds",
		Help: "Total duration of requests in microseconds.",
	}, fieldKeys)
	createdId := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Name: "created_id",
		Help: "The ID created.",
	}, []string{})
	return &instrumentingMiddleware{requestCount, requestLatency, createdId, svc}
}

func (mw instrumentingMiddleware) GetUser(ctx context.Context, id int) (output User, err error) {
	defer func(begin time.Time) {
		methodField := []string{"method", "getuser"}
		errorField := []string{"error", fmt.Sprintf("%v", err)}
		mw.requestCount.With(methodField...).With(errorField...).Add(1)
		mw.requestLatency.With(methodField...).With(errorField...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output, err = mw.next.GetUser(ctx, id)
	return
}

func (mw instrumentingMiddleware) CreateUser(ctx context.Context, name string, age int) (id int, err error) {
	defer func(begin time.Time) {
		methodField := []string{"method", "createuser"}
		errorField := []string{"error", fmt.Sprintf("%v", err)}
		mw.requestCount.With(methodField...).With(errorField...).Add(1)
		mw.requestLatency.With(methodField...).With(errorField...).Observe(time.Since(begin).Seconds())
		mw.createdId.Observe(float64(id))
	}(time.Now())
	id, err = mw.next.CreateUser(ctx, name, age)
	return
}
