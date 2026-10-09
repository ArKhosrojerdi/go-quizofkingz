# Users
POST /users/register
POST /users/login
POST /users/refresh
GET /users/profile

Login and refresh responses contain `access_token` (2-hour lifetime) and
`refresh_token` (30-day lifetime). Refresh requests send the refresh token as
`{"refresh_token":"..."}`. Refresh tokens are stateless and cannot be revoked
server-side; refreshing returns a new token pair.