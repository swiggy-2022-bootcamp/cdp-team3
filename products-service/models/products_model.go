package models

type Product struct {
	Id                 string                  `dynamodbav:"productId" json:"id,omitempty"`
	Model              string                  `dynamodbav:"model" json:"model,omitempty"`
	Quantity           int32                   `dynamodbav:"quantity" json:"quantity,omitempty"`
	MaxQuantity        int32                   `dynamodbav:"max_quantity" json:"max_quantity,omitempty"`
	CanPlaceOrder      bool                    `dynamodbav:"can_place_order" json:"can_place_order,omitempty"`
	Price              float64                 `dynamodbav:"price" json:"price,omitempty"`
	ManufacturerId     int32                   `dynamodbav:"manufacturer_id" json:"manufacturer_id,omitempty"`
	Sku                string                  `dynamodbav:"sku" json:"sku,omitempty"`
	Points             int32                   `dynamodbav:"points" json:"points,omitempty"`
	Rewards            int32                   `dynamodbav:"rewards" json:"rewards,omitempty"`
	Image              string                  `dynamodbav:"image" json:"image,omitempty"`
	Weight             float64                 `dynamodbav:"weight" json:"weight,omitempty"`
	Length             float64                 `dynamodbav:"len" json:"length,omitempty"`
	Width              float64                 `dynamodbav:"width" json:"width,omitempty"`
	Height             float64                 `dynamodbav:"height" json:"height,omitempty"`
	Minimun            int32                   `dynamodbav:"minimum" json:"minimum,omitempty"`
	ProductCategory    []int32                 `dynamodbav:"product_category" json:"product_category,omitempty"`
	ProductRelated     []int32                 `dynamodbav:"product_related" json:"product_related,omitempty"`
	ProductSeoUrl      []ProductSeoUrlModel    `dynamodbav:"product_seo_url" json:"product_seo_url,omitempty"`
	ProductDescription ProductDescriptionModel `dynamodbav:"product_description" json:"product_description,omitempty"`
}

type ProductSeoUrlModel struct {
	Keyword    string `dynamodbav:"keyword" json:"keyword,omitempty"`
	LanguageId int32  `dynamodbav:"language_id" json:"language_id,omitempty"`
	StoreId    int32  `dynamodbav:"store_id" json:"store_id,omitempty"`
}

type ProductDescriptionModel struct {
	LanguageId      int32  `dynamodbav:"language_id" json:"language_id,omitempty"`
	Name            string `dynamodbav:"name" json:"name,omitempty"`
	Description     string `dynamodbav:"description" json:"description,omitempty"`
	MetaTitle       string `dynamodbav:"meta_title" json:"meta_title,omitempty"`
	MetaDescription string `dynamodbav:"meta_description" json:"meta_description,omitempty"`
	MetaKeyword     string `dynamodbav:"meta_keyword" json:"meta_keyword,omitempty"`
	Tag             string `dynamodbav:"tag" json:"tag,omitempty"`
}
