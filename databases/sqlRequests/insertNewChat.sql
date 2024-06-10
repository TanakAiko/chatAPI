INSERT INTO chats (
        senderId,
        receiverId,
        content,
        statusReceived,
        statusRead,
        createdAt
    )
VALUES (?, ?, ?, ?, ?, ?)