package repositories

import (
	"context"
	"github.com/99-66/go-gin-project-template/consts"
	"github.com/99-66/go-gin-project-template/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
	MongoDB *mongo.Client
}

func ProvideTodoRepository(DB *gorm.DB, MongoDB *mongo.Client) TodoRepository {
	return TodoRepository{DB: DB, MongoDB: MongoDB}
}

// FindAll
func (t *TodoRepository) FindAll() (todo []models.Todo, err error) {
	return todo, t.DB.Find(&todo).Error
}

// FindById
func (t *TodoRepository) FindById(id uint) (todo models.Todo, err error) {
	return todo, t.DB.First(&todo, id).Error
}


// Create
func (t *TodoRepository) Create(todo *models.Todo) error {
	if err := t.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

// Update
func (t *TodoRepository) Update(todo, updateTodo *models.Todo) error {
	if err := t.DB.Model(&todo).Updates(updateTodo).Error; err != nil {
		return err
	}
	return nil
}

// DeleteById
func (t *TodoRepository) DeleteById(todo *models.Todo, id uint) error {
	if err := t.DB.Where("id = ?", id).Delete(todo).Error; err != nil {
		return err
	}
	return nil
}

// FindAllCollections
func (t *TodoRepository) FindAllCollection() ([]models.Location, error) {
	dbName := consts.MONGODB_DATABASE
	var Locations []models.Location

	collection := t.MongoDB.Database(dbName).Collection("chauffeurLocation")
	docs, err := collection.Find(context.TODO(), bson.D{})

	defer docs.Close(context.TODO())
	if err != nil {
		return nil, err
	}
	for docs.Next(context.TODO()) {
		var Loc models.Location
		err := docs.Decode(&Loc)
		if err != nil {
			return nil, err
		}
		Locations = append(Locations, Loc)
	}

	return Locations, nil
}