# Package google.rpc

## Index

- [`Status`](#M_Status) (message)

---

<h2 id="M_Status">Message <code>Status</code></h2>

The `Status` type defines a logical error model that is suitable for
 different programming environments, including REST APIs and RPC APIs. It is
 used by [gRPC](https://github.com/grpc). Each `Status` message contains
 three pieces of data: error code, error message, and error details.

 You can find out more about this error model and how to work with it in the
 [API Design Guide](https://cloud.google.com/apis/design/errors).

| Fields | |
|---|---|
|`code`|The status code, which should be an enum value of<br/>[google.rpc.Code][google.rpc.Code].|
|`message`|A developer-facing error message, which should be in English. Any<br/>user-facing error message should be localized and sent in the<br/>[google.rpc.Status.details][google.rpc.Status.details] field, or localized<br/>by the client.|
|`details`|A list of messages that carry the error details.  There is a common set of<br/>message types for APIs to use.|

