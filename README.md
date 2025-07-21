
🛒 E-Commerce Go App
=====================

A simple e-commerce backend built with Go and SQLite, featuring clean architecture and modular design.

------------------------------------------------------------

🚀 Getting Started
------------------

Follow the steps below to set up and run the project.

1. Clone the Repository
------------------------

    git clone https://github.com/vandannandwana/Basic-E-Commerce.git
    cd Basic-E-Commerce

2. Install Dependencies
------------------------

Ensure you have Go installed (v1.18 or higher), then run:

    go mod tidy

3. Configuration Setup
-----------------------

Create a config folder and a local.yaml file inside it:

    mkdir config

Create config/local.yaml with the following content:

    env: "dev"
    storage_path: "storage/storage.db"
    http_server:
      address: "localhost:8082"

4. Database Setup
------------------

Create the storage folder and initialize the SQLite database:

    mkdir storage
    sqlite3 storage/storage.db

(You can also allow your Go code to create the database schema automatically if implemented.)

✅ Run the Application
----------------------

Start the application using:

    go run ./cmd/e-commerce/main.go -config config/local.yaml

The server will start on http://localhost:8082 (as per local.yaml).

------------------------------------------------------------

📌 Notes
--------

- Make sure local.yaml and storage.db are listed in .gitignore to avoid pushing sensitive or environment-specific data.
- This project is a great foundation for learning or building real-world Go services.

📄 License
----------

MIT License — feel free to use and modify.

------------------------------------------------------------

📡 API Endpoints
----------------

Below are the available product-related API endpoints:

1. Create a New Product
------------------------
    POST /api/products

    Example (using curl):
    curl -X POST http://localhost:8082/api/products \
         -H "Content-Type: application/json" \
         -d '{"name":"Laptop","price":50000,"quantity":10}'

2. Get a Product by ID
-----------------------
    GET /api/products/{id}

    Example:
    curl http://localhost:8082/api/products/1

3. Get All Products
--------------------
    GET /api/products

    Example:
    curl http://localhost:8082/api/products

4. Delete a Product by ID
--------------------------
    DELETE /api/products/{id}

    Example:
    curl -X DELETE http://localhost:8082/api/products/1
