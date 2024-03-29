# 3. Requirements and Consumption Calculations

Date: 2023-12-25

## Status

Accepted

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

## Load and Memory Consumption Calculations

### Load Calculation

1. **Message Traffic**
- Total messages sent per day: **2 billion** messages
- Total message views per day: **4 billion** views

2. **Data Transfer for Messages**
- Text messages: **200 GB/day**
- Image messages: **2,000,000 GB/day**

### Memory Consumption

1. **In-Memory Data Storage**
- Total in-memory storage required: **20,000,000 GB**

2. **Database Storage**
- Total database storage for one day of messages: **3,000,300 GB/day**

### Considerations

- These are rough estimates and provide a baseline for the required infrastructure.
- Actual consumption may vary based on user behavior and system efficiency.
- The architecture must be scalable to handle peak loads and growth.
- Data handling strategies like compression and optimized database schemas are crucial.
