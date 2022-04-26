package repository

import (
	"fmt"
	"time"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
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
func toPersistedDynamodbEntitySA(o *models.Category) *models.Category {
	return &models.Category{

		CategoryId :        uuid.New().String(),
		CategoryDescription: o.CategoryDescription,
		
	}
}
func (p CategoryRepositoryImpl) AddCategoryToDB(category *models.Category) *apperrors.AppError {
	fmt.Println("Inside category repo")
	fmt.Println("category",category)
	categoryRecord := toPersistedDynamodbEntitySA(category)
	fmt.Println(categoryRecord)

	data, err := dynamodbattribute.MarshalMap(categoryRecord)
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
func (categoryRepo CategoryRepositoryImpl) DeleteCategoryByIDFromDB(categoryId string) (bool, *apperrors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//check if category is associated with products
	// var queryInput = &dynamodb.QueryInput{
	// 	TableName: aws.String("ProductCategoryRelation"),
	// 	KeyConditions: map[string]*dynamodb.Condition{
	// 		"category_id": {
	// 			ComparisonOperator: aws.String("EQ"),
	// 			AttributeValueList: []*dynamodb.AttributeValue{
	// 				{
	// 					S: aws.String(categoryId),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// var resp, err = categoryRepo.categoryDB.Query(queryInput)
	// if err != nil {
	// 	return false,apperrors.NewUnexpectedError(err.Error())
	// }
	// if resp != nil {
	// 	return false, apperrors.NewUnexpectedError(err.Error())
	// }
	//delete the category
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"category_id": {
				S: aws.String(categoryId),
			},
		},
		TableName: aws.String(categoryCollection),
	}

	_, err := categoryRepo.categoryDB.DeleteItemWithContext(ctx, input)
	if err != nil {
		return false, apperrors.NewUnexpectedError(err.Error())
	}
	return true, nil
}
func (categoryRepo CategoryRepositoryImpl) DeleteCategoriesFromDB(categoryIds []string) (bool,*apperrors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// for _, categoryId := range categoryIds {
	// 	var queryInput = &dynamodb.QueryInput{
	// 		TableName: aws.String("ProductCategoryRelation"),
	// 		KeyConditions: map[string]*dynamodb.Condition{
	// 			"category_id": {
	// 				ComparisonOperator: aws.String("EQ"),
	// 				AttributeValueList: []*dynamodb.AttributeValue{
	// 					{
	// 						S: aws.String(categoryId),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	}
	// 	var resp, err = categoryRepo.categoryDB.Query(queryInput)
	// 	if err != nil {
	// 		return false,apperrors.NewUnexpectedError(err.Error())
	// 	}
	// 	if resp != nil {
	// 		return false, apperrors.NewUnexpectedError(err.Error())
	// 	}
		//delete the category
		for _,categoryId := range categoryIds{
		input := &dynamodb.DeleteItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"category_id": {
					S: aws.String(categoryId),
				},
			},
			TableName: aws.String(categoryCollection),
		}

		_, err := categoryRepo.categoryDB.DeleteItemWithContext(ctx, input)
		if err != nil {
			return false, apperrors.NewUnexpectedError(err.Error())
		}
	}
	//}
	return true, nil
}

func (categoryRepo CategoryRepositoryImpl) UpdateCategoryByIdFromDB(categoryId string,category *models.Category) (bool,*apperrors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	prevCategoryInput := &dynamodb.GetItemInput{
		TableName: aws.String("Category"),
		Key: map[string]*dynamodb.AttributeValue{
			"category_id": {
				S: aws.String(category.CategoryId),
			},
		},
	}
	prevCategoryResult, err := categoryRepo.categoryDB.GetItemWithContext(ctx, prevCategoryInput)
	if err != nil {
		return false, apperrors.NewUnexpectedError(err.Error())
	}
	oldCategory := &models.Category{}
	err = dynamodbattribute.UnmarshalMap(prevCategoryResult.Item, &oldCategory)

	if err != nil {
		return false, apperrors.NewUnexpectedError(err.Error())
	}

	if category.CategoryDescription[0].Description == "" {
		category.CategoryDescription[0].Description = oldCategory.CategoryDescription[0].Description
	}
	if category.CategoryDescription[0].Name == "" {
		category.CategoryDescription[0].Name = oldCategory.CategoryDescription[0].Name
	}
	if category.CategoryDescription[0].MetaDescription == "" {
		category.CategoryDescription[0].MetaDescription = oldCategory.CategoryDescription[0].MetaDescription
	}
	if category.CategoryDescription[0].MetaKeyword == "" {
		category.CategoryDescription[0].MetaKeyword = oldCategory.CategoryDescription[0].MetaKeyword
	}
	if category.CategoryDescription[0].MetaTitle == "" {
		category.CategoryDescription[0].MetaTitle = oldCategory.CategoryDescription[0].MetaTitle
	}
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
				S: aws.String(category.CategoryDescription[0].MetaTitle),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(category.CategoryId),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set category_description.name =:s, category_description.description = :s1, category_description.mata_description = :s2,  category_description.mata_keyword= :s3, category_description.meta_title = :s4"),
		TableName:        aws.String("Category"),
	}
	_, err = categoryRepo.categoryDB.UpdateItemWithContext(ctx, input)
	if err != nil {
		return false, apperrors.NewUnexpectedError(err.Error())
	}
	return true, apperrors.NewUnexpectedError(err.Error())
}
