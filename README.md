<h1> Modern Go Web App Project: Bookings Reservation System </h1>


<h3>
How to run the project:

1. Install Go: 
   
   - Download and follow the installation guide from Go's website: [Go](https://go.dev/dl)

2. Clone the repository

3. Install the packages:
   
   - Run: `go get -u`
  
4. Run the application:

   - Run the command: `go run cmd/web/*.go`

<hr>

Pages and routes:

   - Home page: "/"
   
   - About page: "/about"

Specifications:
- Built in Go 1.18
- Uses the [chi router](https://github.com/go-chi/chi)
- Uses [Alex Edwards SCS](https://github.com/alexedwards/scs/v2) session manager
- Uses [nosurf](https://github.com/justinas/nosurf) to prevent CSRF (Cross-Site Request Forgery) attacks

</h3>