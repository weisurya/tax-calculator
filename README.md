Project name: Tax Calculator
Programming language: GO language v1.10.3
Environment: Windows 10
IDE: Microsoft Visual Code

Database: AWS - Amazon DynamoDB - PostgreSQL 10.4 db.t2.micro
Table structure: Please see it at storage/create_table_items.sql

Several options to utilize the project:
1. go main.go
2. docker build . -t tax-calculator
3. docker-compose up

Utilize all unit test:
1. At root project directory, type "go test ./..."


Github: https://github.com/weisurya/tax-calculator
API Documentation: https://www.getpostman.com/collections/4ff25583a0c9a4954cd2

Concern:
- Missing "value" parameter which could be used to calculate the tax amount. Instead, I use amount as the "value" to calculate the tax amount