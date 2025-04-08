# Golang TaskManager Folder Structure

### How to start the project
```bash
go run cmd/api/main.go
go run cmd/cron/main.go
```

```bash
# Create a new task
curl -X POST http://localhost:8080/api/tasks/ -H "Content-Type: application/json" -d '{"title":"My First Task","description":"Task description here","due_date":"2025-04-15"}'

# List all tasks
curl http://localhost:8080/api/tasks/

# Get a specific task (replace 1 with the actual task ID)
curl http://localhost:8080/api/tasks/1

# Update a task
curl -X PUT http://localhost:8080/api/tasks/1 -H "Content-Type: application/json" -d '{"title":"Updated Task Title"}'

# Mark a task as complete
curl -X PUT http://localhost:8080/api/tasks/1/complete

# Delete a task
curl -X DELETE http://localhost:8080/api/tasks/1
```

