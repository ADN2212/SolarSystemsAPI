# Solar Systems API

This project was created with the purpose of putting my Go knowledge into practice.  
It is a simple API designed to store information about stars and the planets that orbit them.

## Technologies Used
- **GORM** Provides an abstraction layer over SQL, making it easier to define models, handle migrations, and interact with the database using idiomatic Go code instead of writing raw SQL queries.
- **Gin** A lightweight yet powerful HTTP web framework for Go, used to define routes, handle requests and responses, and provide middleware support with high performance.
- **PostgreSQL** A reliable, open-source relational database known for its robustness, advanced features, and compliance with SQL standards, used to persist and query application data.
- **golang-jwt** A library to generate, sign, and validate JSON Web Tokens (JWT), enabling secure authentication and authorization across the application.
- **bcrypt** to encrypt the user credentials: Used for hashing passwords securely before storing them in the database, protecting sensitive user data from being compromised even if the database is exposed.

## Entities
The API have three main entities:

- **Stars**: The primary entity of the system, as planets exist *within* their corresponding star system.  
- **Planets**: These belong to specific stars.
- **Users**: Which, based on their roles, can perform actions on the endpoints.

## Endpoints

### **STARS**
- **POST**: Receives the data for a new star in the request body, adds it to the database, and returns the ID of the newly created star.  
```json
"Example request body:"
{
    "name": "Proxima A",
    "solarMass": 10
}
```
- **GET**: Receives the star's ID via the URL and returns a solar system from the database (a star with its orbiting planets).  
```json
"Example response body:"
{
    "id": 1,
    "name": "Sol",
    "solarMass": 1,
    "planets": [
        {
            "id": 1,
            "name": "Tierra",
            "mass": 20,
            "isLivable": true
        }
    ]
}
```
- **DELETE**: Receives the star's ID via the URL and deletes the entire solar system from the database (the star and all its orbiting planets).

- **PATCH**: Receives the star's ID via the URL and updates it with the provided data. The request body follows the same structure as the **POST** method for this endpoint.

### **PLANETS**

- **POST**: Receives the data for a new planet in the request body, adds it to the database, and links it to the specified star. Returns the planet's ID.  
```json
"Example request body:"
{
    "name": "Jupiter",
    "mass": 30,
    "isLivable": false,
    "starId": 1
}
```
- **DELETE**: Receives the planet's ID via the URL and deletes it from the database.

- **PATCH**: Receives the data to update the planet in the request body. It does **not** allow changing the `starId` field to move the planet to another solar system. The request body follows the same structure as the **POST** method for this endpoint.

### **USERS**

- **POST**: Receives the data for a new user in the request body and adds it to the database. This endpoint works as a **SignUp** method.  
```json
"Example request body:"
{
    "userName": "Manetha8",
    "passWord": "123456",
    "role": "human"
}
```
As it can bee seen in the example there is a rol key that can be "human" or "god". user of type "human" can only perform actions on the **PLANETS** endpoint while users of type "god" can only perform actions in the **STARS** endpoints. 

- **GET**: Receives the user credentials in the request body and validates them against the database. If the credentials are correct, it returns a JWT token that expires in 5 minutes.  
This token must be sent in the **Authorization** header of each request using the format `'Bearer ${myJWT}'` (JavaScript string literal syntax).  
This endpoint works as a **LogIn** method.  

"Example request body:"
```json
{
    "userName": "Manetha5",
    "passWord": "123456"
}
```
- **DELETE**: Receives the user credentials via the request body, extracts the JWT from the **Authorization** header, and stores it in a **deleted_tokens** table that functions as a blacklist. Any token in that table becomes unusable, even if its lifecycle has not yet expired.  
This endpoint works as a **LogOut** method, and the request body and headers are the same as in the login method.  

## About the Entity Fields

### Stars
- `name`: Name of the star. *(string)*
- `solarMass`: Number of solar masses of the star — how many times the Sun would fit inside it. *(unsigned integer)*

### Planets
- `name`: Name of the planet. *(string)*
- `mass`: Mass of the planet. *(unsigned integer)*
- `isLivable`: `true` if the planet can be inhabited by humans.
- `starId`: ID of the star the planet orbits.
