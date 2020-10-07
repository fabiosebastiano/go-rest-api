package repository

import (
	"fmt"

	"com.github/fabiosebastiano/go-rest-api/entity"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type dynamoDBRepo struct {
	tableName string
}

const ()

// NewDynamoDBRepository create a new repo
func NewDynamoDBRepository() PostRepository {
	return &dynamoDBRepo{
		tableName: "posts",
	}
}

func CreateDynamoDBClient() *dynamodb.DynamoDB {
	// crea sessione usando credenziali .aws/credential e .aws/config
	sess := session.Must(
		session.NewSessionWithOptions(
			session.Options{
				SharedConfigState: session.SharedConfigEnable,
			}))

	return dynamodb.New(sess)
}

func (repo *dynamoDBRepo) Save(post *entity.Post) (*entity.Post, error) {

	dynamoDBClient := CreateDynamoDBClient()

	attributeValue, err := dynamodbattribute.MarshalMap(post)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	fmt.Println(attributeValue)

	item := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(repo.tableName),
	}

	_, err = dynamoDBClient.PutItem(item)

	if err != nil {
		fmt.Printf("ERRORE DURANTE PUT DELL'ITEM" + err.Error())
		return nil, err
	}

	return post, nil
}

func (repo *dynamoDBRepo) FindAll() ([]entity.Post, error) {
	dynamoDBClient := CreateDynamoDBClient()

	params := &dynamodb.ScanInput{
		TableName: aws.String(repo.tableName),
	}

	results, err := dynamoDBClient.Scan(params)
	if err != nil {
		return nil, err
	}

	var posts []entity.Post = []entity.Post{}

	for _, i := range results.Items {
		post := entity.Post{}
		err := dynamodbattribute.UnmarshalMap(i, &post)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (repo *dynamoDBRepo) FindByID(id string) (*entity.Post, error) {

	dynamoDBClient := CreateDynamoDBClient()

	result, err := dynamoDBClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repo.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(id),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	post := entity.Post{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &post)
	if err != nil {
		panic(err)
	}
	return &post, nil
}
func (repo *dynamoDBRepo) Delete(postID string) error {
	dynamoDBClient := CreateDynamoDBClient()
	fmt.Println("DELETE IN REPO: ", postID)
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(string(postID)),
			},
		},
		TableName: aws.String(repo.tableName),
	}

	_, err := dynamoDBClient.DeleteItem(input)

	if err != nil {
		return err
	}

	return nil
}
