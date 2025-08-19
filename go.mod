// This play demonstrates the use of "layer constructors"
// to get around the naming collisions caused by resource names
// like "User" or "Order" shared across multiple layers like
// "Repository", "Service" and "Handlers".
//
// internal
// ├── common
// │   ├── cache/cache.go
// │   ├── config/config.go
// │   ├── database/database.go
// │   └── logger/logger.go
// ├── handler
// │   ├── handler.go
// │   ├── order/order.go
// │   └── user/user.go
// ├── repository
// │   ├── repository.go
// │   ├── order/order.go
// │   └── user/user.go
// └── service
//     ├── service.go
//     ├── email/email.go
//     ├── order/order.go
//     └── user/user.go

// Reddit post:
// https://www.reddit.com/r/golang/comments/1mt9clj/how_i_went_from_hating_di_frameworks_to_building

// Output:
// Constructing Config
// Constructing Logger
// Constructing Database Pool
// Constructing Cache
// Constructing Repository
// Constructing Repository/Order
// Constructing Repository/User
// Constructing Service
// Constructing Service/Email
// Constructing Service/User
// Constructing Service/Order
// Constructing Handler
// Constructing Handler/User
// Constructing Handler/Order
// Registering routes
// Registering routes for Handler/Order
// Registering routes for Handler/User

module github.com/ufukty/goplays
