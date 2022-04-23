package repositories

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	log "github.com/sirupsen/logrus"
	"github.com/swiggy-ipp/cart-service/models"
)

// CRUD Repository implementation for Cart Collection in DynamoDB
type cartRepositoryImpl struct {
	cartDB    *dynamodb.Client
	tableName string
}

// NewCartRepository returns a new CartRepository
func NewCartRepository(db *dynamodb.Client, tableName string) CartRepository {
	if tableName == "" {
		tableName = "cart"
	}
	return &cartRepositoryImpl{cartDB: db, tableName: tableName}
}

// Create creates a new cart
func (cr *cartRepositoryImpl) Create(ctx context.Context, cart *models.Cart) error {
	// Form key if not present
	if cart.ID == "" {
		cart.GenerateKey()
	}

	// Serialize item
	data, err := cart.Marshal()
	if err != nil {
		log.Error("Failed to serialize: ", err)
		return err
	}

	// Put item into DB
	_, err = cr.cartDB.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &cr.tableName,
		Item:      data,
	})
	if err != nil {
		log.Errorf("Failed to create cart: %v", err)
		return err
	}
	log.Info("Created cart: ", cart)
	return nil
}

// Read reads a cart by its ID
func (cr *cartRepositoryImpl) Read(ctx context.Context, id string) (*models.Cart, error) {
	// Get item from DB
	out, err := cr.cartDB.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(cr.tableName),
		KeyConditionExpression: aws.String("#id = :id"),
		ExpressionAttributeNames: map[string]string{
			"#id": "id",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberS{
				Value: id,
			},
		},
	})
	if err != nil {
		log.Errorf("Failed to read cart: %v", err)
		return nil, err
	} else if out.Items == nil || len(out.Items) == 0 {
		log.Infof("Cart with ID %s not found", id)
		return nil, errors.New("Cart not found")
	} else if len(out.Items) > 1 {
		log.Infof("Multiple carts with ID %s found", id)
		return nil, errors.New("Multiple carts with same ID found")
	} else {
		// Model *dynamodb.GetItemOutput into cart
		var cart models.Cart
		err = cart.UnMarshal(out.Items[0])
		if err != nil {
			log.Errorf("Failed to deserialize: %v", err)
			return nil, err
		}
		return &cart, nil
	}
}

// ReadByUserID reads a cart by its associated userID
func (cr *cartRepositoryImpl) ReadByUserID(ctx context.Context, userID string) (*models.Cart, error) {
	// Get item from DB
	out, err := cr.cartDB.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(cr.tableName),
		FilterExpression: aws.String("#user_id = :user_id"),
		ExpressionAttributeNames: map[string]string{
			"#user_id": "user_id",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":user_id": &types.AttributeValueMemberS{
				Value: userID,
			},
		},
	})

	if err != nil {
		log.Errorf("Failed to read cart by User ID: %v", err)
		return nil, err
	} else if out.Items == nil || len(out.Items) == 0 {
		log.Infof("Cart with User ID %s not found", userID)
		return nil, errors.New("Cart not found")
	} else if len(out.Items) > 1 {
		log.Infof("Multiple carts with User ID %s found", userID)
		return nil, errors.New("Multiple carts with same ID found")
	} else {
		// Model *dynamodb.GetItemOutput into cart
		var cart models.Cart
		err = cart.UnMarshal(out.Items[0])
		if err != nil {
			log.Errorf("Failed to deserialize cart: %v", err)
			return nil, err
		}
		return &cart, nil
	}
}

// Update updates a cart by its ID
func (cr *cartRepositoryImpl) UpdateCartItems(ctx context.Context, cart *models.Cart) error {
	// Serialize item
	data, err := models.MarshalGeneralList(cart.Items)
	if err != nil {
		log.Error("Failed to serialize: ", err)
		return err
	}

	// Get current item from DB
	out, err := cr.Read(ctx, cart.ID)
	if err != nil {
		log.Error("Failed to read cart: ", err)
		return err
	}

	// Put item into DB
	_, err = cr.cartDB.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(cr.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: out.ID,
			},
			"user_id": &types.AttributeValueMemberS{
				Value: out.UserID,
			},
		},
		UpdateExpression: aws.String("SET #items = :items"),
		ExpressionAttributeNames: map[string]string{
			"#items": "items",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":items": &types.AttributeValueMemberL{
				Value: data,
			},
		},
	})
	if err != nil {
		log.Errorf("Failed to update cart: %v", err)
		return err
	}
	return nil
}

// Delete deletes a cart by its ID
func (cr *cartRepositoryImpl) Delete(ctx context.Context, id string) error {
	// Get current item from DB
	out, err := cr.Read(ctx, id)
	if err != nil {
		return err
	}

	// Delete item from DB
	_, err = cr.cartDB.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(cr.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: out.ID,
			},
			"user_id": &types.AttributeValueMemberS{
				Value: out.UserID,
			},
		},
	})
	if err != nil {
		log.Errorf("Failed to delete cart: %v", err)
		return err
	}
	return nil
}
