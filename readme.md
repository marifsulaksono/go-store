# About This Project
go-store is a simple marketplace RESTful API backend development side with references API Design on website [Shopee](https://shopee.co.id).

What tech stack do I use?
- Go Programming Language for common language
- MySQL for Database Management
- gorm.io/gorm for Object Relational Mapping (ORM)
- gorilla/mux for Route Management
- joho/godotenv for Env Management
- github.com/golang-jwt/jwt for Authorization
- golang.org/x/crypto for Generate and Hash Password <br />

# Installation
This is the instruction how you can test _go-store_ on your local computer:
## Prerequisites
- Make sure that you have installed MySQL on your computer.
  ```sh
  mysql --version
  ``` 
  If not installed, you can get it on [MySQL Documentaion](https://dev.mysql.com/doc/mysql-installation-excerpt/8.0/en/)
- Make sure that you have installed Go on your computer.
  ```sh
  go version
  ```
  If not installed and you have an IDE like Visual Studio Code, you can get it on [Go Documentaion](https://go.dev/doc/install)
- Make sure that you have installed Git on your computer.
  ```sh
  git --version
  ```
  If not installed and you have an IDE like Visual Studio Code, you can get it on [Git Documentaion](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) <br />
## Clone and Prepare Project
- Clone this project.
  ```sh
  git clone https://github.com/marifsulaksono/go-store.git
  ```
- Create new env file on project directory. Write the following into the env file.
  Example:
  ```
  DB_USERNAME=root
  DB_PASSWORD=root
  DB_HOST=127.0.0.1:3306
  DB_NAME=go-store?parseTime=true
  SERVER_PORT=8080
  ```
  Adjust to your environmental values. Then, **save file with name ```.env```** <br />
## Run Project
- Make sure your MySQL service is running. You can confirm this with a command:<br />
```Linux:```
  ```sh
  sudo systemctl status mysql
  ```
  ```Windows:```
  ```sh
  mysqladmin -u root -p status
  ```
- Run Project on project directory path.
  ```sh
  go run .
  ```
  After the output write "server starting...." response, you can use this project. Good Luck! <br />

# Project Usage
> [!NOTE]
> Some endpoint is need authorization to be accessed. You can get the token JWT by login, don't forget to register new account first if no user exist on database.
## User Route
### GET
- ```/users/{id}```, Get user data by id user.
### POST
- ```/register```, Register new user. Example body:
```json
{
	"name": "Muhammad Arif Sulaksono",
	"username": "arif",
	"password": "arif",
	"email": "arif@gmail.com",
	"phonenumber": 81234567890
}
```
> [!NOTE]
> Username only allowed alphabetic and numeric. Minimum length is 8 and maximum length is 50.<br />
> Password must have one numeric, lowercase, uppercase, and special character. Minimum length is 8 and maximum length is 30.
- ```/login```, Login user with _basic auth_ method.
### PUT
- ```/users/profile```, Update logined user account. Example body:
```json
{
	"name": "Muhammad Arif Sulaksono",
	"username": "arif",
	"email": "arif@gmail.com",
	"phonenumber": 81234567890
}
```
### PATCH
- ```/users/password```, Change password logined user . Example body:
```json
{
	"old_password": "arif",
  "new_password": "arif123"
}
```
> [!NOTE]
> Password must have one numeric, lowercase, uppercase, and special character. Minimum length is 8 and maximum length is 30.
### DELETE
- ```/users```, Delete logined user account. <br />

## SHIPPING ADDRESS
### GET
- ```/users/address```, Get all logined user's shipping address.
- ```/users/address/{id}```, Get logined user's shipping address by id.
### POST
- ```/users/address```, Create new shipping address, by default user id is logined user. Example body:
```json
{
	"recepient_name": "Arif",
	"Address": "Probolinggo",
	"phonenumber": "81234567890"
}
```
### PUT
- ```/users/address/{id}```, Update shipping address by id. Example body:
```json
{
	"recepient_name": "Arif",
	"Address": "Kraksaan",
	"phonenumber": "81987654320"
}
```
### DELETE
- ```/users/address/{id}```, Delete shipping address by id. <br />

## CATEGORY
### GET
- ```/categories```, Get all categories.
- ```/categories/{id}```, Get Category By id.
### POST
- ```/categories```, Create new category. Example body:
```json
{
  "name": "Computer & Laptop"
}
```
### PUT
- ```/categories/{id}```, Update category by id. Example body:
```json
{
  "name": "Computer & Laptop"
}
```
### DELETE
- ```/categories/{id}```, Delete category by id. <br />

## STORE
### GET
- ```/stores```, Get all stores.
- ```/stores/{id}```, Get store by id.
### POST
- ```/stores```, Create new store. Example body:
```json
{
	"name_store": "Arif Comp",
	"address": "Kraksaan, Kabupaten Probolinggo",
	"email": "arif.comp@gmail.com",
	"desc": "Kami menyediakan berbagai macam kartu perdana"
}
```
> [!NOTE]
> 1 user only has allowed to create 1 store.
### PUT
- ```/stores/{id}```, Update store by id. Example body:
```json
{
	"name_store": "Arif Comp",
	"address": "Kraksaan, Kabupaten Probolinggo",
	"email": "arif.comp@gmail.com",
	"desc": "Kami menyediakan berbagai macam kartu perdana"
}
```
### DELETE
- ```/stores/{id}```, Delete store by id. <br />

## PRODUCT
> [!NOTE]
> Seller side is on development, so this project is on buyer side. If you want to insert new product, please create the store and insert the store id manually.
### GET
- ```/products/search```, Get all product. at this endpoint, you can use some parameter for filtering and sorting.
  | Parameter | Description | Example |
  | --- | --- | --- |
  | Keyword | Filtering product by name of product | /products/search?keyword=laptop |
  | Status | Filtering product by status (sale or soldout) | /products/search?status=sale |
  | Sort By | Sorting product by name, stock, price, and sold product. By default sort by id | /products/search?sortBy=DESC |
  | Order | Ordering product by ascending or descanding. By default order is ascending | /products/search?order=DESC |
  | Minimum Price | Set the minimum price of product search | /products/search?minPrice=10000 |
  | Maximum Price | Set the maximum price of product search | /products/search?minPrice=10000000 |
  | Limit | Set limit of pagination. By default limit is 25 | /products/search?limit=10 |
  | Page | Set page of pagination. By defailt is 1 | /products/search?page=2 |
  | Category ID | Filtering product by Category ID | /products/search?categoryId=1 |
  | Store ID | Filtering product by Store ID | /products/search?storeId=1 |
- ```/products/{id}```, Get product by id.
### POST
- ```/products```, Create new product. Example body:
```json
{
	"name": "Lenovo G40-70",
	"stock": 50,
	"price": 7000000,
	"desc": "Spesifikasi: Intel i3 RAM 4GB HDD 500GB",
	"category_id": 1,
	"store_id": 1
}
```
### PUT
- ```/products/{id}```, Update product by id. Example body:
```json
{
	"name": "Lenovo G40-70",
	"stock": 50,
	"price": 7000000,
	"desc": "Spesifikasi: Intel i3 RAM 4GB HDD 500GB",
	"category_id": 1
}
```
- ```/products/{id}/restore```, Restore deleted product by id.
### DELETE
- ```/products/{id}/delete```, Soft delete product by id.
- ```/products/delete/{id}```, Hard delete product by id. <br />
## CART
### GET
- ```/carts```, Get all cart by logined user id.
- ```/carts/{id}```, Get cart by cart id.
### POST
- ```/carts```, Create new cart. Example body:
```json
{
	"product_id": 1,
	"qty": 10
}
```
> [!NOTE]
> If there same product id on cart user, it will suplement the quantity, else create new cart.
### PUT
- ```/carts/{id}```, Update cart by id. Example body:
```json
{
	"qty": 10
}
```
### DELETE
- ```/carts/{id}```, Delete cart by id. <br />
## TRANSACTION
### GET
- ```/transactions```, Get all logined user transaction.
- ```/transactions/{id}```, Get transaction by id.
### POST
- ```/transactions```, Create new transaction. Example body:
```json
{
    "shipping_address_id": 1,
    "items": [
        {
            "product_id": 1,
            "qty": 10
        },
        {
            "product_id": 2,
            "qty": 10
        }
    ]
}
```
<br />

# Contact Me
- [LinkedIn](https://www.linkedin.com/in/marifsulaksono/)
- [Instagram](https://www.instagram.com/marfs.2102)