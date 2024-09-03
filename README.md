# go-application
simple application in go-lang

## Resource - Byte size go learning
https://www.bytesizego.com/blog/learning-golang-2024

### Initialize go module 
go mod init go-application

### Spotify Client
## Authorise spotify
- get access token
- using clientId and clientSecret
- send POST request to spotify Token API endpoint by passing content-type Header and body with grant type as clientCredentials and clientId + clientSecret

curl -X POST "https://accounts.spotify.com/api/token" \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "grant_type=client_credentials&client_id=your-client-id&client_secret=your-client-secret"

