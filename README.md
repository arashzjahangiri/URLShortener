# URLShortener

![Go Version](https://img.shields.io/badge/Golang-1.24-66ADD8?style=for-the-badge&logo=go)
![App Version](https://img.shields.io/github/v/tag/arashzjahangiri/URLShortener?sort=semver&style=for-the-badge&logo=github)
![Repo Size](https://img.shields.io/github/repo-size/arashzjahangiri/URLShortener?logo=github&style=for-the-badge)
![Coverage](https://img.shields.io/codecov/c/github/arashzjahangiri/URLShortener?logo=codecov&style=for-the-badge)

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

## Features

- Base62 Encoding – Generates clean, URL-safe short codes using Base62 (0–9, a–z, A–Z).
- PostgreSQL as Primary Storage – Reliable and scalable relational database backing the core data layer.
- Redis Caching – Boosts performance by caching frequently accessed short URLs.
- Fiber Web Framework – Blazing fast and minimal web server built with Fiber, inspired by Express.js.
- Strong Typing – Ensures clarity and maintainability throughout the codebase.
- No Globals or init() Functions – Keeps the application predictable and explicit.
- Standardized Naming – Follows idiomatic Go practices and project-layout conventions.
- Independent Packages – Each module is self-contained and easy to extend or test.
- Intuitive Structure – Designed for readability and ease of navigation.

## Development

```bash
# local
go run cmd/migration/* --direction=up
go run cmd/server/*

# by compose
cd ./hacks/compose
podman compose -f ./compose.local.yml up -d 
```

## Tests

### Benchmark (Key Generation)

```txt
goos: linux
goarch: amd64
pkg: github.com/mohammadne/fesghel/internal/urls
cpu: Intel(R) Core(TM) i5-8265U CPU @ 1.60GHz
BenchmarkGenerateKey
BenchmarkGenerateKey-8     1419451     871.4 ns/op     112 B/op     6 allocs/op
```

### Load Test (K6)

To install K6:

```bash
sudo dnf install https://dl.k6.io/rpm/repo.rpm
sudo dnf install k6
```

Sample result:

```txt
TOTAL RESULTS

checks_total.......................: 32     0.799981/s
checks_succeeded...................: 53.12% 17 out of 32
checks_failed......................: 46.87% 15 out of 32

✓ liveness is status 200
✓ readiness is status 200
✓ shorten status is 201
✗ response has ID
    ↳  0% — ✓ 0 / ✗ 15

HTTP
http_req_duration.......................................................: avg=15.13ms min=321.37µs med=14.44ms max=26.83ms p(90)=26.59ms p(95)=26.68ms
    { expected_response:true }............................................: avg=15.13ms min=321.37µs med=14.44ms max=26.83ms p(90)=26.59ms p(95)=26.68ms
http_req_failed.........................................................: 0.00%   0 out of 17
http_reqs...............................................................: 17      0.42499/s

EXECUTION
iteration_duration......................................................: avg=61.93µs min=6.47µs   med=21.52µs max=1.02s   p(90)=36.06µs p(95)=42.54µs
iterations..............................................................: 5259588 131486.572865/s
vus.....................................................................: 50      min=0           max=50
vus_max.................................................................: 56      min=56          max=56

NETWORK
data_received...........................................................: 3.0 kB  75 B/s
data_sent...............................................................: 2.9 kB  72 B/s

running (00m40.0s), 00/56 VUs, 5259588 complete and 0 interrupted iterations
healthz ✓ [======================================] 1 VUs   00m00.0s/10m0s  1/1 shared iters
shorten ✓ [======================================] 5 VUs   00m03.1s/10m0s  15/15 shared iters
get     ✓ [======================================] 50 VUs  30
```

## TODOs

- complete readme and docs
- integration api tests
- functional tests
- the pkg helper for integration
