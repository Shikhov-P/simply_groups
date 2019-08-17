## Notes on developing the application
I've taken these notes in order to better understand the technology stack in use.

### How API works
<details>
    <summary>Click to expand!</summary>
    1. `main` package: Once the server is started, it listens to incoming requests and, if a request url is registered, a relevant handler function is invoked from the `controllers` package. 
    2. `controllers` package: incoming json data is decoded into a struct and passed to a relevant function from the `models` package. 
    3. `models` package: request data is manipulated; all necessary db operations take place. The result is returned back to the `controller` function.
    4. `controllers` package: received data is encoded  into json and sent as a reply to the initial request.
</details>


###`auth.go`:

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


#### JWT
JWT is a standard for securely transmitting information as a JSON object.
JWT tokens can be signed or encrypted:
1)  Signed tokens verify the integrity of the claims withing it.
2) Encrypted tokens hide the claims from other parties and provide secrecy between parties.

JWT tokens can be signed using secret (with the HMAC algorithm) or public/private (RSA or ECDSA algorithms) keys.
It can be guaranteed that the party holding the private key is the one to sign the token.

JWT tokens consist of three parts:
1) Header(usually made up of 2 parts: token type (JWT), signing algorithm)
2) Payload:
    * Payload consists of claims, the statements about an entity and additional data. Claims can be public, private and registered.
3) Signature.
    * Produced by signing the encoded header, encoded payload and the secret with the specified algorithm.
    * Is used to verify message integrity.
    * Is used to verify the sender.
Only public information should be included in the Header and Payload parts of the JWT token.

Pros of JWT tokens:
1) Compactness (if compared to SAML) => a good option for HTML anf HTTP environments.
2) Security.
3) Ease of parsing with built-in JSON parsers.


