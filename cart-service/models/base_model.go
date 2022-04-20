package models

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Base Data Model Interface
type dataModel interface {
	GenerateKey() string
}

// Serializable Data Model Interface
type dynamoSerializableDataModel interface {
	Marshal() (map[string]types.AttributeValue, error)
	UnMarshal() (interface{}, error)
}

//  Encoder for DynamoDB that uses JSON Tag key
func encoderJSON(eo *attributevalue.EncoderOptions) {
	eo.TagKey = "json"
}

// Decoder for DynamoDB that uses JSON Tag key
func decoderJSON(do *attributevalue.DecoderOptions) {
	do.TagKey = "json"
}
