package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/fmo/tm-players/internal/application/core/domain"
)

type Adapter struct {
	Connection *dynamodb.DynamoDB
	TableName  string
}

func NewAdapter(tableName string) (*Adapter, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &Adapter{
		Connection: dynamodb.New(sess),
		TableName:  tableName,
	}, nil
}

func (a Adapter) Save(ctx context.Context, player *domain.Player) error {
	playerParsed, err := dynamodbattribute.MarshalMap(player)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      playerParsed,
		TableName: aws.String(a.TableName),
	}

	_, err = a.Connection.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
