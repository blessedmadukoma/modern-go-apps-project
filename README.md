<h1> Modern Go Web App Project </h1>

<h4>The Go web application built in the Udemy course: Building modern web applications in Go</h4>

<h5>
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

</h5>

<!-- Lesson 5 commits -->
- added the scs package to handle sessions
- created a session manager and SessionLoad method to load and save sessions based on request