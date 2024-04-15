Alif developer test
=================

## Summary
The app receives event info via post request and creates notification for the user. Then, the job (worker) takes unsent
notification list and prints them to terminal.

- The app runs in localhost
- The base url is http://localhost:8080/api/vi
- The information about API is covered with swagger API

### Sample Events
```javascript
{
  "orderType": "Purchase",
  "sessionId": "29827525-06c9-4b1e-9d9b-7c4584e82f56",
  "card": "4433**1409",
  "eventDate": "2023-01-04 13:44:52.835626 +00:00",
  "websiteUrl": "https://amazon.com"
},

{
  "orderType": "CardVerify",
  "sessionId": "500cf308-e666-4639-aa9f-f6376015d1b4",
  "card": "4433**1409",
  "eventDate": "2023-04-07 05:29:54.362216 +00:00",
  "websiteUrl": "https://adidas.com"
},

{
  "orderType": "SendOtp",
  "sessionId": "500cf308-e666-4639-aa9f-f6376015d1b4", 
  "card": "4433**1409", 
  "eventDate": "2023-04-06 22:52:34.930150 +00:00",
  "websiteUrl": "https://somon.tj"
}
```
