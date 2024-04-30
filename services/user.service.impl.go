package services

import (
	"context"

	"example.com/demo-apis/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService()  {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx: ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	-, err := u.usercollection.Insertone(u.ctx, user)
	return nil
}

// passing ctx bcz helpful to perform the operation within certain period of time

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, query) .Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User,error) {
	var users[]*models.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nill{
		return nil, err
	}
	for cursor.Next(c.ctx){
		var user models.User
		err := cursor.Decode(&user)
		if err != nil{
			return nil, err
		}
		user = append(users, &user)
	}

	if err := cursor.Err(); err!= nil{
		return nill, err
	}

	cursor.Close(u.ctx)
	if len(users) == 0{
		return nil, errors.New("documents not found")
	}
	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.name}}
	update := bson.D(bson.E{Key:"$set", value: bson.D{bson.E{Key:"user_name", Value: user.Name}, bson.E{Key:"user_age", Value: user:Age}, bson.E{Key:"user_address", Value: user.Address}}})
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1{
		return errors.New("no new matched document found for update")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.name}}
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1{
		return errors.New("no new matched document found for update")
	}
	return nil
}
