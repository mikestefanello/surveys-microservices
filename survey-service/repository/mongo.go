package repository

import (
	"context"

	"github.com/mikestefanello/surveys-microservices/survey-service/config"
	"github.com/mikestefanello/surveys-microservices/survey-service/survey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoSurveyRepository struct {
	client *mongo.Client
	config config.MongoConfig
}

// NewSurveyMongoRepository creates a new Mongo DB survey repository
func NewSurveyMongoRepository(cfg config.MongoConfig) (survey.Repository, error) {
	repo := &mongoSurveyRepository{config: cfg}
	err := repo.connect()
	return repo, err
}

func (r *mongoSurveyRepository) connect() error {
	ctx, cancel := r.contextWithTimeout()
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(r.config.URL))
	if err != nil {
		return err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	r.client = client

	return nil
}

func (r *mongoSurveyRepository) contextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), r.config.Timeout)
}

func (r *mongoSurveyRepository) getCollection() *mongo.Collection {
	return r.client.Database(r.config.DB).Collection("surveys")
}

func (r *mongoSurveyRepository) Insert(survey *survey.Survey) error {
	ctx, cancel := r.contextWithTimeout()
	defer cancel()

	_, err := r.getCollection().InsertOne(ctx, survey)

	return err
}

func (r *mongoSurveyRepository) LoadByID(id string) (*survey.Survey, error) {
	ctx, cancel := r.contextWithTimeout()
	defer cancel()

	s := &survey.Survey{}
	err := r.getCollection().FindOne(ctx, bson.M{"id": id}).Decode(&s)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, survey.ErrNotFound
		}
		return nil, err
	}

	return s, nil
}

func (r *mongoSurveyRepository) Load() (*survey.Surveys, error) {
	ctx, cancel := r.contextWithTimeout()
	defer cancel()

	s := &survey.Surveys{}
	opts := options.Find()
	opts.SetSort(bson.D{{"createdAt", -1}})
	opts.SetLimit(25)
	cursor, err := r.getCollection().Find(ctx, bson.M{}, opts)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, s)

	return s, err
}
