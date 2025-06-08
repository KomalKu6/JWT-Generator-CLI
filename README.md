# JWT-Generator-CLI -GoLang

A command-line tool that generates secure **JWT tokens** (JSON Web Tokens) based on user input.  
This tool simulates **how login systems** issue tokens in real-world cloud platforms.

---

## 📌 Why I Built This

In cloud-native applications, we need to authenticate users **without storing sessions on the server** (i.e., *stateless authentication*). JWTs allow us to:
- ✅ Authenticate users securely
- ✅ Carry user roles (admin, user) in the token
- ✅ Set an expiration to limit token validity
- ✅ Avoid server-side session storage

This CLI tool helps simulate the **JWT issuance** process during login.

----

## 🛠 Features

- Take username & role as input via CLI
- Create JWT with `username`, `role`, and `expiry`
- Sign token using HMAC-SHA256
- Print full JWT token to use in Authorization header

##  Example:
Enter username: komal
Enter role (admin/user): admin

Generated JWT token:
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

