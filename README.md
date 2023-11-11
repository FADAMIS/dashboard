# FADAMIS Dashboard

FADAMIS dashboard is a web dashboard used for administrating SSPŠ camps like TechDays or HackDays. API reference below.

---
<details>

<summary><code>POST /register</code></summary>

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

---

**JSON example**
```json
{
    "name": "John",
    "surname": "Smith",
    "email": "john@smith.com",
    "phone": "777888999"
}
```

---

**cURL example**
```javascript
curl -X POST "http://localhost/register" -d '{"name": "John", "surname": "Smith", "email": "john@smith.com", "phone": "777888999"}' -H "Content-Type: application/json"
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

<summary><code>GET /participants</code></summary>

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
curl "http://localhost/participants"
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
            "ID": 1,
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

<summary><code>GET /food</code></summary>

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
curl "http://localhost/food"
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
                        "ID": 1,
                        "name": "John",
                        "surname": "Smith", 
                        "email": "john@smith.com",
                        "phone": "777888999",
                        "food_id": 1
                    }
                ]
        }
    ]
}
```

</details>






---




<details>

<summary><code>POST /food</code></summary>

Adds new available food to the database

---
### Request

Headers:

- `Content-Type: application/json`

---

| request data |  data type  |  description |
|--------------|-------------|--------------|
| name         |  string     |  food's name |

---

**JSON example**
```json
{
    "name": "Pizza Prosciutto",
}
```

---

**cURL example**
```javascript
curl -X POST "http://localhost/food" -d '{"name": "Pizza Prosciutto"}' -H "Content-Type: application/json"
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
    "message": "food added",
    "food": "Pizza Prosciutto"
}
```

</details>





---



<details>

<summary><code>POST /order/:name</code></summary>

Orders food for the participant specified by `:name` parameter

- `:name` must be a sha256 sum of `name+surname`
- Example:
```
sha256(JohnSmith)
---->
9d3e2c3ef4399d27897e1d918151cac74ed7b2bee028fea50d29d7d8ea3f925e
---->
http://localhost/order/9d3e2c3ef4399d27897e1d918151cac74ed7b2bee028fea50d29d7d8ea3f925e
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
curl -X POST "http://localhost/order/9d3e2c3ef4399d27897e1d918151cac74ed7b2bee028fea50d29d7d8ea3f925e" -d '{"id": 1, "name": "Pizza Prosciutto"}' -H "Content-Type: application/json"
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


Backend API written in Go and Gin web framework by [Fabucik](https://github.com/Fabucik)

Frontend by [DuckyScr](https://github.com/DuckyScr)