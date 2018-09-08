package user

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

func NewLoggingMiddleware(next UserService) UserService {
	logger := log.NewLogfmtLogger(os.Stdout)
	return &loggingMiddleware{logger, next}
}

func (mw loggingMiddleware) GetUser(ctx context.Context, id int) (output User, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "getuser",
			"input", id,
			"output", fmt.Sprintf("%v", output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.GetUser(ctx, id)
	return
}

func (mw loggingMiddleware) CreateUser(ctx context.Context, name string, age int) (id int, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "createuser",
			"input", fmt.Sprintf("{name:%v, age: %v}", name, age),
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	id, err = mw.next.CreateUser(ctx, name, age)
	return
}
