# http-mux
 
Request:
```
[
        "http://localhost:8080/test?1",
        "http://localhost:8080/test?2",
        "http://localhost:8080/test?3"
]
```
Response:
```
{
    "Responses": [
        {
            "Code": 200,
            "Response": "/test?1",
            "Error": null
        },
        {
            "Code": 200,
            "Response": "/test?2",
            "Error": null
        },
        {
            "Code": 200,
            "Response": "/test?3",
            "Error": null
        }
    ]
}
```
