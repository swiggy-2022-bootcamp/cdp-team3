package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id                 primitive.ObjectID      `json:"id,omitempty"`
	Model              string                  `json:"model,omitempty"`
	Quantity           int32                   `json:"quantity,omitempty"`
	MaxQuantity        int32                   `json:"max_quantity,omitempty"`
	CanPlaceOrder      bool                    `json:"can_place_order,omitempty"`
	Price              float64                 `json:"price,omitempty"`
	ManufacturerId     int32                   `json:"manufacturer_id,omitempty"`
	Sku                string                  `json:"sku,omitempty"`
	Points             int32                   `json:"points,omitempty"`
	Rewards            int32                   `json:"rewards,omitempty"`
	Image              string                  `json:"image,omitempty"`
	Weight             float64                 `json:"weight,omitempty"`
	Length             float64                 `json:"length,omitempty"`
	Width              float64                 `json:"width,omitempty"`
	Height             float64                 `json:"height,omitempty"`
	Minimun            int32                   `json:"minimum,omitempty"`
	ProductCategory    []int32                 `json:"product_category,omitempty"`
	ProductRelated     []int32                 `json:"product_related,omitempty"`
	ProductSeoUrl      []ProductSeoUrlModel    `json:"product_seo_url,omitempty"`
	ProductDescription ProductDescriptionModel `json:"product_description_model,omitempty"`
}

type ProductSeoUrlModel struct {
	Keyword    string `json:"keyword,omitempty"`
	LanguageId int32  `json:"language_id,omitempty"`
	StoreId    int32  `json:"store_id,omitempty"`
}

type ProductDescriptionModel struct {
	LanguageId      int32  `json:"language_id,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	MetaTitle       string `json:"meta_title,omitempty"`
	MetaDescription string `json:"meta_description,omitempty"`
	MetaKeyword     string `json:"meta_keyword,omitempty"`
	Tag             string `json:"tag,omitempty"`
}
