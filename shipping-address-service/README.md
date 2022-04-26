# Shipping-Address-Service

Contains the Shipping Address Microservice for the final Swiggy I++ E-Commerce Application project created to showcase the skills learnt throughout the programme.

## Project Structure

| ![Folder Structure](shippingfolder.png) |
| --- |

## Relevant REST APIs (:3003)

| HTTP Verb  | Endpoint Path                                          |  Description                                       |
| ---------- | ------------------------------------------------------ | -------------------------------------------------- |
| GET        | /                                                      | Health Check                                       |
| GET        | /swagger/\*                                            | Swagger UI                                         |
| POST       | /shipping-service/api/shippingaddress                  | Adds a new Shipping Address to DB.                 |
| GET        | /shipping-service/api/shippingaddress/shipping_id      | Gets Shipping Address By Shipping Id         |
| PUT        | /shipping-service/api/shippingaddress/shipping_id      | Update Shipping Address to DB.                     |
| DELETE     | /shipping-service/api/shippingaddress/shipping_id      | Delete Shipping Adddress to DB.                    |
| POST       | /shipping-service/api//shippingaddress/existing/:id    | Set Shipping Address to default Shipping Address.  |
| GET        | /shipping-service/api/shippingaddress/existing/:userId | Gets Default Shipping Address of User              |

| ![Shipping Address Service REST Swagger]() |
| ---------- |

## GRPC Connections

1. Shipping Service -> Auth Service (:8012): To verify token from Auth service.

2. Shipping Service (:8003) <- Admin Service : Admin service sends the shipping address object as response, the shipping service will store the address in shipping DB and return shipping ID to the admin service to store in Admin service DB.

3. Shipping Service (:8003)<- Customer Service : Customer service sends the shipping address object as response, the shipping service will store the address in shipping DB and return shipping ID to the customer service to store in Customer service DB.

4. Shipping Service (:8003)<- Checkout Service : Shipping service give checkout service the default shipping address of the customer.

## Sonarqube Quality Check

| ![Sonarqube Quality Dashboard](sonar.png) |
| ---------- |


