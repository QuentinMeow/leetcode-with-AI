# Object Storage - S3 / GCS

Examples: Amazon S3, Google Cloud Storage, Azure Blob Storage, MinIO.

## 30-Second Interview Answer

Use object storage for large binary objects, immutable files, media, backups, logs, and data lakes. Store metadata in a database; store bytes in object storage; serve hot public content through a CDN.

## Use When

- Data is large: images, videos, PDFs, exports, backups, logs.
- Access is by object key rather than rich queries.
- Objects are immutable or rarely updated.
- Durability and low storage cost matter.
- CDN delivery or signed upload/download URLs are needed.

## Avoid When

- You need row-level transactions and indexes.
- You need frequent small updates inside a file.
- You need low-latency random writes.
- You need relational constraints or full-text search directly over contents.

## Core Model

Object storage stores objects in buckets. Each object has a key, bytes, metadata, and access policy.

Typical split:

- Database row: object ID, owner, content type, size, status, permissions, storage key.
- Object storage: actual bytes.
- CDN: cached public/read-heavy delivery.

## Query And Indexing

Object storage is not for rich querying. Query metadata from SQL/document/key-value, then fetch object bytes by key.

Interview detail: never list a bucket on a hot path. Maintain object metadata in a real database.

## Consistency And Durability

Modern object stores provide very high durability and generally strong read-after-write behavior for object operations, but application workflows still need state handling.

Upload flow:

1. Create metadata row with `pending` status.
2. Return signed upload URL.
3. Client uploads bytes.
4. Callback/event marks object `ready`.
5. Background workers transcode, scan, index, or generate thumbnails.

## Scaling

Object stores scale naturally for storage and bandwidth, but design still matters:

- Use CDN for hot downloads.
- Use multipart uploads for large files.
- Use lifecycle policies for archival/deletion.
- Avoid hot object keys if provider guidance requires key distribution.
- Separate originals, derived thumbnails, and transcoded variants.

## Failure Modes

- Metadata says ready but upload failed.
- Object exists but DB transaction failed.
- CDN serves stale content.
- Signed URLs leak or live too long.
- Background processing fails after upload.
- Deletion must remove metadata, objects, derived files, and CDN cache.

## Data Modeling

Metadata fields:

- object_id
- owner_id
- bucket/key
- content_type
- size
- checksum
- upload_status
- visibility
- created_at
- retention/deletion policy

## Interview Examples

- Photo/video sharing app.
- File upload service.
- User-generated attachments.
- Data lake for analytics.
- Database backups and export files.

## Senior-Level Tradeoffs

- Object storage is cheap and durable, but not queryable like a database.
- Signed URLs reduce server bandwidth, but require authorization and expiry design.
- CDN improves latency, but adds cache invalidation complexity.
- Metadata/object dual writes need reconciliation jobs.

## Common Mistakes

- Storing large media bytes in SQL rows.
- Listing buckets for user-facing pages.
- Forgetting upload state transitions.
- Ignoring virus scanning, content validation, and access control.
