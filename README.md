**Run the applicaiton using docker compose:**
  Build the images
  - docker-compose up --build
  Start the postgres
  - docker exec -it postgres_db psql -U postgres -d todoapp
  Create the schemas in the postgres
    CREATE TABLE todo_items (
        id VARCHAR(50) PRIMARY KEY,
        title TEXT NOT NULL,
        description TEXT,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL
    );
    
    CREATE TABLE attachments (
        id VARCHAR(50) PRIMARY KEY,
        todo_id VARCHAR(50) REFERENCES todo_items(id) ON DELETE CASCADE,
        file_name TEXT NOT NULL,
        file_type TEXT NOT NULL,
        file_data BYTEA NOT NULL,
        created_at TIMESTAMP NOT NULL,
    	updated_at TIMESTAMP NOT NULL
    );

**Run the application by deploying the helm chart**
  - run helm install to-do-app .\to-do-app-0.1.0.tgz to install the helm chart
  - access the service on port 8080

Below is the sample curl call to create a todo
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: multipart/form-data" \
  -F "title=My First Task" \
  -F "description=This is a test task" \
  -F "file_data=@/path/to/file.pdf"
