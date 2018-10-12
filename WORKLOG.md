# WORKLOG

## Chapters


| Chapter | Title                                                    | Book | Video |
|---------|----------------------------------------------------------|------|-------|
| 1       |                                                          |x     |x      |
| 2.0     | Creating a code directory and a git repo                 |x     |x      |
| 2.1     | A basic web application                                  |x     |x      |
| 2.1     | [ASIDE] - What is a web request?                         |x     |x      |
| 2.2     | Explaining our web application in detail                 |x     |x      |
| 2.3     | Dynamic reloading                                        |x     |x      |
| 3.1     | Setting the Content-Type header                          |x     |x      |
| 3.2     | Adding a contact page                                    |x     |x      |
| 3.3     | Adding a 404 page                                        |x     |x      |
| 3.4     | net/http's ServeMux                                      |x     |x      |
| 3.5     | github.com/julienschmidt's httprouter                    |x     |x      |
| 3.6     | Gorilla Web Toolkit's mux package                        |x     |x      |
| 3.7     | Implementing the gorilla/mux router                      |x     |x      |
| 3.Ex0   | Exercises overview                                       |x     |x      |
| 3.Ex1   | Adding an FAQ page                                       |x     |x      |
| 3.Ex2   | Adding a 404 page                                        |x     |x      |
| 3.Ex3   | Using httprouter                                         |x     |x      |
| 4.0     | What are templates?                                      |x     |x      |
| 4.1     | Our first HTML template                                  |x     |x      |
| 4.2     | Code injection and contextual encoding                   |x     |x      |
| 4.Ex0   | Exercises overview                                       |x     |x      |
| 4.Ex1   | Rendering a custom field                                 |x     |x      |
| 4.Ex2   | Rendering additional data attributes                     |x     |x      |
| 4.Ex3   | Rendering with nested structures                         |x     |x      |
| 5.0     | Intro to the MVC videos                                  |x     |x      |
| 5.1     | What is MVC?                                             |x     |x      |
| 5.2     | Walking through a web request with MVC                   |x     |x      |
| 5.Ex0   | Exercises                                                |x     |x      |
| 6.0     | Creating our first view                                  |x     |x      |
| 6.1     | Creating the contact view                                |x     |x      |
| 6.2     | Named and nested templates                               |x     |x      |
| 6.3     | Creating the View type                                   |x     |x      |
| 6.4     | Using the View type                                      |x     |x      |
| 6.5     | Creating a Bootstrap layout                              |x     |x      |
| 6.6     | Adding a navigation bar                                  |x     |x      |
| 6.7     | Cleaning up our code by globbing template files          |x     |x      |
| 6.8     | Simplifying our view rendering logic                     |x     |x      |
| 6.9     | Moving our footer to the bootstrap layout                |x     |x      |
| 6.10    | Summary                                                  |x     |x      |
| 7.0     | Creating the signup page                                 |x     |x      |
| 7.1     | Wrapping the signup form in a bootstrap panel            |x     |x      |
| 7.2     | Adding a signup link to the navbar                       |x     |x      |
| 7.3     | An introduction to REST                                  |x     |x      |
| 7.4     | Creating our first controller - the users controller     |x     |x      |
| 7.5     | CRUD, HTTP verbs with Gorilla mux, and the create action |x     |x      |
| 7.6     | Parsing the signup form (parsing POST forms)             |x     |x      |
| 7.7     | Parsing forms with gorilla schema                        |x     |x      |
| 7.8     | DRYing up our form parsing code                          |x     |x      |
| 7.9     | Creating a controller for our mostly static pages        |x     |x      |
| 7.10    | Making views easier to create                            |x     |x      |
| 8.0     | What does it mean to persist data?                       |x     |x      |
| 8.1     | Web applications use databases to persist data           |x     |x      |
| 8.2     | We will be using PostgreSQL                              |x     |x      |
| 8.3     | SQL has many great educational resources                 |x     |x      |
| 8.4     | Postgres is scalable and relatively easy to use          |x     |x      |
| 8.5     | Setting up and connecting to Postgres                    |x     |x      |
| 8.6     | SQL basics and creating tables to interact with          |x     |x      |
| 8.7     | Connecting to our database with Go's sql package         |x     |x      |
| 8.8     | Writing records to our database with Go's sql package    |      |       |
| 8.9     | Querying for records with Go's sql package               |      |       |
| 8.10    | Handling relational data with Go's sql package           |      |       |
| 8.11    | Setting up GORM                                          |      |       |
| 8.12    | Creating our first model with GORM                       |      |       |
| 8.13    | Creating records and logging with GORM                   |      |       |
| 8.14    | Querying records with GORM                               |      |       |
| 8.15    | Error handling with GORM                                 |      |       |
| 8.16    | Relational data with GORM                                |      |       |
| 9.1     | Creating the User model                                  |      |       |
| 9.2     | Creating the UserService                                 |      |       |
| 9.3     | The create user method                                   |      |       |
| 9.4     | What does a model test look like?                        |      |       |
| 9.5     | Finishing the UserService implementation                 |      |       |
| 9.6     | Connecting models and controllers                        |      |       |
| 10.0    | An intro to building an auth system                      |      |       |
| 10.1    | Store hashes, not passwords                              |      |       |
| 10.2    | Implementing bcrypt hashing                              |      |       |
| 10.3    | Using passwords from the signup form                     |      |       |
| 10.4    | Salt and pepper                                          |      |       |
| 10.5    | Creating the login form                                  |      |       |
| 10.6    | Authenticating users                                     |      |       |
| 11.0    | Remembering users and creating our first cookie          |      |       |
| 11.1    | Viewing cookies via code                                 |      |       |
| 11.2    | Creating cookies on login and signup                     |      |       |
| 11.3    | Securing our cookies from tampering                      |      |       |
| 11.4    | Generating remember tokens                               |      |       |
| 11.4    | [ASIDE] - Why 32 bytes?                                  |      |       |
| 11.5    | Writing a remember token hasher                          |      |       |
| 11.6    | Hashing remember tokens on user creation and update      |      |       |
| 11.7    | Storing remember tokens in cookies                       |      |       |
| 11.8    | Securing our cookies from XSS                            |      |       |
| 11.9    | Securing our cookies from theft and CSRF                 |      |       |
| 12.0    | Validating and normalizing                               |      |       |
| 12.1    | Static types vs interfaces                               |      |       |
| 12.1    | [ASIDE] - Emebedding, interfaces, and chaining           |      |       |
| 12.2    | The UserDB interface                                     |      |       |
| 12.3    | The UserService interface                                |      |       |
| 12.4.0  | Organizing validation code                               |      |       |
| 12.4.1  | Remember token normalizer                                |      |       |
| 12.4.2  | Ensuring remember tokens are set on create               |      |       |
| 12.4.3  | Ensuring a valid ID on delete                            |      |       |
| 12.5.0  | Converting emails to lowercase and trimming whitespace   |      |       |
| 12.5.1  | Requiring email addresses                                |      |       |
| 12.5.2  | Verifying emails match a pattern                         |      |       |
| 12.5.3  | Verifying an email address isn't taken                   |      |       |
| 12.6    | Validating and normalizing passwords                     |      |       |
| 12.7    | Validating and normalizing remember tokens               |      |       |
| 13.1    | Bootstrap alerts                                         |      |       |
| 13.2    | Dynamic alerts                                           |      |       |
| 13.3    | Only show alerts when necessary                          |      |       |
| 13.4    | Creating the views.Data type                             |      |       |
| 13.5    | Handling signup errors                                   |      |       |
| 13.6    | Only display public errors                               |      |       |
| 13.7    | Handling login errors                                    |      |       |
| 13.8    | Handling rendering errors                                |      |       |
| 14.0    | Intro to the gallery chapter                             |      |       |
| 14.1    | The gallery model                                        |      |       |
| 14.2    | Sharing a GORM connection and the GalleryService         |      |       |
| 14.3    | Utilizing our shared GORM connection (the Services type) |      |       |
| 14.4    | Moving close to the Services type                        |      |       |
| 14.5    | Implementing the GalleryService                          |      |       |
| 14.6    | Creating the galleries controller                        |      |       |
| 14.7    | Implementing the Gallery create action                   |      |       |
| 14.8    | Gallery validators and normalizers                       |      |       |
| 14.9    | Requiring users to be logged in to view specific pages   |      |       |
| 14.10   | Adding users to the request context                      |      |       |
| 14.11   | Rendering individual galleries                           |      |       |
| 14.12   | Galleries edit action                                    |      |       |
| 14.13   | Galleries update action                                  |      |       |
| 14.14   | GalleryService update method                             |      |       |
| 14.15   | Gallery delete action                                    |      |       |
| 14.16   | Gallery index action                                     |      |       |
| 14.17   | Navbar updates                                           |      |       |
| 15.1    | File upload HTML form                                    |      |       |
| 15.2    | File upload handler                                      |      |       |
| 15.3    | The ImageService and create method                       |      |       |
| 15.4    | Finding images by Gallery ID                             |      |       |
| 15.5    | Rendering images                                         |      |       |
| 15.6    | Deleting images                                          |      |       |
| 16.1    | Error cleanup                                            |      |       |
| 16.2    | Serving static assets                                    |      |       |
| 16.3    | CSRF protection                                          |      |       |
| 16.4    | Don't lookup the user when serving static assets         |      |       |
| 16.5    | URL encoding image path bug                              |      |       |
| 16.6    | Starting with config variables                           |      |       |
| 16.7    | Functional options for services                          |      |       |
| 16.8    | Loading a JSON config                                    |      |       |
| 16.9    | Setting up a droplet (server) on Digital Ocean           |      |       |
| 16.10   | Installing postgres on our prod server                   |      |       |
| 16.11   | Caddy server                                             |      |       |
| 16.12   | Deploy script                                            |      |       |
| 17.1    | Deleting cookies and logging out users                   |      |       |
| 17.2    | Redirecting with alerts                                  |      |       |
| 17.3    | Emailing users                                           |      |       |
| 17.4    | Persisting form data                                     |      |       |
| 17.5.1  | Resetting passwords - Creating the pwReset DB table      |      |       |
| 17.5.2  | Resetting passwords - Implementing UserService methods   |      |       |
| 17.5.3  | Resetting passwords - Form HTML templates                |      |       |
| 17.5.4  | Resetting passwords - Controller updates                 |      |       |
| 17.5.5  | Resetting passwords - Sending emails and building links  |      |       |

## Done Log

* Connect go with psql server from laptop and computer
* Setup postgres on my home server
    * Open port for 5432
    * Add listen_adress '*'
    * Add ip range
