
# AWS Simple API

[![Deploy master branch](https://github.com/alijany/simpleAPI/actions/workflows/deploy.yml/badge.svg)](https://github.com/alijany/simpleAPI/actions/workflows/deploy.yml)
[![codecov](https://codecov.io/gh/alijany/simpleAPI/branch/master/graph/badge.svg?token=I7SQ62E80I)](https://codecov.io/gh/alijany/simpleAPI)

A simple AWS Lambda with Go

``Note:`` The valid rules for defining Model and device identifiers are:
  -  It should start with `id`.
  -  It should end with digits.
  -  It should be a single word, the white spaces are not allowed.
  
example : id23

# Table of content:

- [AWS Simple API](#aws-simple-api)
- [Table of content:](#table-of-content)
- [AWS Credentials](#aws-credentials)
- [Deploy](#deploy)
- [Project Structure](#project-structure)
  - [**/common**](#common)
  - [**/service**](#service)
    - [**service.main.go**](#servicemaingo)
    - [**service.types.go**](#servicetypesgo)
    - [**service.service.*.go**](#serviceservicego)
    - [**service.functions.ts**](#servicefunctionsts)
    - [**service.db.*.go**](#servicedbgo)
    - [**service.resource.*.go**](#serviceresourcego)
    - [**service.mock.go**](#servicemockgo)

---

# AWS Credentials

As a quick setup you can export them as `Actions secrets` so they would be accessible to Serverless and the AWS SDK

* secrets:
  - AWS_ACCESS_KEY_ID
  - AWS_SECRET_ACCESS_KEY

# Deploy

To deploy changes, simply push theme to github, no further action is needed !

---

# Project Structure

## **/common**

If you have a code that will be used everywhere in the project, add it here...

## **/service**

This project uses `Services Pattern` which each service use a single Lambda function can handle a few jobs that are usually related via a data model or a shared infrastructure dependency, like devices and...

services can have different parts, for adding a new service, create a folder with the service name and add these files:

### **service.main.go**
Implement a lambda handler function 

### **service.types.go**
Put every type and interface that related to this service in here.

### **service.service.*.go**
Put your logics like calling apis, working with storage and... in here.

### **service.functions.ts**
Define a `lambda function` thats related to this service

example:

```typescript
const deviceFunctions: NonNullable<Serverless['functions']> = {
    devices: {
        handler: 'bin/main',
        events: [
            { http: { method: 'post', path: '/api/devices', cors: true } },
            { http: { method: 'get', path: '/api/devices/{id}', cors: true } },
        ]
    }
}
```

### **service.db.*.go**
Manipulates dynamoDb data.

### **service.resource.*.go**
Define any kind of AWS resource like dynamoDB tables in here.

### **service.mock.go**
Define all function and resources which be used in testing
