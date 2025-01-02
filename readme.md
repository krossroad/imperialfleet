# Imperialfleet

This is api for inventory of General R3-D3's Imperial Fleet. To get this system running, one needs to have docker & docker compose install. And run following command
```
docker-compose up -d
```


## Galactic Fleet API Documentation

### Table of Contents

- [Endpoints](#endpoints)
  - [List Spaceships](#list-spaceships)
  - [Get Spaceship Details](#get-spaceship-details)
  - [Create Spaceship](#create-spaceship)
  - [Update Spaceship](#update-spaceship)
  - [Delete Spaceship](#delete-spaceship)
---

## Endpoints

### 1. **List Space crafts**

- **URL:** `/spaceships`
- **Method:** `GET`
- **Query Parameters:**
  - `name` (optional) - Filter by spaceship name.
  - `class` (optional) - Filter by spaceship class.
  - `status` (optional) - Filter by spaceship status.
- **Response:**
```json
{
  "spaceships": [
    {
      "id": 1,
      "name": "Devastator",
      "status": "Operational"
    },
    {
      "id": 2,
      "name": "Red Five",
      "status": "Damaged"
    }
  ]
}
```

---

### 2. **Get Spaceship Details**

- **URL:** `/spaceships/{id}`
- **Method:** `GET`
- **Path Parameters:**
  - `id` - ID of the spaceship.
- **Response:**
```json
{
  "id": 1,
  "name": "Devastator",
  "class": "Star Destroyer",
  "crew": 35000,
  "image": "https://url.to.image",
  "value": 199999,
  "status": "Operational",
  "armament": [
    { "title": "Turbo Laser", "quantity": "60" },
    { "title": "Ion Cannons", "quantity": "60" },
    { "title": "Tractor Beam", "quantity": "10" }
  ]
}
```

---

### 3. **Create Spaceship**

- **URL:** `/spaceships`
- **Method:** `POST`
- **Request Body:**
```json
{
  "name": "string",
  "class": "string",
  "crew": "integer",
  "value": "integer",
  "status": "string",
  "armament": [
    { "title": "string", "quantity": "integer" }
  ],
  "image": "string"
}
```
- **Response:**
```json
{ "success": true }
```

---

### 4. **Update Spaceship**

- **URL:** `/spaceships/{id}`
- **Method:** `PUT`
- **Path Parameters:**
  - `id` - ID of the spaceship.
- **Request Body:** Same as [Create Spaceship](#create-spaceship).
- **Response:**
```json
{ "success": true }
```

---

### 5. **Delete Spaceship**

- **URL:** `/spaceships/{id}`
- **Method:** `DELETE`
- **Path Parameters:**
  - `id` - ID of the spaceship.
- **Response:**
```json
{ "success": true }
```
