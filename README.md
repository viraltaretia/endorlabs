# Endorlabs project


### Installation

- `make all`: Builds docker app image and launches app and redis containers
- `make clean`: Cleans up and remove all containers and images.


## Usage

- `bash post_script.sh`: Storing multiple values inside launched redis store
- Getting values by name
```shell
curl -XGET http://0.0.0.0:8080/objects/name/Dia  | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   474  100   474    0     0  32494      0 --:--:-- --:--:-- --:--:-- 67714
[
  {
    "name": "Dia",
    "id": "Person:Dia:376ec918-2aed-4873-8335-2477e41d41c3",
    "age": 7,
    "kind": "Person"
  },
  {
    "name": "Dia",
    "id": "Person:Dia:a7b301b0-f294-4abb-a7b6-f31639bc4cb0",
    "age": 35,
    "kind": "Person"
  },
]
```
- Getting all items by kind
```shell
curl -XGET http://0.0.0.0:8080/objects/kind/Person | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  8015    0  8015    0     0   221k      0 --:--:-- --:--:-- --:--:--  301k
[
  {
    "name": "Samar",
    "id": "Person:Samar:20725b5f-a2d3-4198-89bf-67dca97a6fe0",
    "age": 5,
    "kind": "Person"
  },
  {
    "name": "Viral",
    "id": "Person:Viral:1252ae73-14a3-4fb2-a698-7f56a1f25c2c",
    "age": 80,
    "kind": "Person"
  },
]
```

- Getting item by ID
```shell
curl -XGET http://0.0.0.0:8080/objects/Person:Samar:20725b5f-a2d3-4198-89bf-67dca97a6fe0 | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    97  100    97    0     0  11288      0 --:--:-- --:--:-- --:--:-- 48500
{
  "name": "Samar",
  "id": "Person:Samar:20725b5f-a2d3-4198-89bf-67dca97a6fe0",
  "age": 5,
  "kind": "Person"
}
```

- Deleting item by ID
```shell
curl -XDELETE http://0.0.0.0:8080/objects/Person:Samar:20725b5f-a2d3-4198-89bf-67dca97a6fe0 | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
```
