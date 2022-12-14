# gradebook_app

Gradebook Microservices Application with Go

---

# High Level Architecture

![HLA](img/hla_gradebook.png "High Level Architecture")


---

# Components
1. **Service Registry**
- Service Registration
- Health Monitoring
2. **Log Service**
- Centralized Logging
3. **Grading Service**
- Business Logic
- Data Persistance
4. **Teacher Portal**
- Web Application
- API Gateway
---

# Workflow
* Create a web service (log service)
* Create the registry service (service registry)
* Register the web service
* Deregister the web service
