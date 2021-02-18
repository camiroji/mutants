package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Repository interface {
	SaveDNA(dna Dna) error
	GetStats() (Stats, error)
}

type DynamoRepo struct {
	DB *dynamodb.DynamoDB
}

type Dna struct {
	dna      *string
	isMutant *bool
}

type Stats struct {
	CountMutantsDna int
	CountTotalDnas  int
}

func (r DynamoRepo) SaveDNA(dna Dna) error {
	item := map[string]*dynamodb.AttributeValue{
		"dna":      {S: dna.dna},
		"isMutant": {BOOL: dna.isMutant},
	}
	input := &dynamodb.PutItemInput{Item: item, TableName: aws.String("dnas"), ReturnValues: aws.String("NONE")}
	_, err := r.DB.PutItem(input)
	return err
}

func (r DynamoRepo) GetStats() (Stats, error) {
	filt := expression.Name("isMutant").Equal(expression.Value(true))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		return Stats{}, err
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String("dnas"),
	}
	result, err := r.DB.Scan(params)
	stats := Stats{
		CountMutantsDna: int(*result.Count),
		CountTotalDnas:  int(*result.ScannedCount),
	}
	return stats, err
}
