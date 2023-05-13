# gradebook_app

Gradebook Microservices Application with Go

![GitHub Last Commit](https://img.shields.io/github/last-commit/vmilovanovicc/gradebook_app)
![Github Top Language](https://img.shields.io/github/languages/top/vmilovanovicc/gradebook_app)
![Go Version](https://img.shields.io/github/go-mod/go-version/vmilovanovicc/gradebook_app)
![Licence](https://img.shields.io/github/license/vmilovanovicc/gradebook_app)


---

# High Level Architecture

---

# Components

| Service Registry     | Log Service         | Grading Service  | Teacher Portal  |
|----------------------|---------------------|------------------|-----------------|
| Service Registration | Centralized Logging | Business Logic   | Web Application |
| Health Monitoring    |                     | Data Persistence | API Gateway     |

---

# Workflow

**Service Registration**
* Create a web service (log service)
* Create the registry service (service registry)
* Register the web service
* Deregister the web service
---
**Service Discovery**
* Create the grading service
* Request required services on startup
* Notify when new services start
* Notify when services shut down
---
**Service Monitoring**
* Create the Teacher Portal - a web app
* Add health checks

