package repositories

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	log "github.com/sirupsen/logrus"
	"github.com/swiggy-ipp/cart-service/models"
)

// CRUD Repository implementation for Cart Collection in DynamoDB
type cartRepositoryImpl struct {
	cartDB *dynamodb.Client
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
	if cart.Id == "" {
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
		Item: data,
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
	out, err := cr.cartDB.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(cr.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: id,
			},
		},
	})
	if err != nil {
		log.Errorf("Failed to read cart: %v", err)
		return nil, err
	} else if out.Item == nil {
		log.Infof("Cart with ID %s not found", id)
		return nil, nil
	} else {
		// Model *dynamodb.GetItemOutput into cart
		var cart models.Cart
		err = cart.UnMarshal(out.Item)
		if err != nil {
			log.Errorf("Failed to deserialize cart: %v", err)
			return nil, err
		}
		return &cart, nil
	}
}

// ReadByUserID reads a cart by its associated userID
func (cr *cartRepositoryImpl) ReadByUserID(ctx context.Context, userID string) (*models.Cart, error) {
	// Get item from DB
	out, err := cr.cartDB.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(cr.tableName),
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberS{
				Value: userID,
			},
		},
	})
	if err != nil {
		log.Errorf("Failed to read cart by User ID: %v", err)
		return nil, err
	} else if out.Item == nil {
		log.Infof("Cart for user ID %s not found", userID)
		return nil, nil
	} else {
		// Model *dynamodb.GetItemOutput into cart
		var cart models.Cart
		err = cart.UnMarshal(out.Item)
		if err != nil {
			log.Errorf("Failed to deserialize cart: %v", err)
			return nil, err
		}
		return nil, nil
	}
}

// Update updates a cart by its ID
func (cr *cartRepositoryImpl) Update(ctx context.Context, cart *models.Cart) error {
	// Serialize item
	data, err := cart.Marshal()
	if err != nil {
		log.Error("Failed to serialize: ", err)
		return err
	}

	// Put item into DB
	_, err = cr.cartDB.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(cr.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: cart.Id,
			},
		},
		UpdateExpression: aws.String("SET #cart = :cart"),
		ExpressionAttributeNames: map[string]string{
			"#cart": "cart",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":cart": &types.AttributeValueMemberM{
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
	// Delete item from DB
	_, err := cr.cartDB.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(cr.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: id,
			},
		},
	})
	if err != nil {
		log.Errorf("Failed to delete cart: %v", err)
		return err
	}
	return nil
}

// EmptyCart fetches the cart identified by Cart ID and empties it
func (cr *cartRepositoryImpl) EmptyCart(ctx context.Context, id string) error {
	// Get cart from DB
	cart, err := cr.Read(ctx, id)
	if err != nil {
		log.Errorf("Failed to empty cart: %v", err)
		return err
	}

	// Empty Cart fetched from DB
	cart.Items = []models.CartItem{}
	err = cr.Update(ctx, cart)
	if err != nil {
		log.Errorf("Failed to empty cart: %v", err)
		return err
	}
	return nil
}
