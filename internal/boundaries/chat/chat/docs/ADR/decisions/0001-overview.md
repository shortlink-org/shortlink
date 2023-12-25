# 1. Overview

Date: 2023-12-25

## Status

Accepted

## Context

Read more about this service we can be in common [ADR](../../../../docs/ADR/README.md) this boundary context.

## Decision

Write a chat service that will be used for communication between users.

## Consequences

We use C4 model for describe architecture of this service.
In this ADR we describe only component level.

### C4 Component diagram

```puml
@startuml Component_Diagram
!includeurl https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml

Container(chat_service, "Chat Service", "Elixir")

Component(crud_component, "CRUD Component", "Manages CRUD operations for chats and messages.")
Component(pin_component, "Pin/Unpin Component", "Handles pinning and unpinning of chats and messages.")
Component(notification_component, "Notification Component", "Handles setting and sending notifications.")
Component(media_component, "Media Component", "Manages image, file, and audio handling.")
Component(unread_message_counter, "Unread Message Counter", "Tracks and updates the count of unread messages for users.")
Component(user_status_manager, "User Status Manager", "Tracks and updates user status (online/offline) and time since last online.")
Component(cross_device_sync_manager, "Cross-Device Sync Manager", "Ensures synchronization of user actions across multiple devices.")

System_Ext(speech_service, "Speech-to-Text Service", "External service that converts voice messages to text.")

Rel(chat_service, crud_component, "Uses")
Rel(chat_service, pin_component, "Uses")
Rel(chat_service, notification_component, "Uses")
Rel(chat_service, media_component, "Uses")
Rel(chat_service, unread_message_counter, "Uses")
Rel(chat_service, user_status_manager, "Uses")
Rel(chat_service, cross_device_sync_manager, "Uses")
Rel(chat_service, speech_service, "Sends voice messages for conversion", "Uses API")

@enduml
```

---

## Requirements

### Functional Requirements

1. **Global Accessibility**: The Chat Service must be accessible to users worldwide, supporting internationalization and localization as needed.
2. **Unread Message Tracking**: The application must clearly display the count and list of unread messages to users.
3. **Chat Types**: The service should support both group chats and private messaging functionalities.
4. **Media Handling in Messages**:
  - Maximum image size in a message: 1 MB.
  - Maximum number of images per message: 3.
5. **Text Limitation**: The maximum length of text in a message should be 2000 characters.

### Performance and Scalability Requirements

1. **User Load Handling**: The system should be capable of supporting 200 million unique users daily, with scalable architecture to handle peak loads.
2. **Message Traffic**: On average, each user is expected to send 10 messages and view messages 20 times per day.
3. **High Availability**: The system must aim for 99.995% uptime, translating to a maximum allowable downtime of approximately 4 hours and 23 minutes per year.
4. **Message Delivery Time**: Messages should be delivered to the recipient within 3 seconds under normal network conditions. If the recipient is offline, a push notification must be sent to their mobile device.
5. **Cross-Device Synchronization**: The service must ensure real-time synchronization across multiple devices. If a message is read on one device, it should be marked as read on all other devices where the user's account is active.

### Non-Functional Requirements

1. **Usability**: The user interface should be intuitive and user-friendly, accommodating a diverse global user base.
2. **Security and Privacy**: Strong security measures should be in place to protect user data and communication. Privacy settings should allow users to control their visibility and information sharing.
3. **Data Management**: Efficient handling of large volumes of data, including messages, media files, and user activity logs.
4. **Scalability**: The architecture should be scalable to accommodate growing user numbers and data volumes.
5. **Maintainability and Monitoring**: The system should be maintainable with effective monitoring and logging mechanisms to quickly identify and address issues.
