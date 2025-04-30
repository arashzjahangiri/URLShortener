# URLShortener

Go Version | App Version | Repo Size | Coverage
-----------|-------------|-----------|---------
1.24.2     | v0.1.0      | Compact   | Comprehensive

**URLShortener** is a modern, efficient, and scalable URL shortener built with Go. The name reflects its purpose: simplifying long URLs into concise, manageable links.

## Introduction

URLShortener is a practical project designed to demonstrate clean architecture, efficient data handling, and scalable service design. It serves as both a robust backend application and a learning resource for Go developers. This project showcases best practices in API design, database management, and performance optimization.

## Features

- **Base62 Encoding** – Generates clean, URL-safe short codes using Base62 (0–9, a–z, A–Z).
- **PostgreSQL as Primary Storage** – Reliable and scalable relational database for core data storage.
- **Redis Caching** – Enhances performance by caching frequently accessed short URLs.
- **Fiber Web Framework** – Minimal and blazing-fast web server inspired by Express.js.
- **Strong Typing** – Ensures clarity and maintainability throughout the codebase.
- **No Globals or `init()` Functions** – Keeps the application predictable and explicit.
- **Standardized Naming** – Follows idiomatic Go practices and project-layout conventions.
- **Independent Packages** – Self-contained modules for easy extension and testing.
- **Intuitive Structure** – Designed for readability and ease of navigation.

## Development

### Local Development
```sh
go run cmd/migration/* --direction=up
go run cmd/server/*