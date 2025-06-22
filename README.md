# JWT-Generator-CLI -GoLang
    
A command-line tool that generates secure **JWT tokens** (JSON Web Tokens) based on user input.  
This tool simulates **how login systems** issue tokens in real-world cloud platforms.

---

## Why I Built This

In cloud-native applications, we need to authenticate users **without storing sessions on the server** (i.e., *stateless authentication*). JWTs allow us to:
- Authenticate users securely
- Carry user roles (admin, user) in the token
- Set an expiration to limit token validity
- Avoid server-side session storage

This CLI tool helps simulate the **JWT issuance** process during login.

----

## Features

- Take username & role as input via CLI
- Create JWT with `username`, `role`, and `expiry`
- Sign token using HMAC-SHA256
- Print full JWT token to use in Authorization header

##  Example:
Enter username: komal
Enter role (admin/user): admin

Generated JWT token:
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

---------------------------------------------------------------------------------
Project: JWT Generator CLI in Go

 1. Why Are You Building This Project?
Real-world use case: In cloud platforms (like Cisco Secure Access), when a user logs in, we don’t want to keep them logged in using a session stored on the server.

 Solution? Generate a JWT (JSON Web Token) that carries the user’s data + role + expiry, and is signed using a secret so it can't be tampered with.

 This is exactly how secure APIs, cloud dashboards, microservices, and API gateways handle stateless authentication today.

 2. What Problem Does This Project Solve?

Problem	Solution:

You want a user to log in and receive a token securely ->	Generate JWT containing user details
You want role-based access control -> 	Embed role (admin/user) in JWT
You want the token to expire ->	Add expiry using exp claim
You want to simulate login in a CLI or automation tool ->	Accept input in CLI and output a token

 3. What Are You Implementing Technically?
Feature	Implementation:
CLI input	-> fmt.Scanln() for username and role
JWT generation	-> Using github.com/golang-jwt/jwt/v5
Expiry ->	time.Now().Add(24 * time.Hour).Unix()
Claims -> (data)	jwt.MapClaims{ "username": ..., "role": ..., "exp": ... }
Signing token	-> SignedString(jwtKey) using HS256

HS256

 4. Logic Behind the Project

 Step-by-Step Thinking:
Step	Thought Process
1. Take input from user ->	Ask for username and role
2. Use this input to create a secure token	-> Use JWT claims to hold the data
3. Add an expiry to the token	-> Prevent misuse or forever-valid tokens
4. Sign the token with a secret	-> So it can be verified but not faked
5. Print the token	-> So client can use it in future requests (e.g., API header)

 This mimics exactly how login APIs work in the backend.

 5. How To Build This Project (Step by Step)
mkdir jwt token
cd jwt token
 Terminal : 
 go mod init jwt token
 go get github.com/golang-jwt/jwt/v5

-> Write a code in main.go
Run the file : go run main.go

-----------------------------------------------
Interview Questions
------------------------------------------------

| Question                                                      | Ideal Answer                                                                             |
| ------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| What is JWT?                                                  | A stateless token that holds JSON data (claims) securely and is signed. Used for auth.   |
| Why use JWT instead of sessions?                              | Sessions need server storage; JWTs are stored on the client and verified via signature.  |
| What does `exp` mean in a JWT?                                | It's the token expiration time. The token is invalid after this UNIX timestamp.          |
| What are the parts of a JWT?                                  | Header, Payload (claims), Signature. Each part is base64 encoded and separated by dots.  |
| How did you generate JWT in Go?                               | Used the `golang-jwt` package, created claims, signed with `SignedString()`.             |
| What is the difference between a signed and an encrypted JWT? | Signed = can't be tampered, readable. Encrypted = fully hidden, for high-security needs. |
| How is role-based access controlled in JWT?                   | Embed a `role` field in claims; backend checks it before allowing access.                |


### What Is Stateless Authentication?
Imagine logging into a system like Gmail or Cisco Secure Access. Normally:

 In Traditional (Stateful) Authentication:
You log in → Server creates a session and stores it in memory

Each time you make a request → Server looks up your session

If server crashes or restarts → sessions may be lost

This means:

Server must remember you

Server needs to store session data (RAM/Database)

Not scalable in microservices or cloud

In Stateless Authentication (using JWT):

You log in → Server gives you a JWT token (with your info inside)

You store it on your browser or mobile app

Every request you make → You send the token → Server verifies it, doesn’t store anything

If token is valid, you’re allowed access. No server memory needed.

Stateless = Server doesn’t store who is logged in — it just verifies the token.

--> Where Is Stateless Authentication Used Today?

Industry	Example
Web Applications	Google, Facebook, Netflix use JWT tokens
 Cloud Security	Cisco Secure Access, AWS IAM uses stateless tokens
Mobile Apps	Android/iOS apps store JWT to access APIs
API Gateways	Authenticated access to microservices
DevOps Tools	Kubernetes Dashboard, Jenkins login tokens
Zero Trust Networks	Where every request must be verified using a token

Why Do Engineers Prefer Stateless?
Reason	Why it Matters

Scalable	No sessions to store per user
Faster	Verifies token, doesn’t need DB check
Secure	Tokens have expiry & are signed
Portable	Token can be used across services or APIs
Works well in Microservices	Each service can validate token independently

Note: Stateless authentication means the server doesn't store any user session. Instead, it issues a signed JWT token to the user, which includes user details, roles, and expiry. Each request carries the token, and the server just verifies it using a secret. This is scalable, secure, and used in all modern cloud platforms, including Cisco Secure Access.


Concept	Explanation: 

fmt.Scanln()	Reads user input from terminal (Go CLI)
time.Now().Add(24 * time.Hour)	Adds 1 day to current time for token expiry
MapClaims	Map-like structure where we add data into token
SignedString(jwtKey)	Creates a secure JWT by signing it with secret key
Unix()	Converts time to UNIX timestamp (used in JWT exp)

 Tech Stack
Technology	Use

GoLang	Core programming
golang-jwt	JWT signing & claims
CLI	Terminal-based input
HMAC-SHA256	Secure signing algorithm

Learn GoLang – Resources for Beginners

 Go by Example: https://gobyexample.com/ or
   https://golangr.com/exercises
 Interactive Go Tour: https://go.dev/tour/
 Crash Course	YouTube: FreeCodeCamp GoLang Crash Course
 GitHub	GoLang JWT Repo Examples.

JWT Token Output: 
https://github.com/KomalKu6/JWT-Generator-CLI/blob/b90203382b6f21a0ce2a0fc2c240fa644413f756/Screenshot%20(2).png

