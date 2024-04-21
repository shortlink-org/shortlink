# 16. Adopting Domain-Driven Design (DDD) and Clean Architecture Principles

Date: 2024-02-08

## Status

Accepted

## Context

Our goal is to implement robust and scalable solutions from the first attempt, overcoming the complexities and 
collaboration challenges inherent in software development. Recognizing the need for a unified approach to our 
software's business domain and architecture, we see Domain-Driven Design (DDD) as a pathway to improved domain modeling 
and inter-team communication. Simultaneously, Clean Architecture principles, as proposed by Robert C. Martin, 
offer a blueprint for creating a decoupled, maintainable, and adaptable system. This strategic combination is anticipated 
to directly address our project's needs, setting a strong foundation for future development.

## Decision

To comprehensively address the identified challenges, we have decided to integrate both Domain-Driven Design (DDD) 
and Clean Architecture principles into our development and architectural practices. This encompasses:

#### 1. **Adopting Domain-Driven Design (DDD) Principles:** 

To improve communication and collaboration across teams and to more closely align our software design with the business domain.

##### 1.1 Use Rich Domain Models:
 
- Use rich domain models to encapsulate business logic and behavior.
  - [protoc-gen-rich-model](../pkg/protoc/protoc-gen-rich-model/README.md) is a code generator that creates rich domain models from protobuf definitions.

##### 1.2. Value Objects:

Value objects are only private setters and public getters.

#### 2. **Incorporating Clean Architecture Principles:** Structuring our system to ensure:

  - **Independent of Frameworks:** The architecture does not depend on the existence of some library 
of feature laden software. This allows you to use such frameworks as tools, rather than having 
to cram your system into their limited constraints.
  - **Testable:** The business rules can be tested without the UI, Database, Web Server, or any other external element.
  - **Independent of UI:** The UI can change easily, without changing the rest of the system. 
A Web UI could be replaced with a console UI, for example, without changing the business rules.
  - **Independent of Database:** You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. 
Your business rules are not bound to the database.
  - **Independent of any external agency:** In fact your business rules simply don’t know anything at all about the outside world.

## Consequences

What becomes easier or more difficult to do and any risks introduced by the change that will need to be mitigated.

### References

- **DDD:**
  - 📖 [Learning Domain-Driven Design](https://www.oreilly.com/library/view/learning-domain-driven-design/9781098100124/)
  - 🎥🇷🇺[Денис Цветцих — Rich Model и Anemic Model: враги или друзья](https://youtu.be/kVO2OcsWFXs?si=PjqV1793Gh_LDJ3N)
  - 🎥🇷🇺[Domain-Driven Design. Практический минимум](https://gofunc.ru/talks/3f7283f58edb4f0291e683b2afb25d36/?referer=/schedule/days/)
- **Clean Architecture:**
  - [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
