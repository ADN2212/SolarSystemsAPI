# Solar Systems API

This project was created with the purpose of putting my Go knowledge into practice.  
It is a simple API designed to store information about stars and the planets that orbit them.

## Technologies Used
- **GORM** as the ORM.
- **Gin** for creating the endpoints.
- **PostgreSQL** as database.

## Entities
As mentioned, the API consists of two main entities:

- **Stars**: The primary entity of the system, as planets exist *within* their corresponding star system.  
- **Planets**: These belong to specific stars.

## Endpoints

### **stars**
- **POST**: Receives the data for a new star in the request body, adds it to the database, and returns the ID of the newly created star.  
Example request body:
```json
{
    "name": "Proxima A",
    "solarMass": 10
}
```
- **GET**: Receives the star's ID via the URL and returns a solar system from the database (a star with its orbiting planets).  
Example response body:
```json
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

### +planets

- **POST**: Receives the data for a new planet in the request body, adds it to the database, and links it to the specified star. Returns the planet's ID.  
Example request body:
```json
{
    "name": "Jupiter",
    "mass": 30,
    "isLivable": false,
    "starId": 1
}
```
- **DELETE**: Receives the planet's ID via the URL and deletes it from the database.

- **PATCH**: Receives the data to update the planet in the request body. It does **not** allow changing the `starId` field to move the planet to another solar system. The request body follows the same structure as the **POST** method for this endpoint.


> **Note:** Currently, the API does not have user management, but this feature will be added in the future.

## About the Entity Fields

### Stars
- `name`: Name of the star. *(string)*
- `solarMass`: Number of solar masses of the star — how many times the Sun would fit inside it. *(unsigned integer)*

### Planets
- `name`: Name of the planet. *(string)*
- `mass`: Mass of the planet. *(unsigned integer)*
- `isLivable`: `true` if the planet can be inhabited by humans.
- `starId`: ID of the star the planet orbits.
