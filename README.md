# Receipt Processor Challenge

A web service that accepts receipts and calculates points based on predefined rules.

## Overview

This service processes receipts and awards points based on specific rules. It provides two main endpoints:
- Process receipts and generate unique IDs
- Retrieve points for processed receipts using their IDs

## Technical Stack

- Go (Golang)
- Gin Web Framework
- In-memory storage
- RESTful API design

## API Endpoints

### 1. Process Receipt
- **Endpoint**: POST `/receipt/process`
- **Content Type**: application/json
- **Request Body**:
```json
{
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
        {
            "shortDescription": "Mountain Dew",
            "price": "1.25"
        }
    ],
    "total": "1.25"
}
```
- **Success Response**: 
```json
{
    "id": "<unique-id>"
}
```

### 2. Get Points
- **Endpoint**: GET `/receipt/:id/points`
- **Success Response**:
```json
{
    "points": 32
}
```

## Points Calculation Rules

1. One point for every alphanumeric character in the retailer name
2. 50 points if the total is a round dollar amount with no cents
3. 25 points if the total is a multiple of 0.25
4. 5 points for every two items on the receipt
5. If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
6. 6 points if the day in the purchase date is odd
7. 10 points if the time of purchase is after 2:00pm and before 4:00pm

## Project Structure

```
fetch_receipt_processor_challenge/
├── main.go           # Main application entry point and route handlers
├── main_test.go      # Test cases for the main application
├── types/
│   └── types.go      # Data structures and types
└── helpers/
    └── helpers.go    # Helper functions for points calculation
```

## Running the Application

1. Clone the repository
2. Install dependencies:
```bash
go mod tidy
```
3. Run the application:
```bash
go run main.go
```
The server will start on `localhost:8080`

## Testing with Postman

1. Launch Postman
2. Import the following cURL commands:

Process Receipt:
```bash
curl -X POST \
  http://localhost:8080/receipts/process \
  -H 'Content-Type: application/json' \
  -d '{
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
        {
            "shortDescription": "Mountain Dew",
            "price": "1.25"
        }
    ],
    "total": "1.25"
}'
```

Get Points:
```bash
curl -X GET http://localhost:8080/receipts/{id}/points
```

3. Replace `{id}` in the Get Points request with the ID received from the Process Receipt response
4. Send the requests and verify the responses match the expected format

## Error Handling

The service includes comprehensive error handling for:
- Invalid JSON input
- Empty receipt data
- Invalid date/time formats
- Missing required fields
- Invalid price formats
