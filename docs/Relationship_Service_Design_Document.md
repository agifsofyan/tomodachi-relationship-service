# Relationship Service Design Document

**Version:** 1.0\
**Status:** Draft\
**Language:** Go 1.24+\
**Framework:** Gin\
**Communication:** - External API → REST - Internal API → gRPC -
Database → PostgreSQL - Future Event → Kafka - Cache → Redis (optional)

------------------------------------------------------------------------

# 1. Overview

Relationship Service bertanggung jawab mengelola seluruh hubungan antar
user pada aplikasi Tomodachi.

Service ini **tidak mengetahui detail Profile**, **tidak mengirim
Notification**, dan **tidak mengelola Chat**.

Semua komunikasi dilakukan melalui REST, gRPC, atau Event.

------------------------------------------------------------------------

# 2. Responsibility

-   Friend Request
-   Accept Friend Request
-   Reject Friend Request
-   Cancel Friend Request
-   Remove Friend
-   Block User
-   Unblock User
-   Get Friend List
-   Get Friend Request
-   Mutual Friends
-   Friend Status
-   Privacy Validation

Tidak bertanggung jawab terhadap: - Authentication - Profile - Chat -
Notification - Nearby Search

------------------------------------------------------------------------

# 3. High Level Architecture

``` text
React / Mobile
      │
REST API (Gin)
      │
Relationship Service
      ├── PostgreSQL
      ├── gRPC Client → Profile Service
      ├── gRPC Client → Auth Service (optional)
      └── Kafka Producer (future)
```

------------------------------------------------------------------------

# 4. Tech Stack

  Component    Technology
  ------------ -------------------------
  Language     Go 1.24+
  Framework    Gin
  ORM          GORM
  Database     PostgreSQL
  Migration    golang-migrate
  Validation   go-playground/validator
  Logger       slog
  Config       Viper / env
  UUID         google/uuid
  gRPC         grpc-go
  Proto        protobuf
  Testing      testify

------------------------------------------------------------------------

# 5. Folder Structure

``` text
relationship-service/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── relationship/
│   │   ├── delivery/
│   │   ├── application/
│   │   ├── domain/
│   │   └── infrastructure/
│   └── shared/
│       ├── config/
│       ├── database/
│       ├── middleware/
│       ├── logger/
│       ├── grpc/
│       └── utils/
├── api/
│   ├── rest/
│   └── proto/
├── migrations/
├── docs/
├── Dockerfile
├── Makefile
├── go.mod
└── README.md
```

------------------------------------------------------------------------

# 6. Layer Architecture

``` text
REST (Gin)
    ↓
Handler
    ↓
Application
    ↓
Domain
    ↓
Repository Interface
    ↓
PostgreSQL Repository
```

------------------------------------------------------------------------

# 7. Domain Model

## FriendRequest

-   id
-   requester_id
-   receiver_id
-   status
-   created_at
-   updated_at

Status: - PENDING - ACCEPTED - REJECTED - CANCELLED

## Friendship

-   id
-   user_id
-   friend_id
-   created_at

## Block

-   id
-   blocker_id
-   blocked_id
-   created_at

------------------------------------------------------------------------

# 8. REST API

Base URL:

    /api/v1/relationships

  Method   Endpoint                       Description
  -------- ------------------------------ ---------------------
  POST     /friends/request               Send friend request
  POST     /friends/request/{id}/accept   Accept request
  POST     /friends/request/{id}/reject   Reject request
  DELETE   /friends/request/{id}          Cancel request
  DELETE   /friends/{friendId}            Remove friend
  POST     /blocks                        Block user
  DELETE   /blocks/{userId}               Unblock user
  GET      /friends                       Friend list
  GET      /friends/request               Pending requests
  GET      /friends/status/{userId}       Relationship status

------------------------------------------------------------------------

# 9. gRPC Service

``` proto
service RelationshipService {
  rpc IsFriend(IsFriendRequest) returns (IsFriendResponse);
  rpc IsBlocked(IsBlockedRequest) returns (IsBlockedResponse);
  rpc FriendStatus(FriendStatusRequest) returns (FriendStatusResponse);
  rpc GetFriendIds(GetFriendIdsRequest) returns (GetFriendIdsResponse);
  rpc GetBlockedIds(GetBlockedIdsRequest) returns (GetBlockedIdsResponse);
}
```

Relationship Service juga bertindak sebagai **gRPC Client** untuk
memanggil Profile Service guna memvalidasi keberadaan user.

------------------------------------------------------------------------

# 10. Database

## friend_requests

-   id UUID
-   requester_id UUID
-   receiver_id UUID
-   status
-   created_at
-   updated_at

## friendships

-   id UUID
-   user_id UUID
-   friend_id UUID
-   created_at

Setiap pertemanan disimpan dua arah: - A → B - B → A

## blocks

-   id UUID
-   blocker_id UUID
-   blocked_id UUID
-   created_at

------------------------------------------------------------------------

# 11. Business Rules

## Send Friend Request

-   Tidak boleh mengirim ke diri sendiri.
-   User tujuan harus ada.
-   Tidak sedang diblokir.
-   Belum berteman.
-   Tidak ada request pending.

## Accept

-   Request harus PENDING.
-   Receiver harus sesuai JWT.
-   Membuat dua record friendship.

## Reject

-   Status menjadi REJECTED.

## Block

-   Menghapus friendship bila ada.
-   Membatalkan pending request.
-   Menyimpan data block.

------------------------------------------------------------------------

# 12. Authentication

Menggunakan JWT dari Auth Service.

    Authorization: Bearer <token>

Middleware memvalidasi token dan menyimpan UserID pada request context.

------------------------------------------------------------------------

# 13. Kafka (Future)

Topics:

-   friend.requested
-   friend.accepted
-   friend.rejected
-   friend.removed
-   friend.blocked
-   friend.unblocked

Notification Service akan menjadi consumer.

------------------------------------------------------------------------

# 14. Redis (Future)

-   Friend count cache
-   Mutual friend cache
-   Recommendation cache
-   Rate limiting

------------------------------------------------------------------------

# 15. Request Flow

``` text
Client
  ↓
Gin Handler
  ↓
Application
  ↓
gRPC Profile Validation
  ↓
Repository
  ↓
PostgreSQL
  ↓
Kafka (future)
  ↓
Response
```

------------------------------------------------------------------------

# 16. Roadmap

1.  Friend Request
2.  Accept / Reject
3.  Block / Unblock
4.  Remove Friend
5.  Mutual Friends
6.  Friend Recommendation
7.  Kafka Integration
8.  Redis Cache
9.  OpenTelemetry
10. Prometheus
