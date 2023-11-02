### Golang Project Readme
**Description**
This is a Golang project with a seeder located in `cmd/seeder/seed.go` and the main file located in `cmd/main.go`.

**Requirements**
* Go 1.21
* copy the `.env.example` file into `.env`

**Installation**
1. Clone the repository:
```
git clone https://github.com/A-H-Pooladvand/example
```
2. Change into the project directory:
```
cd example
```
3. Install the dependencies:
```
go mod init
go mod tidy
```
4. Copy the `.env.example` file to `.env` and edit it with your own environment variables:
```
cp .env.example .env
```
5. Start the seeder:
```
go run cmd/seeder/seed.go
```
6. Start the main application:
```
go run cmd/main.go
```

**Usage**
The seeder can be used to seed the database with dummy data. To run the seeder, simply run the following command:
```
go run cmd/seeder/seed.go
```

The main application can be started with the following command:
```
go run cmd/main.go
```