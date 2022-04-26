package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/swiggy-ipp/cart-service/dto/requests"
	repositories_mock "github.com/swiggy-ipp/cart-service/mocks/repositories"
	"github.com/swiggy-ipp/cart-service/models"
)

const mockUserID string = "mockUserID"
const mockProductID string = "mockProductID"
const mockProductID2 string = "mockProductID2"

var mockCart models.Cart = models.Cart{
	UserID: mockUserID,
	Items: []models.CartItem{
		{
			ProductID: mockProductID,
			Quantity:  1,
		},
	},
}
var mockCartItemRequest requests.CartItemRequest = requests.CartItemRequest{
	ProductID: mockProductID2,
	Quantity:  1,
}

func TestNewCartService(t *testing.T) {
	// Act
	got := NewCartService(nil)

	// Assert: Check if got is a pointer to a CartService
	if reflect.TypeOf(got).Kind() != reflect.Ptr {
		t.Errorf("Expected a pointer to a CartService, got %s", reflect.TypeOf(got).Kind())
	}
}

func TestCreateCartExisting(t *testing.T) {
	// Mock Setup
	ctrl := gomock.NewController(t)
	repo := repositories_mock.NewMockCartRepository(ctrl)
	defer ctrl.Finish()

	// Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Arrange
	repo.EXPECT().
		ReadByUserID(gomock.Any(), mockUserID).
		Return(&mockCart, nil)
	cs := NewCartService(repo)

	// Assert
	if err := cs.CreateCart(ctx, mockUserID); err != nil {
		t.Errorf("CreateCart() error = %v", err)
	}
}

func TestCreateCartNotExisting(t *testing.T) {
	// Mock Setup
	ctrl := gomock.NewController(t)
	repo := repositories_mock.NewMockCartRepository(ctrl)
	defer ctrl.Finish()

	// Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Arrange
	repo.EXPECT().
		ReadByUserID(gomock.Any(), mockUserID).
		Return(nil, nil)
	repo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, nil)
	cs := NewCartService(repo)

	// Assert
	if err := cs.CreateCart(ctx, mockUserID); err != nil {
		t.Errorf("CreateCart() error = %v", err)
	}
}

func TestCreateCartItem(t *testing.T) {
	// Mock Setup
	ctrl := gomock.NewController(t)
	repo := repositories_mock.NewMockCartRepository(ctrl)
	defer ctrl.Finish()

	// Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Arrange
	repo.EXPECT().
		ReadByUserID(ctx, mockUserID).
		Return(&mockCart, nil)
	repo.EXPECT().UpdateCartItems(ctx, gomock.Any()).Return(nil)
	cs := NewCartService(repo)

	// Assert
	if err := cs.CreateCartItem(ctx, mockCartItemRequest, mockUserID); err != nil {
		t.Errorf("CreateCartItem() error = %v", err)
	}
}

func TestGetCartItems(t *testing.T) {
	// Mock Setup
	ctrl := gomock.NewController(t)
	repo := repositories_mock.NewMockCartRepository(ctrl)
	defer ctrl.Finish()

	// Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Arrange
	repo.EXPECT().
		ReadByUserID(ctx, mockUserID).
		Return(&mockCart, nil)
	repo.EXPECT().
		Read(ctx, mockUserID).
		Return(nil, nil)
	cs := NewCartService(repo)

	// Assert
	if out, err := cs.GetCartItems(ctx, "", mockUserID); (err != nil) && (out != nil) {
		t.Errorf("GetCartItems() error = %v", err)
	}
	if out, err := cs.GetCartItems(ctx, mockUserID, ""); (err != nil) && (out != nil) {
		t.Errorf("GetCartItems() error = %v", err)
	}
}

func TestUpdateCartItem(t *testing.T) {
	// Mock Setup
	ctrl := gomock.NewController(t)
	repo := repositories_mock.NewMockCartRepository(ctrl)
	defer ctrl.Finish()

	// Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Arrange
	repo.EXPECT().
		ReadByUserID(ctx, mockUserID).
		Return(&mockCart, nil)
	repo.EXPECT().
		UpdateCartItems(ctx, gomock.Any()).Return(nil)
	cs := NewCartService(repo)

	// Assert
	if err := cs.UpdateCartItem(ctx, mockCartItemRequest, mockUserID); err != nil {
		t.Errorf("UpdateCartItem() error = %v", err)
	}
}

func TestDeleteCartItem(t *testing.T) {
	// Mock Setup
	ctrl := gomock.NewController(t)
	repo := repositories_mock.NewMockCartRepository(ctrl)
	defer ctrl.Finish()

	// Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Arrange
	repo.EXPECT().
		ReadByUserID(ctx, mockUserID).
		Return(&mockCart, nil)
	repo.EXPECT().
		UpdateCartItems(ctx, gomock.Any()).Return(nil)
	cs := NewCartService(repo)

	// Assert
	if err := cs.DeleteCartItem(ctx, mockProductID, mockUserID); err != nil {
		t.Errorf("DeleteCartItem() error = %v", err)
	}
}

func TestEmptyCart(t *testing.T) {
	// Mock Setup
	ctrl := gomock.NewController(t)
	repo := repositories_mock.NewMockCartRepository(ctrl)
	defer ctrl.Finish()

	// Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Arrange
	repo.EXPECT().
		ReadByUserID(ctx, mockUserID).
		Return(&mockCart, nil)
	repo.EXPECT().
		UpdateCartItems(ctx, gomock.Any()).Return(nil)
	cs := NewCartService(repo)

	// Assert
	if err := cs.EmptyCart(ctx, requests.CartIDRequest{UserID: mockUserID}); err != nil {
		t.Errorf("EmptyCart() error = %v", err)
	}
}

func TestDeleteCart(t *testing.T) {
	// Mock Setup
	ctrl := gomock.NewController(t)
	repo := repositories_mock.NewMockCartRepository(ctrl)
	defer ctrl.Finish()

	// Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Arrange
	repo.EXPECT().
		ReadByUserID(ctx, mockUserID).
		Return(&mockCart, nil)
	repo.EXPECT().
		Delete(ctx, gomock.Any()).
		Return(nil)
	cs := NewCartService(repo)

	// Assert
	if err := cs.DeleteCart(ctx, requests.CartIDRequest{UserID: mockUserID}); err != nil {
		t.Errorf("DeleteCart() error = %v", err)
	}
}
