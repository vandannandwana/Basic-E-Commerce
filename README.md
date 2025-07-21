# E-Commerce Go App

A simple e-commerce application built with Go, using SQLite for storage and a clean folder architecture.

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/vandannandwana/Basic-E-Commerce.git
cd Basic-E-Commerce

### 2. Download Required Dependencies

go mod tidy

### 3. Configuration Part

mkdir config

# config/local.yaml

env: "dev"
storage_path: "storage/storage.db"
http_server:
  address: "localhost:8082"

### 4. Database Setup

mkdir storage

sqlite3 storage/storage.db

## You are All Set to Go

go run .\cmd\e-commerce\main.go

