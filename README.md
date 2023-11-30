# FADAMIS Dashboard

FADAMIS dashboard is a web dashboard used for administrating SSPŠ camps like TechDays or HackDays. API reference below.

## Dependencies
- Docker (instalation guide [here](https://docs.docker.com/engine/install/))

## Instalation
**If your machine is booted with systemd (dashboard runs automatically on startup):**
``` bash
git clone https://github.com/FADAMIS/dashboard
cd dashboard
sudo make # this adds systemd service to your machine
```

---

**Others:**
``` bash
git clone https://github.com/FADAMIS/dashboard
cd dashboard
docker-compose build
```
Then run with `docker-compose up`





## After install

When the dashboard is first started, the database will have sample data in `camps`, `admins` and `foods` tables.

**Default admin login** - `admin`:`admin`


## API Reference

<details>

<summary><code>POST /api/register</code></summary>

Registers a new participant

---
### Request

Headers:

- `Content-Type: application/json`

---

| request data |  data type  |  description                 |
|--------------|-------------|------------------------------|
| name         |  string     | participant's name           |
| surname      |  string     | participant's surname        |
| email        |  string     | participant's email          |
| phone        |  string     | participant's phone number   |
| camp_id      |  int        | camp's id

---

**JSON example**
```json
{
    "name": "John",
    "surname": "Smith",
    "email": "john@smith.com",
    "phone": "777888999",
    "camp_id": 1
}
```

---

**cURL example**
```javascript
curl -X POST "http://localhost/api/register" -d '{"name": "John", "surname": "Smith", "email": "john@smith.com", "phone": "777888999"}' -H "Content-Type: application/json"
```

### Response

---

| Status code  |  Content Type    |
|--------------|------------------|
| 200          | application/json |

---


**Successful response**
```json
{
    "message": "register successful"
}
```

</details>




---





<details>

<summary><code>GET /api/camp</code></summary>

Returns all available camps

---
### Request

Headers:

- None

---

| request data |  data type  |
|--------------|-------------|
| none         |  none       |

---

**cURL example**
```javascript
curl "http://localhost/api/camp"
```

### Response

---

| Status code  |  Content Type    |
|--------------|------------------|
| 200          | application/json |

---

`date` is `int64` (it is stored as Unix timestamp)


**Successful response**
```json
{
    "camps": [
        {
            "id": 1,
            "name": "TechDays!",
            "participants": null,
            "date": 1701302400,
            "processed": false
        }
    ]
}
```

</details>






---




<details>

<summary><code>GET /api/food</code></summary>

Returns all available food

- Also returns ID of the specific food which is later used for participants to order said food

---
### Request

Headers:

- None

---

| request data |  data type  |
|--------------|-------------|
| none         |  none       |

---

**cURL example**
```javascript
curl "http://localhost/api/food"
```

### Response

---

| Status code  |  Content Type    |
|--------------|------------------|
| 200          | application/json |

---


**Successful response**
```json
{
    "foods": [
        {
            "id": 1,
            "name": "pizza",
            "participants": null,
            "image_path": "/images/322989cf390c0e81fbd89727f2ee7d5f402bf652789d88229751eab26ef2e162.jpeg"
        }
    ]
}
```

</details>



---


<details>

<summary><code>POST /api/order/:name</code></summary>

Orders food for the participant specified by `:name` parameter

- `:name` must be a sha256 sum of `name+surname+camp_id`
- Example:
```
sha256(JohnSmith1)
---->
4850eecac63c272e9009385e399f958599a79d96345d2a0e3500d63e7cf2839c
---->
http://localhost/api/order/4850eecac63c272e9009385e399f958599a79d96345d2a0e3500d63e7cf2839c
```

---
### Request

Headers:

- `Content-Type: application/json`

---

| request data |  data type  |  description |
|--------------|-------------|--------------|
| id           |  int        |    food's ID |
| name         | string      | food's name  |

---

**JSON example**
```json
{
    "id": 1,
    "name": "Pizza Prosciutto",
}
```

---

**cURL example**
```javascript
curl -X POST "http://localhost/api/order/9d3e2c3ef4399d27897e1d918151cac74ed7b2bee028fea50d29d7d8ea3f925e" -d '{"id": 1, "name": "Pizza Prosciutto"}' -H "Content-Type: application/json"
```

### Response

---

| Status code  |  Content Type    |
|--------------|------------------|
| 200          | application/json |

---


**Successful response**
```json
{
    "food": "Pizza Prosciutto",
    "name": "JohnSmith"
}
```

</details>


---





<details>

<summary><code>POST /api/admin/login</code></summary>

Login endpoint for administrators


---
### Request

Headers:

- `Content-Type: application/json`

---

| request data |  data type  |  description     |
|--------------|-------------|------------------|
| username     |  string     | admin's username |
| password     | string      | admin's password |

---

**JSON example**
```json
{
    "username": "admin",
    "password": "supersecretpassword",
}
```

---

**cURL example**
```javascript
curl -X POST "http://localhost/api/admin/login" -d '{"username": "admin", "password": "supersecretpassword"}' -H "Content-Type: application/json"
```

### Response

---

| Status code  |  Content Type    |
|--------------|------------------|
| 200          | application/json |

---

**Sets a session cookie that expires after 6 hours**

**Successful response**
```json
{
    "message": "login successful",
}
```

</details>


---





<details>

<summary><code>GET /api/admin/participants</code></summary>

Returns all registered participants

---
### Request

Headers:

- None

---

| request data |  data type  |
|--------------|-------------|
| none         |  none       |

---

**cURL example**
```javascript
curl "http://localhost/api/participants"
```

### Response

---

| Status code  |  Content Type    |
|--------------|------------------|
| 200          | application/json |

---


**Successful response**
```json
{
    "participants": [
        {
            "id": 1,
            "name": "John",
            "surname": "Smith",
            "email": "john@smith.com",
            "phone": "777888999",
            "food_id": 1                // ID of ordered food
        }
    ]
}
```

</details>






---




<details>

<summary><code>GET /api/admin/food</code></summary>

Returns all available food along with all the participants with specifed food already ordered

- Also returns ID of the specific food which is later used for participants to order said food

---
### Request

Headers:

- None

---

| request data |  data type  |
|--------------|-------------|
| none         |  none       |

---

**cURL example**
```javascript
curl "http://localhost/api/admin/food"
```

### Response

---

| Status code  |  Content Type    |
|--------------|------------------|
| 200          | application/json |

---


**Successful response**
```json
{
    "foods": [
        {
            "id": 1,
            "name": "none",
            "participants": 
                [
                    {
                        "id": 1,
                        "name": "John",
                        "surname": "Smith", 
                        "email": "john@smith.com",
                        "phone": "777888999",
                        "food_id": 1
                    }
                ],
            "image_path": "/images/322989cf390c0e81fbd89727f2ee7d5f402bf652789d88229751eab26ef2e162.jpeg"
        }
    ]
}
```

</details>






---









<details>

<summary><code>POST /api/admin/food</code></summary>

Adds new available food to the database

---
### Request

Headers:

- `Content-Type: multipart/form-data`

---

| request data |  data type  |  description  |
|--------------|-------------|---------------|
| name         |  string     |  food's name  |
| file         |  File       |  food's image |

---

**cURL example**
```javascript
curl -X POST "http://localhost/api/admin/food" -F 'name=Pizza' -F 'file=@/home/admin/pizza.jpg' -H "Content-Type: multipart/form-data"
```

Image's filename is sha256 checksum of itself + file extension

---

**Allowed file types**
- jpg
- png

---

### Response

---

| Status code  |  Content Type    |
|--------------|------------------|
| 200          | application/json |

---


**Successful response**
```json
{
    "message": "food added",
    "food": "Pizza"
}
```

</details>





---




<details>

<summary><code>GET /api/admin/camp</code></summary>

Returns all (even expired) camps

---
### Request

**Requires session cookie**

Headers:

- None

---

| request data |  data type  |
|--------------|-------------|
| none         |  none       |

---

**cURL example**
```javascript
curl "http://localhost/api/admin/camp"
```

### Response

---

| Status code  |  Content Type    |
|--------------|------------------|
| 200          | application/json |

---


**Successful response**
```json
{
    "foods": [
        {
            "id": 1,
            "name": "TechDays",
            "participants": 
                [
                    {
                        "id": 1,
                        "name": "John",
                        "surname": "Smith", 
                        "email": "john@smith.com",
                        "phone": "777888999",
                        "food_id": 1
                    }
                ],
            "date": 1701302400,
            "processed": false
        }
    ]
}
```

</details>



---




<details>

<summary><code>POST /api/admin/camp</code></summary>

Adds new camp to the database

---
### Request

Headers:

- `Content-Type: application/json`

---

| request data |  data type  |  description                   |
|--------------|-------------|--------------------------------|
| name         |  string     |  camp's name                   |
| date         |  int64      |  camp's date as Unix timestamp |

---

**cURL example**
```javascript
curl -X POST "http://localhost/api/admin/camp" -d "{'name': 'TechDays', 'date': 1701302400}" -H "Content-Type: application/json"
```

---

### Response

---

| Status code  |  Content Type    |
|--------------|------------------|
| 200          | application/json |

---


**Successful response**
```json
{
    "message": "camp added",
    "camp": "TechDays"
}
```

</details>





---






<details>

<summary><code>POST /api/admin/close</code></summary>

Closes camp registration and sends email with information.

---
### Request

Headers:

- `Content-Type: application/json`

---

| request data |  data type  |  description                   |
|--------------|-------------|--------------------------------|
| id           |  uint       |  camp's ID                   |
| name         |  string     |  camp's name                   |
| date         |  int64      |  camp's date as Unix timestamp |

---

**cURL example**
```javascript
curl -X POST "http://localhost/api/admin/close" -d "{'id': 1, 'name': 'TechDays', 'date': 1701302400}" -H "Content-Type: application/json"
```

---

### Response

---

| Status code  |  Content Type    |
|--------------|------------------|
| 200          | application/json |

---


**Successful response**
```json
{
    "message": "camp closed",
    "camp": "TechDays"
}
```

</details>





---


Backend API written in Go and Gin web framework by [Fabucik](https://github.com/Fabucik)

Frontend by [DuckyScr](https://github.com/DuckyScr)