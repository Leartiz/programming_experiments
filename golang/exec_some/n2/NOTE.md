## n2 - `context.*` on web request

- implementation [here](./exec.go)

### Test

Use in windows console:

`Invoke-WebRequest -Uri "http://localhost:8000" -Headers @{ 'User-Id' = '3' }`