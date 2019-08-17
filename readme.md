### Notes on developing the application
I've taken these notes in order to better understand the technology stack in use.

####`auth.go`:

The `JwtAuthentication` is an http.Handler function which is a *wrapper*.
Wrappers are usually used for:
* Logging and tracing
* Writing common response headers
* Validating the request (e.g. checking auth credentials which is the case in my app)

The main idea of using wrappers is to take in the `original` http.Handler and return a new one that does something before and/or after calling the `ServeHTTP` on the `original`.

The general pattern for the `JwtAuthentication` wrapper is:
```
func Middleware(next http.Handler) http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Do something before next.ServeHTTP(w, r) 
    next.ServeHTTP(w, r)
    // Do something after next.ServeHTTP(w, r) 
  })
}
```

#### Contexts
**Intro and definition**

Contexts are a means to move data between callsites in a request chain. Often this reduces to moving data between middleware.
Contexts are type-unsafe, i.e. context data can't be checked by a compiler. That's why they should be used in certain cases only.
**When to use contexts?**
Contexts should be used for request scoped data, i.e. data that exists once a request has begun.
E.g.: 
* user ids extracted from headers
* auth tokens tied to cookies or session ids, etc.

**Links:**
[Article from peter.bourgon.org](https://peter.bourgon.org/blog/2016/07/11/context.html)

