# Design

1. Client_A requests new session url `POST https://msg.etwo.dev/session/new`
  - server returns session url
2. Client_A connects to session `GET https://msg.etwo.dev/session/{id}/`
  - server returns a session secret
3. Client_B connects to session `GET https://msg.etwo.dev/session/{id}/`
4. Client_A asks for identification `GET https://msg.etwo.dev/session/{id}/`

## Create a new session
Endpoint: POST https://msg.etwo.dev/session/new
Description: Client_A requests a new session.
Payload: Client_A creates passphrase for session.
```json
{
  "passphrase": "alternate-doggy-genius"
}
```
Response: Server returns a new session URL.
```json
{
  "session_id": "1123788108584984636",
  "session_secret": "UG91cmluZzYlRW1vdGljb24lU2N1YmE="
}
```


## Connect to a session
Endpoint: POST https://msg.etwo.dev/session/{id}
Description: Client_B connects to the session.
Payload: Client_B uses passphrase to connect.
```json
{
  "passphrase": "alternate-doggy-genius"
}
```
Response: Server returns session secret
```json
{
  "session_secret": "UG91cmluZzYlRW1vdGljb24lU2N1YmE="
}
```

## WebSocket connection
Endpoint: ws://msg.etwo.dev/session/{id}/stream
```js
[type][auth][binary]
```
Description: Client_A or Client_B establishes a WebSocket connection with the server. This will be used to send and receive messages in real-time. During the handshake, the client generates a pair of public/private encryption keys locally. The public key is sent to the server as part of the handshake.
