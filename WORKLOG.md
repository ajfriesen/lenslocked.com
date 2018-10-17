# WORKLOG

## Todo

### Book

| Chapter | Title                                                                              | Read? | Notes |
|---------|------------------------------------------------------------------------------------|-------|------|
| 1.1     | Who is this book for?                                                              | x      |      |
| 1.      | How to use this book                                                               | x      |      |
| 1.3     | Reference materials                                                                | x      |      |
| 1.4     | What if I am already an expert developer?                                          | x      |      |
| 1.5     | What are we building?                                                              | x      |      |
| 1.6     | Conventions used in this book                                                      | x      |      |
| 1.6.1   | Command-line commands are prefixed with a $                                        | x      |      |
| 1.6.2   | The subl or ‘atom’ command means “open with your text editor”                      | x      |      |
| 1.6.3   | Many code listings will be shortened for clarity                                   | x      |      |
| 1.6.4   | All code samples are limited to 60 columns when possible                           | x      |      |
| 1.7     | Accessing the code                                                                 | x      |      |
| 1.7.1   | Why git?                                                                           | x      |      |
| 1.8     | Disclaimer: Not everything is one size fits all                                    | x      |      |
| 1.9     | Commenting your exported types                                                     | x      |      |
| 2.1     | Building the server                                                                | x      |      |
| 2.2     | Demystifying our app                                                               | x      |      |
| 3.1     | Routing with if/else statements                                                    | x      |      |
| 3.2     | Popular routers                                                                    | x      |      |
| 3.2.1   | net/http.ServeMux                                                                  | x      |      |
| 3.2.2   | github.com/julienschmidt/httprouter                                                | x      |      |
| 3.2.3   | github.com/gorilla/mux.Router                                                      | x      |      |
| 3.2.4   | What about <some other router>                                                     | x      |      |
| 3.3     | Using the gorilla/mux router                                                       | x      |      |
| 3.4     | Exercises                                                                          | x      |      |
| 3.4.1   | Ex1 - Add an FAQ page                                                              | x      |      |
| 3.4.2   | Ex2 - Custom 404 page                                                              | x      |      |
| 3.4.3   | Ex3 - [HARD] Try out another router                                                | x      |      |
| 4.1     | What are templates?                                                                | x      |      |
| 4.2     | Why do we use templates?                                                           | x      |      |
| 4.3     | Templates in Go                                                                    | x      |      |
| 4.4     | Creating a template                                                                | x      |      |
| 4.5     | Contextual encoding                                                                | x      |      |
| 4.6     | Exercises                                                                          | x      |      |
| 4.6.1   | Ex1 - Add a new template variable                                                  | x      |      |
| 4.6.2   | Ex2 - Experiment with different data types                                         | x      |      |
| 4.6.3   | Ex3 - [HARD] Learn how to use nested data                                          | x      |      |
| 4.6.4   | Ex4 - [HARD] Create an if/else statement in your template                          | x      |      |
| 5.1     | Model-View-Controller (MVC)                                                        | x      |      |
| 5.2     | Walking through a web request                                                      | x      |      |
| 5.3     | Exercises                                                                          | x      |      |
| 5.3.1   | Ex1 - What does MVC stand for?                                                     | x      |      |
| 5.3.2   | Ex2 - What is each layer of MVC responsible for?                                   | x      |      |
| 5.3.3   | Ex3 - What are some benefits to using MVC?                                         | x      |      |
| 6.1     | The home template                                                                  | x      |      |
| 6.2     | The contact template                                                               | x      |      |
| 6.3     | Creating a reusable Bootstrap layout                                               | x      |      |
| 6.3.1   | Named templates                                                                    | x      |      |
| 6.3.2   | Creating a view type                                                               | x      |      |
| 6.3.3   | Creating the Bootstrap layout                                                      | x      |      |
| 6.4     | Adding a navigation bar                                                            | x      |      |
| 6.5     | Cleaning up our code                                                               | x      |      |
| 6.5.1   | What is globbing?                                                                  | x      |      |
| 6.5.2   | Using filepath.Glob                                                                | x      |      |
| 6.5.3   | Simplifying view rendering                                                         | x      |      |
| 6.5.4   | Moving our footer to the layout                                                    | x      |      |
| 6.6     | Exercises                                                                          | x      |      |
| 6.6.1   | Ex1 - Create an FAQ page with the Bootstrap layout                                 | x      |      |
| 6.6.2   | Ex2 - Update the navbar to link to the FAQ page                                    | x      |      |
| 6.6.3   | Ex3 - Create a new layout                                                          | x      |      |
| 7.1     | Add a sign up page with a form                                                     | x      |      |
| 7.1.1   | Creating a Bootstrap sign up form                                                  | x      |      |
| 7.1.2   | Wrapping our form in a panel                                                       | x      |      |
| 7.1.3   | Adding the sign up link to our navbar                                              | x      |      |
| 7.2     | An intro to REST                                                                   | x      |      |
| 7.2.1   | How REST affects our code                                                          | x      |      |
| 7.3     | Creating our first controller                                                      | x      |      |
| 7.3.1   | Create the users controller                                                        | x      |      |
| 7.3.2   | Moving the sign up page code                                                       | x      |      |
| 7.3.3   | Connecting our router and the users controller together                            | x      |      |
| 7.4     | Processing the sign up form                                                        | x      |      |
| 7.4.1   | Stubbing the create user action                                                    | x      |      |
| 7.4.2   | HTTP request methods and gorilla/mux                                               | x      |      |
| 7.4.3   | Parsing a POST form                                                                | x      |      |
| 7.4.4   | Parsing forms with gorilla/schema                                                  | x      |      |
| 7.4.5   | Keeping our parsing code DRY                                                       | x      |      |
| 7.5     | Cleaning up and creating a static controller                                       | x      |      |
| 7.5.1   | Creating the static controller                                                     | x      |      |
| 7.5.2   | Simplifying the creation of new views                                              | x      |      |
| 7.6     | Exercises                                                                          | x      |      |
| 7.6.1   | Ex1 - Add the FAQ page to the static controller                                    | x      |      |
| 7.6.2   | Ex2 - Create a new controller for galleries                                        | x      |      |
| 7.6.3   | Exercise cleanup                                                                   | x      |      |
| 8.1     | Our web app will use PostgreSQL                                                    | x      |      |
| 8.2     | Setting up PostgreSQL                                                              | x      |      |
| 8.2.1   | Install PostgreSQL                                                                 | x      |      |
| 8.2.2   | Learn how to connect to Postgres                                                   | x      |      |
| 8.2.3   | Learn the basics of SQL                                                            | x      |      |
| 8.2.4   | Gather information needed to connect to your Postgres install                      | x      |      |
| 8.3     | Using Postgres with Go and raw SQL                                                 | x      |      |
| 8.3.1   | Connecting to Postgres with the database/sql package                               | x      |      |
| 8.3.2   | Creating SQL tables to test with                                                   | x      |      |
| 8.3.3   | Writing records with database/sql                                                  |  x    |      |
| 8.3.4   | Querying a single record with database/sql                                         |   x    |     |
| 8.3.5   | Querying multiple records with database/sql                                        |   x   |      |
| 8.3.6   | Writing a relational record                                                        |   x   |      |
| 8.3.7   | Querying related records                                                           |    x   |      |
| 8.3.8   | Delete the SQL tables we were testing with                                         |    x   |      |
| 8.4     | Using GORM to interact with a database                                             |   x    |      |
| 8.4.1   | Installing GORM and connecting to a database                                       |   x    |      |
| 8.4.2   | Defining a GORM model                                                              |    x   |      |
| 8.4.3   | Creating and migrating tables with GORM                                            |   x    |      |
| 8.4.4   | Logging with GORM                                                                  |     x  |      |
| 8.4.5   | Creating a record with GORM                                                        |     x  |      |
| 8.4.6   | Querying a single record with GORM                                                 |     x  |      |
| 8.4.7   | Querying multiple records with GORM                                                |     x  |      |
| 8.4.8   | Creating related models with GORM                                                  |     x  |      |
| 8.4.9   | Querying relational data with GORM                                                 |     x  |      |
| 8.5     | Exercises                                                                          |     x  |      |
| 8.5.1   | Ex1 - What changes won’t the AutoMigrate function provided by GORM handle for you? |     x  |      |
| 8.5.2   | Ex2 - What is gorm.Model used for?                                                 |     x  |      |
| 8.5.3   | Ex3 - Experiment using a few more GORM methods.                                    |       |      |
| 8.5.4   | Ex4 - Experiment with query chaining                                               |       |      |
| 8.5.5   | Ex5 - Learn to execute raw SQL with GORM                                           |       |      |
| 9.1     | Defining a User type                                                               |     x  |      |
| 9.2     | Creating the UserService interface and querying for users                          |     x  |      |
| 9.3     | Creating users                                                                     |   x   |      |
| 9.4     | Querying by email and DRYing up our code                                           |       |      |
| 9.5     | Updating and deleting users                                                        |       |      |
| 9.6     | AutoMigrating and returning errors from DestructiveReset                           |       |      |
| 9.7     | Connecting the user service and controller                                         |       |      |
| 9.7.1   | Adding the name field the sign up form                                             |       |      |
| 9.7.2   | Setting up a user service in our web application                                   |       |      |
| 9.7.3   | Using the users service in our users controller                                    |       |      |
| 9.8     | Exercises                                                                          |       |      |
| 9.8.1   | Ex1 - Add an Age field to our user resource                                        |       |      |
| 9.8.2   | Ex2 - Write a method to query the first user with a specific age                   |       |      |
| 9.8.3   | Ex3 - [HARD] Write a method to query many users by age range                       |       |      |
| 10.1    | Why not use another package or service?                                            |       |      |
| 10.2    | Secure your server with SSL/TLS                                                    |       |      |
| 10.3    | Hash passwords properly                                                            |       |      |
| 10.3.1  | What is a hash function?                                                           |       |      |
| 10.3.2  | Store hashed passwords, not raw passwords                                          |       |      |
| 10.3.3  | Salt and pepper passwords                                                          |       |      |
| 10.4    | Implementing password hashing for our users                                        |       |      |
| 10.4.1  | Adding password fields to the user model                                           |       |      |
| 10.4.2  | Hash passwords with bcrypt before saving                                           |       |      |
| 10.4.3  | Retrieving passwords from the sign up form                                         |       |      |
| 10.4.4  | Salting and peppering passwords                                                    |       |      |
| 10.5    | Authenticating returning users                                                     |       |      |
| 10.5.1  | Creating the login template                                                        |       |      |
| 10.5.2  | Creating the login action                                                          |       |      |
| 10.5.3  | Implementing the Authenticate method                                               |       |      |
| 10.5.4  | Calling Authenticate from our login action                                         |       |      |
| 10.6    | Exercises                                                                          |       |      |
| 10.6.1  | Ex1 - What is a hash function?                                                     |       |      |
| 10.6.2  | Ex2 - What purpose does a salt and pepper serve?                                   |       |      |
| 10.6.3  | Ex3 - Timing attacks                                                               |       |      |
| 11.1    | What are cookies                                                                   |       |      |
| 11.2    | Creating our first cookie                                                          |       |      |
| 11.3    | Viewing cookies                                                                    |       |      |
| 11.3.1  | Viewing cookies in Google Chrome                                                   |       |      |
| 11.3.2  | Viewing cookies with Go code                                                       |       |      |
| 11.4    | Securing our cookies from tampering                                                |       |      |
| 11.4.1  | Digitally signing data                                                             |       |      |
| 11.4.2  | Obfuscating cookie data                                                            |       |      |
| 11.5    | Generating remember tokens                                                         |       |      |
| 11.5.1  | Why do we use 32 bytes?                                                            |       |      |
| 11.6    | Hashing remember tokens                                                            |       |      |
| 11.6.1  | How to use the hash package                                                        |       |      |
| 11.6.2  | Using the crypto/hmac package                                                      |       |      |
| 11.6.3  | Writing our own hash package                                                       |       |      |
| 11.7    | Hashing remember tokens in the user service                                        |       |      |
| 11.7.1  | Adding remember token fields to our User type                                      |       |      |
| 11.7.2  | Setting a remember token when a user is created                                    |       |      |
| 11.7.3  | Adding an HMAC field to the UserService                                            |       |      |
| 11.7.4  | Hashing remember tokens on create and update                                       |       |      |
| 11.7.5  | Retrieving a user by remember token                                                |       |      |
| 11.7.6  | Resetting our DB and testing                                                       |       |      |
| 11.8    | Remembering users                                                                  |       |      |
| 11.8.1  | Storing remember tokens in cookies                                                 |       |      |
| 11.8.2  | Restricting page access                                                            |       |      |
| 11.9    | Securing our cookies from XSS                                                      |       |      |
| 11.10   | Securing our cookies from theft                                                    |       |      |
| 11.11   | Preventing CSRF attacks                                                            |       |      |
| 11.12   | Exercises                                                                          |       |      |
| 11.12.1 | Ex1 - Experiment with redirection                                                  |       |      |
| 11.12.2 | Ex2 - Create cookies with different data                                           |       |      |
| 11.12.3 | Ex3 - [HARD] Experiment with the hash package                                      |       |      |
| 12.1    | Separating responsibilities                                                        |       |      |
| 12.1.1  | Rethinking the models package                                                      |       |      |
| 12.1.2  | Static types vs interfaces                                                         |       |      |
| 12.2    | The UserDB interface                                                               |       |      |
| 12.3    | The UserService interface                                                          |       |      |
| 12.4    | Writing and organizing validation code                                             |       |      |
| 12.4.1  | The straightforward approach                                                       |       |      |
| 12.4.2  | A few more validation examples                                                     |       |      |
| 12.4.3  | Reusable validations                                                               |       |      |
| 12.5    | Writing validators and normalizers                                                 |       |      |
| 12.5.1  | Remember token normalizer                                                          |       |      |
| 12.5.2  | Ensuring remember tokens are set on create                                         |       |      |
| 12.5.3  | Ensuring a valid ID on delete                                                      |       |      |
| 12.6    | Validating and normalizing email addresses                                         |       |      |
| 12.6.1  | Converting emails to lowercase and trimming whitespace                             |       |      |
| 12.6.2  | Requiring email addresses                                                          |       |      |
| 12.6.3  | Verifying emails match a pattern                                                   |       |      |
| 12.6.4  | Verifying an email address isn’t taken                                             |       |      |
| 12.7    | Cleaning up our error naming                                                       |       |      |
| 12.8    | Validating and normalizing passwords                                               |       |      |
| 12.8.1  | Verifying passwords have a minimum length                                          |       |      |
| 12.8.2  | Requiring a password                                                               |       |      |
| 12.9    | Validating and normalizing remember tokens                                         |       |      |
| 13.1    | Rendering alerts in the UI                                                         |       |      |
| 13.2    | Rendering dynamic alerts                                                           |       |      |
| 13.3    | Only display alerts when we need them                                              |       |      |
| 13.4    | A more permanent data type for views                                               |       |      |
| 13.5    | Handling errors in the sign up form                                                |       |      |
| 13.6    | White-listing error messages                                                       |       |      |
| 13.7    | Handling login errors                                                              |       |      |
| 13.8    | Recovering from view rendering errors                                              |       |      |
| 14.1    | Defining the gallery model                                                         |       |      |
| 14.2    | Introducing the GalleryService                                                     |       |      |
| 14.3    | Constructing many services                                                         |       |      |
| 14.4    | Closing and migrating all models                                                   |       |      |
| 14.5    | Creating new galleries                                                             |       |      |
| 14.5.1  | Implementing the gallery service                                                   |       |      |
| 14.5.2  | The galleries controller                                                           |       |      |
| 14.5.3  | Processing the new gallery form                                                    |       |      |
| 14.5.4  | Validating galleries                                                               |       |      |
| 14.6    | Requiring users via middleware                                                     |       |      |
| 14.6.1  | Creating our first middleware                                                      |       |      |
| 14.6.2  | Storing request-scoped data with context                                           |       |      |
| 14.7    | Displaying galleries                                                               |       |      |
| 14.7.1  | Creating the show gallery view                                                     |       |      |
| 14.7.2  | Parsing the gallery ID from the path                                               |       |      |
| 14.7.3  | Looking up galleries by ID                                                         |       |      |
| 14.7.4  | Generating URLs with params                                                        |       |      |
| 14.8    | Editing galleries                                                                  |       |      |
| 14.8.1  | The edit gallery action                                                            |       |      |
| 14.8.2  | Parsing the edit gallery form                                                      |       |      |
| 14.8.3  | Updating gallery models                                                            |       |      |
| 14.9    | Deleting galleries                                                                 |       |      |
| 14.10   | Viewing all owned galleries                                                        |       |      |
| 14.10.1 | Querying galleries by user ID                                                      |       |      |
| 14.10.2 | Adding the Index handler                                                           |       |      |
| 14.10.3 | Iterating over slices in Go templates                                              |       |      |
| 14.11   | Improving the user experience and cleaning up                                      |       |      |
| 14.11.1 | Adding intuitive redirects                                                         |       |      |
| 14.11.2 | Navigation for signed in users                                                     |       |      |
| 15.1    | Image upload form                                                                  |       |      |
| 15.2    | Processing image uploads                                                           |       |      |
| 15.3    | Creating the image service                                                         |       |      |
| 15.4    | Looking up images by gallery ID                                                    |       |      |
| 15.5    | Serving static files in Go                                                         |       |      |
| 15.6    | Rendering images in columns                                                        |       |      |
| 15.7    | Letting users delete images                                                        |       |      |
| 15.8    | Known Bugs                                                                         |       |      |
| 16.1    | Error handling                                                                     |       |      |
| 16.2    | Serving static assets                                                              |       |      |
| 16.3    | CSRF protection                                                                    |       |      |
| 16.4    | Limiting middleware for static assets                                              |       |      |
| 16.5    | Fixing bugs                                                                        |       |      |
| 16.5.1  | URL encoding image paths                                                           |       |      |
| 16.5.2  | Redirecting after image uploads                                                    |       |      |
| 16.6    | Configuring our application                                                        |       |      |
| 16.6.1  | Finding variables that need provided                                               |       |      |
| 16.6.2  | Using functional options                                                           |       |      |
| 16.6.3  | JSON configuration files                                                           |       |      |
| 16.7    | Setting up a server                                                                |       |      |
| 16.7.1  | Digital Ocean droplet                                                              |       |      |
| 16.7.2  | Installing PostgreSQL in production                                                |       |      |
| 16.7.3  | Installing Go                                                                      |       |      |
| 16.7.4  | Setting up Caddy                                                                   |       |      |
| 16.7.5  | Creating a service for our app                                                     |       |      |
| 16.7.6  | Setting up a production config                                                     |       |      |
| 16.7.7  | Creating a deploy script                                                           |       |      |
| 17.1    | Deleting cookies and logging out users                                             |       |      |
| 17.2    | Redirecting with alerts                                                            |       |      |
| 17.3    | Emailing users                                                                     |       |      |
| 17.4    | Persisting and prefilling form data                                                |       |      |
| 17.5    | Resetting passwords                                                                |       |      |
| 17.5.1  | Creating the database model                                                        |       |      |
| 17.5.2  | Updating the services                                                              |       |      |
| 17.5.3  | Forgotten password forms                                                           |       |      |
| 17.5.4  | Controller actions and views                                                       |       |      |
| 17.5.5  | Emailing users and building URLs                                                   |       |      |
| 17.6    | Where do I go next?                                                                |       |      |
| 18.1    | Using interfaces in Go                                                             |       |      |
| 18.2    | Refactoring with gorename                                                          |       |      |
| 18.3    | Handling Bootstrap issues                                                          |       |      |
| 18.4    | Private model errors                                                               |       |      |

## Videos

| Chapter | Title                                                    | Wachted? | Notes |
|---------|----------------------------------------------------------|-------|-------|
| 1       |                                                          | x     |       |
| 2.0     | Creating a code directory and a git repo                 | x     |       |
| 2.1     | A basic web application                                  | x     |       |
| 2.1     | [ASIDE] - What is a web request?                         | x     |       |
| 2.2     | Explaining our web application in detail                 | x     |       |
| 2.3     | Dynamic reloading                                        | x     |       |
| 3.1     | Setting the Content-Type header                          | x     |       |
| 3.2     | Adding a contact page                                    | x     |       |
| 3.3     | Adding a 404 page                                        | x     |       |
| 3.4     | net/http's ServeMux                                      | x     |       |
| 3.5     | github.com/julienschmidt's httprouter                    | x     |       |
| 3.6     | Gorilla Web Toolkit's mux package                        | x     |       |
| 3.7     | Implementing the gorilla/mux router                      | x     |       |
| 3.Ex0   | Exercises overview                                       | x     |       |
| 3.Ex1   | Adding an FAQ page                                       | x     |       |
| 3.Ex2   | Adding a 404 page                                        | x     |       |
| 3.Ex3   | Using httprouter                                         | x     |       |
| 4.0     | What are templates?                                      | x     |       |
| 4.1     | Our first HTML template                                  | x     |       |
| 4.2     | Code injection and contextual encoding                   | x     |       |
| 4.Ex0   | Exercises overview                                       | x     |       |
| 4.Ex1   | Rendering a custom field                                 | x     |       |
| 4.Ex2   | Rendering additional data attributes                     | x     |       |
| 4.Ex3   | Rendering with nested structures                         | x     |       |
| 5.0     | Intro to the MVC videos                                  | x     |       |
| 5.1     | What is MVC?                                             | x     |       |
| 5.2     | Walking through a web request with MVC                   | x     |       |
| 5.Ex0   | Exercises                                                | x     |       |
| 6.0     | Creating our first view                                  | x     |       |
| 6.1     | Creating the contact view                                | x     |       |
| 6.2     | Named and nested templates                               | x     |       |
| 6.3     | Creating the View type                                   | x     |       |
| 6.4     | Using the View type                                      | x     |       |
| 6.5     | Creating a Bootstrap layout                              | x     |       |
| 6.6     | Adding a navigation bar                                  | x     |       |
| 6.7     | Cleaning up our code by globbing template files          | x     |       |
| 6.8     | Simplifying our view rendering logic                     | x     |       |
| 6.9     | Moving our footer to the bootstrap layout                | x     |       |
| 6.10    | Summary                                                  | x     |       |
| 7.0     | Creating the signup page                                 | x     |       |
| 7.1     | Wrapping the signup form in a bootstrap panel            | x     |       |
| 7.2     | Adding a signup link to the navbar                       | x     |       |
| 7.3     | An introduction to REST                                  | x     |       |
| 7.4     | Creating our first controller - the users controller     | x     |       |
| 7.5     | CRUD, HTTP verbs with Gorilla mux, and the create action | x     |       |
| 7.6     | Parsing the signup form (parsing POST forms)             | x     |       |
| 7.7     | Parsing forms with gorilla schema                        | x     |       |
| 7.8     | DRYing up our form parsing code                          | x     |       |
| 7.9     | Creating a controller for our mostly static pages        | x     |       |
| 7.10    | Making views easier to create                            | x     |       |
| 8.0     | What does it mean to persist data?                       | x     |       |
| 8.1     | Web applications use databases to persist data           | x     |       |
| 8.2     | We will be using PostgreSQL                              | x     |       |
| 8.3     | SQL has many great educational resources                 | x     |       |
| 8.4     | Postgres is scalable and relatively easy to use          | x     |       |
| 8.5     | Setting up and connecting to Postgres                    | x     |       |
| 8.6     | SQL basics and creating tables to interact with          | x     |       |
| 8.7     | Connecting to our database with Go's sql package         | x     |       |
| 8.8     | Writing records to our database with Go's sql package    |   x    |     |
| 8.9     | Querying for records with Go's sql package               |   x   |       |
| 8.10    | Handling relational data with Go's sql package           |   x   |       |
| 8.11    | Setting up GORM                                          |   x   |       |
| 8.12    | Creating our first model with GORM                       |    x  |       |
| 8.13    | Creating records and logging with GORM                   |    x  |       |
| 8.14    | Querying records with GORM                               |   x   |       |
| 8.15    | Error handling with GORM                                 |  x    |       |
| 8.16    | Relational data with GORM                                |   x   |       |
| 9.1     | Creating the User model                                  |       |       |
| 9.2     | Creating the UserService                                 |       |       |
| 9.3     | The create user method                                   |       |       |
| 9.4     | What does a model test look like?                        |       |       |
| 9.5     | Finishing the UserService implementation                 |       |       |
| 9.6     | Connecting models and controllers                        |       |       |
| 10.0    | An intro to building an auth system                      |       |       |
| 10.1    | Store hashes, not passwords                              |       |       |
| 10.2    | Implementing bcrypt hashing                              |       |       |
| 10.3    | Using passwords from the signup form                     |       |       |
| 10.4    | Salt and pepper                                          |       |       |
| 10.5    | Creating the login form                                  |       |       |
| 10.6    | Authenticating users                                     |       |       |
| 11.0    | Remembering users and creating our first cookie          |       |       |
| 11.1    | Viewing cookies via code                                 |       |       |
| 11.2    | Creating cookies on login and signup                     |       |       |
| 11.3    | Securing our cookies from tampering                      |       |       |
| 11.4    | Generating remember tokens                               |       |       |
| 11.4    | [ASIDE] - Why 32 bytes?                                  |       |       |
| 11.5    | Writing a remember token hasher                          |       |       |
| 11.6    | Hashing remember tokens on user creation and update      |       |       |
| 11.7    | Storing remember tokens in cookies                       |       |       |
| 11.8    | Securing our cookies from XSS                            |       |       |
| 11.9    | Securing our cookies from theft and CSRF                 |       |       |
| 12.0    | Validating and normalizing                               |       |       |
| 12.1    | Static types vs interfaces                               |       |       |
| 12.1    | [ASIDE] - Emebedding, interfaces, and chaining           |       |       |
| 12.2    | The UserDB interface                                     |       |       |
| 12.3    | The UserService interface                                |       |       |
| 12.4.0  | Organizing validation code                               |       |       |
| 12.4.1  | Remember token normalizer                                |       |       |
| 12.4.2  | Ensuring remember tokens are set on create               |       |       |
| 12.4.3  | Ensuring a valid ID on delete                            |       |       |
| 12.5.0  | Converting emails to lowercase and trimming whitespace   |       |       |
| 12.5.1  | Requiring email addresses                                |       |       |
| 12.5.2  | Verifying emails match a pattern                         |       |       |
| 12.5.3  | Verifying an email address isn't taken                   |       |       |
| 12.6    | Validating and normalizing passwords                     |       |       |
| 12.7    | Validating and normalizing remember tokens               |       |       |
| 13.1    | Bootstrap alerts                                         |       |       |
| 13.2    | Dynamic alerts                                           |       |       |
| 13.3    | Only show alerts when necessary                          |       |       |
| 13.4    | Creating the views.Data type                             |       |       |
| 13.5    | Handling signup errors                                   |       |       |
| 13.6    | Only display public errors                               |       |       |
| 13.7    | Handling login errors                                    |       |       |
| 13.8    | Handling rendering errors                                |       |       |
| 14.0    | Intro to the gallery chapter                             |       |       |
| 14.1    | The gallery model                                        |       |       |
| 14.2    | Sharing a GORM connection and the GalleryService         |       |       |
| 14.3    | Utilizing our shared GORM connection (the Services type) |       |       |
| 14.4    | Moving close to the Services type                        |       |       |
| 14.5    | Implementing the GalleryService                          |       |       |
| 14.6    | Creating the galleries controller                        |       |       |
| 14.7    | Implementing the Gallery create action                   |       |       |
| 14.8    | Gallery validators and normalizers                       |       |       |
| 14.9    | Requiring users to be logged in to view specific pages   |       |       |
| 14.10   | Adding users to the request context                      |       |       |
| 14.11   | Rendering individual galleries                           |       |       |
| 14.12   | Galleries edit action                                    |       |       |
| 14.13   | Galleries update action                                  |       |       |
| 14.14   | GalleryService update method                             |       |       |
| 14.15   | Gallery delete action                                    |       |       |
| 14.16   | Gallery index action                                     |       |       |
| 14.17   | Navbar updates                                           |       |       |
| 15.1    | File upload HTML form                                    |       |       |
| 15.2    | File upload handler                                      |       |       |
| 15.3    | The ImageService and create method                       |       |       |
| 15.4    | Finding images by Gallery ID                             |       |       |
| 15.5    | Rendering images                                         |       |       |
| 15.6    | Deleting images                                          |       |       |
| 16.1    | Error cleanup                                            |       |       |
| 16.2    | Serving static assets                                    |       |       |
| 16.3    | CSRF protection                                          |       |       |
| 16.4    | Don't lookup the user when serving static assets         |       |       |
| 16.5    | URL encoding image path bug                              |       |       |
| 16.6    | Starting with config variables                           |       |       |
| 16.7    | Functional options for services                          |       |       |
| 16.8    | Loading a JSON config                                    |       |       |
| 16.9    | Setting up a droplet (server) on Digital Ocean           |       |       |
| 16.10   | Installing postgres on our prod server                   |       |       |
| 16.11   | Caddy server                                             |       |       |
| 16.12   | Deploy script                                            |       |       |
| 17.1    | Deleting cookies and logging out users                   |       |       |
| 17.2    | Redirecting with alerts                                  |       |       |
| 17.3    | Emailing users                                           |       |       |
| 17.4    | Persisting form data                                     |       |       |
| 17.5.1  | Resetting passwords - Creating the pwReset DB table      |       |       |
| 17.5.2  | Resetting passwords - Implementing UserService methods   |       |       |
| 17.5.3  | Resetting passwords - Form HTML templates                |       |       |
| 17.5.4  | Resetting passwords - Controller updates                 |       |       |
| 17.5.5  | Resetting passwords - Sending emails and building links  |       |       |


## Done Log

* Connect go with psql server from laptop and computer
* Setup postgres on my home server
    * Open port for 5432
    * Add listen_adress '*'
    * Add ip range
