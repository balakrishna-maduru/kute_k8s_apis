


## Examples to Run api's:

```bash
     curl http://localhost:8080/namespaces
```

```bash
     curl -X POST http://localhost:8080/api/v1/namespaces \
     -H "Content-Type: application/json" \
     -d '{ "name": "theia-development" }'
```

```bash
     curl http://localhost:8080/namespaces/theia-development
```

```bash
     curl -X DELETE http://localhost:8080/api/v1/namespaces/theia-development
```