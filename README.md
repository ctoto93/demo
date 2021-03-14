# Golang Students-Courses Demo App Server

This repository is an example of a demo project of students-courses relationship. A student can have many courses, and a course can have many students. The backend should provide CRUD services for both student and course resources. To offer interoperability, it serves its API in both REST and gRPC. Moreover, it should be flexible to persist the resources, either in MongoDB or SQLite.

## Domain Requirement
- Student can only have maximum 30 of total course credits
- A course should have minimum of 5 students and maximum 30 students

## Project Structure

### demo - [domain models declared here]
* db [actual implementation of a repository (MongoDB/SQLite)]
* api [actual implementation of an api server (REST/gRPC)]
* test [test mocks and factories; integration test of db-service]
* cmd [main application executable]


