# A simple user API service

## Installation

### 1. You need a go development environment setup before everything starts taking off.

### 2. Use `git clone` to clone the repo to your local folder.

```
git clone https://github.com/gogf/gf-demo-user
```

### 3. Import `manifest/sql/user.sql` to your database.

### 4. Update `manifest/config/config.yaml` according to your local configurations if necessary.

### 5. Run command `go run main.go`, and you'll see something as follows if success:

```
  ADDRESS | METHOD |        ROUTE         |                                  HANDLER                                   |         MIDDLEWARE           
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | ALL    | /*                   | github.com/gogf/gf/v2/net/ghttp.internalMiddlewareServerTracing            | GLOBAL MIDDLEWARE            
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | ALL    | /*                   | github.com/gogf/gf/v2/net/ghttp.MiddlewareHandlerResponse                  | GLOBAL MIDDLEWARE            
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | ALL    | /api.json            | github.com/gogf/gf/v2/net/ghttp.(*Server).openapiSpec                      |                              
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | ALL    | /swagger/*           | github.com/gogf/gf/v2/net/ghttp.(*Server).swaggerUI                        | HOOK_BEFORE_SERVE            
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | POST   | /user/check-passport | demo/internal/controller.(*cUser).CheckPassport | service.(*sMiddleware).Ctx   
          |        |                      |                                                                            | service.(*sMiddleware).CORS  
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | POST   | /user/is-signed-in   | demo/internal/controller.(*cUser).IsSignedIn    | service.(*sMiddleware).Ctx   
          |        |                      |                                                                            | service.(*sMiddleware).CORS  
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | GET    | /user/profile        | demo/internal/controller.(*cUser).Profile       | service.(*sMiddleware).Ctx   
          |        |                      |                                                                            | service.(*sMiddleware).CORS  
          |        |                      |                                                                            | service.(*sMiddleware).Auth  
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | POST   | /user/sign-in        | demo/internal/controller.(*cUser).SignIn        | service.(*sMiddleware).Ctx   
          |        |                      |                                                                            | service.(*sMiddleware).CORS  
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | POST   | /user/sign-out       | demo/internal/controller.(*cUser).SignOut       | service.(*sMiddleware).Ctx   
          |        |                      |                                                                            | service.(*sMiddleware).CORS  
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | POST   | /user/sign-up        | demo/internal/controller.(*cUser).SignUp        | service.(*sMiddleware).Ctx   
          |        |                      |                                                                            | service.(*sMiddleware).CORS  
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
```

# GoFrame Sites

### GoFrame Repo

* [https://github.com/gogf/gf](https://github.com/gogf/gf)

### GoFrame Home

* [https://goframe.org](https://goframe.org) (中文)
