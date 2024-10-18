# Golang JWT Authentication Template

This project is a JWT-based authentication template written in Golang. It provides basic routes for user registration, login, logout, and a secure route for fetching all users, using JSON Web Tokens (JWT) to protect endpoints. The project includes a simple custom router and middleware for handling token verification.

## Features
- User registration and login
- Password hashing with bcrypt
- JWT token generation and verification
- Middleware for protecting routes
- Custom routing mechanism
- Basic user validation with `validator`

## Technologies Used
- **Go (Golang)**
- **JWT (github.com/golang-jwt/jwt/v5)**: For generating and verifying JWT tokens.
- **bcrypt (golang.org/x/crypto)**: For password hashing.
- **Go Playground Validator (github.com/go-playground/validator/v10)**: For request validation.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/golangJWTAuthTemplate.git
    cd golangJWTAuthTemplate
    ```

2. Initialize the project and install dependencies:
    ```bash
    go mod tidy
    ```

3. Run the server:
    ```bash
    go run main.go
    ```

## Endpoints

### 1. Register a new user
- **Endpoint**: `/register`
- **Method**: `POST`
- **Body** (JSON):
    ```json
    {
      "email": "user@example.com",
      "password": "password123"
    }
    ```
- **Response**: User object or validation error.

### 2. Login
- **Endpoint**: `/login`
- **Method**: `POST`
- **Body** (JSON):
    ```json
    {
      "email": "user@example.com",
      "password": "password123"
    }
    ```
- **Response**: A JWT token is set in an HTTP-only cookie for authentication.

### 3. Logout
- **Endpoint**: `/logout`
- **Method**: `POST`
- **Description**: Logs out the user by invalidating the JWT cookie.

### 4. Get All Users (Protected Route)
- **Endpoint**: `/get-all`
- **Method**: `GET`
- **Description**: Fetches all users. Requires a valid JWT token.
  
  This route is protected with middleware that verifies the JWT token before allowing access.

## Middleware
The project includes an authentication middleware to protect sensitive routes. It checks for the JWT token in cookies and verifies it before allowing access to protected routes.

```go
func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("jwt_token")
        if err != nil {
            http.Error(w, "Token missing in cookie", http.StatusBadRequest)
            return
        }

        token, err := utils.VerifyToken(cookie.Value)
        if err != nil {
            http.Error(w, "Token verification failed", http.StatusBadRequest)
            return
        }

        fmt.Printf("Token verified successfully. \nClaims:%v", token.Claims)
        next.ServeHTTP(w, r)
    })
}
```

## Dependencies

This project uses the following Go packages:

- **github.com/golang-jwt/jwt/v5**: For creating and verifying JWT tokens.
- **golang.org/x/crypto**: For secure password hashing using bcrypt.
- **github.com/go-playground/validator/v10**: For validating user input.

## How to Use

- **Register**: Send a POST request to `/register` with a JSON body containing an email and password.
- **Login**: Send a POST request to `/login` with valid user credentials to get a JWT token set in the cookies.
- **Access Protected Routes**: Use the `/get-all` route to access protected resources by passing the JWT token in the cookies.
- **Logout**: Call `/logout` to invalidate the JWT token and log the user out.

## License

This project is licensed under the MIT License.
