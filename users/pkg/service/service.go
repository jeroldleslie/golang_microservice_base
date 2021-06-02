package service

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"

	"github.com/jeroldleslie/golang_microservice_base/notificator/pkg/grpc/pb"
	"github.com/jeroldleslie/golang_microservice_base/users/pkg/db"
	"github.com/jeroldleslie/golang_microservice_base/users/pkg/io"
)

type Config struct {
	ConsulAddress string
	ConsulPort    string
}

// UsersService describes the service.
type UsersService interface {
	// Add your methods here
	Health(ctx context.Context) (status bool)
	Create(ctx context.Context, user io.User) (u io.User, error error)
	GetById(ctx context.Context, id string) (u io.User, error error)
}

type basicUsersService struct {
	logger                   log.Logger
	notificatorServiceClient pb.NotificatorClient
}

// NewBasicUsersService returns a naive, stateless implementation of UsersService.
func NewBasicUsersService(logger log.Logger) UsersService {

	conn, err := grpc.Dial("notificator:8082", grpc.WithInsecure())
	if err != nil {
		logger.Log("err", err.Error(), "message", "unable to connect to notificator")
		return new(basicUsersService)
	}

	logger.Log("", "connected to notificator")
	return &basicUsersService{
		logger:                   logger,
		notificatorServiceClient: pb.NewNotificatorClient(conn),
	}
}

// New returns a UsersService with all of the expected middleware wired in.
func New(middleware []Middleware, logger log.Logger) UsersService {
	var svc UsersService = NewBasicUsersService(logger)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func bsonObject(idStr string) (bson.ObjectId, error) {
	var _id bson.ObjectId
	if bson.IsObjectIdHex(idStr) == true {
		_id = bson.ObjectIdHex(idStr)
		return _id, nil
	} else {
		return _id, errors.New("invalid id")
	}
}

func (b *basicUsersService) Create(ctx context.Context, user io.User) (u io.User, error error) {
	user.Id = bson.NewObjectId()
	c, err := db.GetUsersCollection(b.logger)
	if err != nil {
		return u, err
	}
	defer c.Database.Session.Close()

	error = c.Insert(&user)

	if error == nil {
		//notify users
		reply, err := b.notificatorServiceClient.Notify(context.Background(), &pb.NotifyRequest{
			Channel: "leslie channel",
			Message: "Hi Leslie! Thank you for registrating in go-microservice-base...",
		})

		if reply != nil {
			b.logger.Log("notificator_reply", reply)
			// TODO handle reply success
		}

		if err != nil {
			b.logger.Log("notificator_err", err)
			// TODO handle reply failure
		}
	}

	return user, error
}

func (b *basicUsersService) GetById(ctx context.Context, id string) (u io.User, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return u, err
	}
	defer session.Close()
	c := session.DB("go-microservice-base").C("users")

	_id, err := bsonObject(id)
	if err != nil {
		return u, err
	}

	error = c.Find(bson.M{"_id": _id}).One(&u)
	return u, error
}

func (b *basicUsersService) Health(ctx context.Context) (status bool) {
	status = true
	return status
}
