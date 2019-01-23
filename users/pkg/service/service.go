package service

import (
	"context"
	"errors"
	"go-microservice-base/users/pkg/db"
	"go-microservice-base/users/pkg/io"

	"go-microservice-base/notificator/pkg/grpc/pb"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type Config struct {
	ConsulAddress string
	ConsulPort    string
}

// UsersService describes the service.
type UsersService interface {
	// Add your methods here
	Health(ctx context.Context) (status bool)
	Create(ctx context.Context, user io.User) (error error)
	GetById(ctx context.Context, id string) (u io.User, error error)
	Login(ctx context.Context, auth io.Authentication) (token string, error error)
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

func (b *basicUsersService) Create(ctx context.Context, user io.User) (error error) {
	l := b.logger

	c, err := db.GetUsersCollection(b.logger)
	if err != nil {
		return err
	}
	defer c.Database.Session.Close()

	user.Id = bson.NewObjectId()
	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(fmt.Sprintf("error hashing password: %v", err))
	}
	user.Password = string(hashedPass)

	l.Log("user.Password",user.Password)
	error = c.Insert(&user)

	if error == nil {
		//notify users

		reply, err := b.notificatorServiceClient.Notify(context.Background(), &pb.NotifyRequest{
			Channel: "leslie channel",
			Message: "Hi Leslie! Thank you for registrating in go-microservice-base...",
		})

		if reply != nil {
			l.Log("notificator_reply", reply)
			// TODO handle reply success
		}

		if err != nil {
			l.Log("notificator_err", err)
			// TODO handle reply failure
		}
	} else {
		l.Log("error", error)
	}

	return error
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

func (b *basicUsersService) Login(ctx context.Context, auth io.Authentication) (token string, error error) {
	// TODO implement the business logic of Login
	l := b.logger

	c, err := db.GetUsersCollection(b.logger)
	if err != nil {
		return token, err
	}
	defer c.Database.Session.Close()
	var user io.User
	error = c.Find(bson.M{"username": auth.Username}).One(&user)
	if error != nil {
		return "",errors.New("invalid username or password")
	}
	//l.Log("logged in user", u.String())

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auth.Password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	l.Log("auth", auth)
	return token, error
}
