package controllers

import (
	"errors"
	"fmt"
	"github.com/HakimHC/altostratus-golang-api/config"
	"github.com/HakimHC/altostratus-golang-api/models"
	"github.com/HakimHC/altostratus-golang-api/responses"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
)

func putAsteroidInDynamoDB(asteroid models.Asteroid, tableName string) error {
	av, err := dynamodbattribute.MarshalMap(asteroid)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = config.DynamoDB.PutItem(input)
	return err
}

func getAsteroidByField(field map[string]string, tableName string) (*[]models.Asteroid, error) {
	if len(field) != 1 {
		return nil, errors.New("you must filter by only one field")
	}
	var key string
	for k := range field {
		key = k
		break
	}

	filt := expression.Name(key).Equal(expression.Value(field[key]))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	params := &dynamodb.ScanInput{
		TableName:                 aws.String(tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	result, err := config.DynamoDB.Scan(params)
	if err != nil {
		return nil, err
	}

	var asteroids []models.Asteroid

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &asteroids)
	if err != nil {
		return nil, err
	}
	return &asteroids, nil
}

func deleteAsteroid(id string, tableName string) (statusCode int, retErr error) {
	key := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(id),
		},
	}

	input := &dynamodb.DeleteItemInput{
		Key:       key,
		TableName: aws.String(tableName),
	}

	asteroid, err := getAsteroidByField(map[string]string{"id": id}, "Asteroids")
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if len(*asteroid) != 1 {
		return http.StatusNotFound, errors.New("asteroid not found")
	}

	_, err = config.DynamoDB.DeleteItem(input)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func fetchAllAsteroids(tableName string) ([]models.Asteroid, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := config.DynamoDB.Scan(input)
	if err != nil {
		return nil, fmt.Errorf("got error performing scan: %v", err)
	}

	var asteroids []models.Asteroid

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &asteroids)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal results: %v", err)
	}

	return asteroids, nil
}

func mergeAsteroid(existing models.Asteroid, update models.AsteroidsPatchDTO) models.Asteroid {
	v := reflect.ValueOf(update)
	t := v.Type()
	result := existing

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if !field.IsZero() {
			reflect.ValueOf(&result).Elem().FieldByName(fieldType.Name).Set(field)
		}
	}

	return result
}

func ErrorResponse(c echo.Context, statusCode int, err string) error {
	return c.JSON(statusCode, responses.BasicResponse{
		Status:  statusCode,
		Message: "error",
		Data:    &echo.Map{"data": err},
	})
}
