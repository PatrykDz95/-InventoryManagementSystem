package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorsk/server/common"
	"net/http"
	"strings"
	"time"
)

type Service struct {
	collection *mongo.Collection
}

func NewUserService(db *mongo.Database) *Service {
	return &Service{
		collection: db.Collection("users"),
	}
}

func (s *Service) Login(c *gin.Context, email string, password string) {
	var user User
	err := s.collection.FindOne(c.Request.Context(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	if !common.CheckPasswordHash(password, user.PasswordHash) {
		token, err := common.GenerateToken(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, token)
	}
}

func (s *Service) Create(c *gin.Context, user *User) {
	user.ID = primitive.NewObjectID()
	hashedPassword, err := common.HashPassword(user.PasswordHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.PasswordHash = hashedPassword
	user.CreatedAt = time.Now()
	if _, err := s.collection.InsertOne(c.Request.Context(), user); err != nil {
		errorCode := ExtractErrorCode(err)
		if errorCode == "E11000" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists"})
			return
		}
	}
	token, err := common.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, token)
}

func (s *Service) GetUserByID(ctx context.Context, id string) (*User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user User
	err = s.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	return &user, err
}

func (s *Service) UpdateUser(ctx context.Context, id string, user *User) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": user})
	return err
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func ExtractErrorCode(err error) string {
	errMsg := err.Error()
	if strings.Contains(errMsg, "E11000") {
		return "E11000"
	}
	return ""
}
