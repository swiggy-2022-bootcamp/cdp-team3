package repository

import (
	"fmt"
	"time"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	//"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	apperrors "github.com/cdp-team3/categories-service/app-errors"
	"github.com/cdp-team3/categories-service/domain/models"
	"github.com/cdp-team3/categories-service/utils/logger"
	
)

const categoryCollection = "Categories"

type CategoryRepositoryImpl struct {
	categoryDB *dynamodb.DynamoDB
	
}

func NewCategoryRepositoryImpl(categoryDB *dynamodb.DynamoDB) CategoryRepository {
	return &CategoryRepositoryImpl{
		categoryDB: categoryDB,
		
	}
}

func (p CategoryRepositoryImpl) DBHealthCheck() bool {

	_, err := p.categoryDB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		logger.Error("Database connection is down.")
		return false
	}
	return true
}
// type CategoryDesciption struct {
// 	Name            string `json:"name"               bson:"name,omitempty"`
// 	Description     string `json:"description"        bson:"description,omitempty"`
// 	MetaDescription string `json:"meta_description"   bson:"meta_description,omitempty"`
// 	MetaKeyword     string `json:"meta_keyword"       bson:"meta_keyword,omitempty"`
// 	MetaTitle       string `json:"meta_title"         bson:"meta_title,omitempty"`
// }
// type Category struct {
// 	CategoryId             string                    `json:"category_id" dynamodbav:"category_id" validate:"required"`
// 	CategoryDesciption     []CategoryDesciption                    `json:"category_description" dynamodbav:"category_description"`
// }
// cd := CategoryDesciption{Name: "testname", Description:"testdesc" ,MetaDescription:"testmetadesc",MetaKeyword:"testmetakey",MetaTitle:"testmetatitle"}
// c:=Category{CategoryDesciption:cd}
func (p CategoryRepositoryImpl) AddCategoryToDB(category *models.Category) *apperrors.AppError {
	fmt.Println("Inside category repo")
	fmt.Println("category",category)

	data, err := dynamodbattribute.MarshalMap(category)
	if err != nil {
		logger.Error("Marshalling of category failed - " + err.Error())
		return apperrors.NewUnexpectedError(err.Error())
	}

	query := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(categoryCollection),
	}

	result, err := p.categoryDB.PutItem(query)
	if err != nil {
		fmt.Println(err)
		logger.Error("Failed to insert category into database - " + err.Error())
		return apperrors.NewUnexpectedError(err.Error())
	}
	fmt.Println(result)
	return nil
}
func (p CategoryRepositoryImpl) FindAllCategoryFromDB() ([]models.Category, *apperrors.AppError) {

	// create the api params
	params := &dynamodb.ScanInput{
		TableName: aws.String(categoryCollection),
	}

	var categoryList []models.Category

	// scan and filter for the items
	err := p.categoryDB.ScanPages(params, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		// Unmarshal the slice of dynamodb attribute values into a slice of custom structs
		var categories []models.Category
		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &categories)
		if err != nil {
			fmt.Printf("\nCould not unmarshal AWS data: err = %v\n", err)
			return true
		}

		categoryList = append(categoryList,categories...)

		return true
	})

	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		logger.Error(err.Error())
		return nil, apperrors.NewUnexpectedError(err.Error())
	}

	return categoryList, nil
}

func (p CategoryRepositoryImpl) GetCategoryFromDB(category_id string) (*models.Category, *apperrors.AppError) {
	category := &models.Category{}

	query := &dynamodb.GetItemInput{
		TableName: aws.String(categoryCollection),
		Key: map[string]*dynamodb.AttributeValue{
			"category_id": {
				S: aws.String(category_id),
			},
		},
	}

	result, err := p.categoryDB.GetItem(query)
	if err != nil {
		logger.Info(result)
		logger.Error("Failed to get item from database - " + err.Error())
		return nil ,  apperrors.NewUnexpectedError(err.Error())
	}

	if result.Item == nil {
		logger.Error("Categories for given ID doesn't exists - ")
		err_ := apperrors.NewNotFoundError("Categories for given ID doesn't exists")
		return nil, err_
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, category)
	if err != nil {
		logger.Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, apperrors.NewUnexpectedError(err.Error())
	}
	return category, nil
}
func (p CategoryRepositoryImpl) DeleteCategoryByIDFromDB(category_id string) *apperrors.AppError{
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String(categoryCollection),
		Key: map[string]*dynamodb.AttributeValue{
			"category_id": {
				S: aws.String(category_id),
			},
		},
	}

	resp, err := p.categoryDB.DeleteItem(params)
	if err != nil {
		logger.Error(err.Error())
		return apperrors.NewUnexpectedError(err.Error())
	} else {
		logger.Info("Success")
		logger.Info(resp)
		return nil
	}
}
// func (p CategoryRepositoryImpl) UpdateCategoryByIDFromDB(category_id string) (*models.Category, *apperrors.AppError) {
// 	category := &models.Category{}

// 	query := &dynamodb.GetItemInput{
// 		TableName: aws.String(categoryCollection),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"category_id": {
// 				S: aws.String(category_id),
// 			},
// 		},
// 	}

// 	result, err := p.categoryDB.GetItem(query)
// 	if err != nil {
// 		logger.Info(result)
// 		logger.Error("Failed to get item from database - " + err.Error())
// 		return nil ,  apperrors.NewUnexpectedError(err.Error())
// 	}

// 	if result.Item == nil {
// 		logger.Error("Categories for given ID doesn't exists - ")
// 		err_ := apperrors.NewNotFoundError("Categories for given ID doesn't exists")
// 		return nil, err_
// 	}

// 	err = dynamodbattribute.UnmarshalMap(result.Item, category)
// 	if err != nil {
// 		logger.Error("Failed to unmarshal document fetched from DB - " + err.Error())
// 		return nil, apperrors.NewUnexpectedError(err.Error())
// 	}
// 	return category, nil
// }
// func (p  CategoryRepositoryImpl) UpdateCategoryByIDFromDB(category_id string, category *models.Category) ( *models.Category,*apperrors.AppError) {
// 	toUpd, err := getCategoryUpdExp(category)

// 	if err != nil {
// 		logger.Error("error while creating expression %s", err)
// 		return nil, apperrors.NewUnexpectedError(err.Error())
// 	}
// 	updItemIn := dynamodb.UpdateItemInput{
// 		TableName:                aws.String(categoryCollection),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"category_id": {
// 				S: aws.String(category_id),
// 			},
// 		},
// 		ExpressionAttributeNames:  toUpd.Names(),
// 		ExpressionAttributeValues: toUpd.Values(),
// 		UpdateExpression:          toUpd.Update(),
// 		ReturnValues:              aws.String("ALL_NEW"),
// 		ConditionExpression:       toUpd.Condition(),
// 	}

// 	fmt.Printf("update item input : %+v", updItemIn)

// 	resp, err :=p.categoryDB.UpdateItem(&updItemIn)

// 	if err != nil {
// 		fmt.Printf("error while updating cateogories %s", err.Error())
// 		return nil,apperrors.NewUnexpectedError(err.Error())
// 	}

// 	fmt.Printf("category update resp %v", resp)

// 	updatedAttributes := models.Category{}
// 	if err = dynamodbattribute.UnmarshalMap(resp.Attributes, &updatedAttributes); err != nil {
// 		fmt.Printf("error while updating cateogories %s", err.Error())
// 		return nil, apperrors.NewUnexpectedError(err.Error())
// 	}

// 	return &updatedAttributes, nil
// }



// func getCategoryUpdExp(cat *models.Category) (expression.Expression, error) {
// 	var updateExp expression.UpdateBuilder

// 	desc := cat.CategoryDesciption[0]
// 	fmt.Println("\n desc",desc)
// 	// descPrefix := "category_description."
// 	//category description
// 	fmt.Println("desc.Name",desc.Name)
// 	if desc.Name != "" {
// 		updateExp = updateExp.Set(expression.Name("category_description.name"), expression.Value(desc.Name))
// 	}

// 	if desc.Description != "" {
// 		updateExp = updateExp.Set(expression.Name("category_description.description"), expression.Value(desc.Description))
// 	}

// 	if desc.MetaDescription != "" {
// 		updateExp = updateExp.Set(expression.Name("category_description.meta_description"), expression.Value(desc.MetaDescription))
// 	}

// 	if desc.MetaKeyword != "" {
// 		updateExp = updateExp.Set(expression.Name("category_description.meta_keyword"), expression.Value(desc.MetaKeyword))
// 	}

// 	if desc.MetaTitle != "" {
// 		updateExp = updateExp.Set(expression.Name("category_description.meta_title"), expression.Value(desc.MetaTitle))
// 	}

// 	fmt.Printf("category update expression %+v", updateExp)
// 	exp, err := expression.
// 		NewBuilder().
// 		WithCondition(expression.AttributeExists(expression.Name("category_id"))).
// 		WithUpdate(updateExp).
// 		Build()

// 	fmt.Printf("%+v", exp.Names())
// 	fmt.Println(exp.Values())
// 	fmt.Println(*exp.Update())
// 	return exp, err
// }

func (p  CategoryRepositoryImpl) UpdateCategoryByIDFromDB1(category_id string, category *models.Category) (bool, *apperrors.AppError) {
	 ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	 defer cancel()
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				S: aws.String(category.CategoryDescription[0].Name),
			}, ":s1": {
				S: aws.String(category.CategoryDescription[0].Description),
			}, ":s2": {
				S: aws.String(category.CategoryDescription[0].MetaDescription),
			}, ":s3": {
				S: aws.String(category.CategoryDescription[0].MetaKeyword),
			}, ":s4": {
				S: aws.String(category.CategoryDescription[0].MetaTitle ),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"category_id": {
				S: aws.String(category_id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set category_description[0].category_name =:s, category_description[0].description = :s1, category_description[0].meta_description= :s2, category_description[0].meta_keyword = :s3, category_description[0].meta_title = :s4"),
		TableName:        aws.String(categoryCollection),
	}

	_, err := p.categoryDB.UpdateItemWithContext(ctx, input)
	if err != nil {
		return false, apperrors.NewUnexpectedError(err.Error())
	}
	return true, nil
}
