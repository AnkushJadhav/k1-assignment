# kloudone assignment

### How to build the project?
Build the executable using the following command:
`go build server/main.go`

### How to run the server?
`./main -dsnstring=DSN_STRING -authsecret=AUTH_SECRET -bindip=BIND_IP -bindport=BIND_PORT`

DSN_STRING (mandatory): The ODBC compatible database string to connect to the MYSQL database
AUTH_SECRET (mandatory): The secret string that will be used to sign JWT tokens
BIND_IP (optional): The IP interface to bind the service to. Defaults to 0.0.0.0 if not provided.
BIND_PORT (optional): The port to bind the service to. Defaults to 8080 if not provided.

### How to use the APIs?
1. Login API :  `POST /auth/login`
  - Headers: `Content-Type - Application/JSON`
  - Body: `{"email": "VALID_EMAIL_ID_IN_THE_SYSTEM" , "password": "VALID_PASSWORD_FOR_THE_USER"}`

2. Logout API : `POST /auth/logout`
  - Cookies: `Authentication`

3. Get Users API : `GET /api/users`
  - Cookies: `Authentication`
  - Query params: 
    - `limit=x` (mandatory) - the page size for results
    - `marker=x` (optional) - the last viewed user id for the next page of users
    - `sort=x:asc|x:desc` (optional) - the attribute to sort the data by (asc or desc). Multiple sort parameters can be passed to sort on multiple attributes
    - `name=x` (optional) - search by user name
    - `email=x` (optional) - search by user email id

4. Get User By ID API : `GET /api/user/:id`
  - Cookies: `Authentication`

5. Create user API : `POST /api/user`
  - Cookies: `Authentication`
  - Headers: `Content-Type - Application/JSON`
  - Body: `{"name": "NAME_OF_THE_USER", "email": "VALID_EMAIL_ID_OF_THE_USER", "password": "A_VALID_PASSWORD"}`
  - Note: The user created by this API can be used in the login API later

6. Edit User API : `PUT /api/user/:id`
  - Cookies: `Authentication`
  - Headers: `Content-Type - Application/JSON`
  - Body: `{"name": "NAME_OF_THE_USER_TO_EDIT", "email": "VALID_EMAIL_ID_OF_THE_USER_TO_EDIT"}`

7. Delete User API : `DELETE /api/user/:id`
  - Cookies: `Authentication`
  - Headers: `Content-Type - Application/JSON`