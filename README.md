# GraphFire

**GraphFire** is a serverless GraphQL API built with Go, leveraging Google Cloud's Firestore and Cloud Run. This project demonstrates a modern approach to building scalable and efficient APIs using serverless architecture and managed databases.

![GraphFire](GraphFire.webp)

## Features

- **Serverless Architecture**: Leveraging Google Cloud Run for a fully managed and scalable API deployment.
- **Firestore Integration**: Utilizing Google Firestore for a highly available and scalable NoSQL database.
- **GraphQL API**: Providing a flexible and efficient API interface for querying and mutating data.

## Setup and Deployment

### Prerequisites

- Go 1.21 or later
- Google Cloud SDK
- Firebase Admin SDK
- Docker

### Local Development

```bash
git clone https://github.com/TFMV/GraphFire.git
cd GraphFire

go mod download

go run main.go
```

### Docker Deployment

```bash
docker build -t graphfire .
docker run -p 8080:8080 graphfire
```
