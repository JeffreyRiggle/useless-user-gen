# Using useless user gen
After you have run the application you should be able to communicate with the application via its REST based API.

## Disclaimer
This is not the worlds best API, it was just some API that worked for a test I was conducting.

## Getting users
Request: `GET localhost:8080/uug/v1/users`

Response: list of users example below.
```json
[
    {"id":0,"firstname":"Ray","lastname":"Lee","occupation":"Superb Cartoonist","gender":"m"},{"id":1,"firstname":"Eva","lastname":"Davis","occupation":"Questionable Oceanographer","gender":"f"}
]
```

## Creating users
Request: `POST localhost:8080/uug/v1/create`

Body:
```json
{
    "NumberToCreate": 2
}
```

Response: list of users created example below.
```json
[
    {"id":0,"firstname":"Ray","lastname":"Lee","occupation":"Superb Cartoonist","gender":"m"},{"id":1,"firstname":"Eva","lastname":"Davis","occupation":"Questionable Oceanographer","gender":"f"}
]
```

## Updating users
Request: `PUT localhost:8080/uug/v1/updaterandom`

Body:
```json
{
    "NumberToUpdate": 2
}
```

Response: list of users updated example below
```json
[
    {"id":0,"firstname":"Bryant","lastname":"Smith","occupation":"Superb Oceanographer","gender":"m"},
    {"id":1,"firstname":"Pete","lastname":"Sanders","occupation":"Questionable Mechanic","gender":"m"}
]
```

## Clearing users
Request: `DELETE localhost:8080/uug/v1/clear`

Response: None.