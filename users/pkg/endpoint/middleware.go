package endpoint

import (
	"context"
	"fmt"
	"time"
	"errors"
	error1 "go-microservice-base/users/pkg/errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/auth/jwt"
	jwt1 "github.com/dgrijalva/jwt-go"
	"go-microservice-base/users/pkg/utils"
)

// InstrumentingMiddleware returns an endpoint middleware that records
// the duration of each invocation to the passed histogram. The middleware adds
// a single field: "success", which is "true" if no error is returned, and
// "false" otherwise.
func InstrumentingMiddleware(duration metrics.Histogram) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				duration.With("success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
			}(time.Now())

			return next(ctx, request)
		}
	}
}

// LoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// AuthMiddleware returns an endpoint middleware
func AuthMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// Add your middleware logic here
			jwtAuth, err := utils.InitJWTAuthenticationBackend()

			tokenString, ok := ctx.(context.Context).Value(jwt.JWTTokenContextKey).(string)
			if !ok {
				return nil, errors.New(error1.TokenNotExists)
			}

			token, err := jwt1.Parse(tokenString, func(token *jwt1.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt1.SigningMethodRSA); !ok {
					return nil, fmt.Errorf(error1.IncorrectSigningMethod)
				} else {
					return jwtAuth.PublicKey, nil
				}
			})

			if token.Valid {
				return next(ctx, request)
			} else if ve, ok := err.(*jwt1.ValidationError); ok {
				if ve.Errors&jwt1.ValidationErrorMalformed != 0 {
					return nil, errors.New(error1.NotToken)
				} else if ve.Errors&(jwt1.ValidationErrorExpired|jwt1.ValidationErrorNotValidYet) != 0 {
					// Token is either expired or not active yet
					return nil, errors.New(error1.TimeExpired)
				} else {
					return nil, errors.New(error1.InvalidToken)
				}
			} else {
				return nil, errors.New(error1.CantHandleToken)
			}

		}
	}
}
