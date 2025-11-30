# Go Real-Time Broadcast Chat

![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go&logoColor=white)
![License: MIT](https://img.shields.io/badge/License-MIT-green)

**A minimal real-time chat system in Go with TCP sockets, goroutines,
channels, and mutex.**\
Converts a previous RPC chat assignment into a real-time broadcasting
system.

------------------------------------------------------------------------

## 🔹 Features

-   Real-time message broadcasting\
-   Notify users when someone joins\
-   No self-echo (sender does not receive own messages)\
-   Concurrent handling of multiple clients\
-   Goroutines + channels for communication\
-   Mutex for safe shared client list access

## 🚀 How to Run

### Run Server

``` bash
go run server.go
```

### Run Client

``` bash
go run client.go localhost:9000
```

## 🛠 Technologies Used

-   Go (Golang)\
-   TCP Sockets\
-   Concurrency (goroutines)\
-   Channels\
-   Mutex synchronization

## ✅ Assignment Requirements

✔ Real-time broadcasting\
✔ Notify all clients when user joins\
✔ No self-echo\
✔ Goroutines + channels\
✔ Mutex-protected shared list\
✔ Multiple clients

## 📂 Project Structure

    chat-broadcast/
    │
    ├── server.go
    ├── client.go
    └── go.mod
