package models

type CategoryDescription struct {
	Name            string `json:"category_name"               dynamodbav:"category_name,omitempty"`

	Description     string `json:"description"        dynamodbav:"description,omitempty"`
	MetaDescription string `json:"meta_description"   dynamodbav:"meta_description,omitempty"`
	MetaKeyword     string `json:"meta_keyword"       dynamodbav:"meta_keyword,omitempty"`
	MetaTitle       string `json:"meta_title"         dynamodbav:"meta_title,omitempty"`
}