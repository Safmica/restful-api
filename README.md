# Warehouse API

The Warehouse API is designed to manage and track a business that operates a warehouse. This API enables you to perform CRUD (Create, Read, Update, Delete) operations on various entities related to warehouse management, such as users, warehouses, products, categories, and stock. By using this API, you can effectively monitor and organize information related to the warehouse and products, as well as restrict access to certain endpoints to admins through the use of JWT (JSON Web Tokens).

## Features
- **User Management**: Create, read, update, and delete user information.
- **Warehouse Management**: Manage warehouse details including name, location, and capacity.
- **Product Management**: Handle product information, including linking products to categories and warehouses.
- **Category Management**: Manage product categories.
- **Stock Management**: Track and manage stock levels for products in the warehouse.
- **JWT Authentication**: Secure specific endpoints to ensure only authorized users can access or modify certain resources.

## Tech Stack
- ![VSCode](https://img.shields.io/badge/VSCode-0078D4?style=for-the-badge&logo=visual%20studio%20code&logoColor=white) **Visual Studio Code** - Used as the IDE for developing the Restful API.
- ![XAMPP](https://img.shields.io/badge/Xampp-F37623?style=for-the-badge&logo=xampp&logoColor=white) **XAMPP** - Used for running local servers and managing databases.
- ![MySQL](https://img.shields.io/badge/MySQL-005C84?style=for-the-badge&logo=mysql&logoColor=white) **MySQL** - Used as the database management system.
- ![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=Postman&logoColor=white) **Postman** - Used for testing and documenting the API endpoints.
- ![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) **Golang** - Used as the programming language for building the API.

## How to Run
This API can be run on your local development system using the following methods:

### Prerequisites
- Golang
- MySQL

### Environment Variables
Set up your environment variables as follows:
```bash
export PORT=<port>
export DBCONN="<username>:<password>@tcp(<hostname>:<port>)/<dbname>?charset=utf8&parseTime=True&loc=Local"
```

### Steps to Run
1. Clone the repository:
```bash
git clone https://github.com/username/warehouse-api.git
cd warehouse-api
```
2. Install dependencies:
```bash
go mod tidy
```
3. Run the application:
```bash
go run main.go
```
4. Use Postman to interact with the API endpoints.


## Documentation
You can access the documentation each endpoint from this API here:
> https://documenter.getpostman.com/view/36503501/2sA3kSnNSZ
