# POS Receipt System

A small business Point-of-Sale (POS) system for managing products, sales, and generating PDF receipts. Built with **Go**, **Gin**, and **Zap** for logging.

## Features

- User authentication (register/login) with JWT
- Product management: add, list, update, delete
- Sales management: create, list, update, delete
- PDF receipt generation for sales
- Business settings management
- JSON API with OpenAPI 3.0 documentation
- Debug logging with **Zap**

## Tech Stack

- **Backend**: Go, Gin
- **Database**: SQLite
- **Logging**: Uber Zap
- **API Documentation**: OpenAPI / Swagger

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/pos-receipt-system.git
cd pos-receipt-system
