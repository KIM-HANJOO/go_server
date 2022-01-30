package main
/*
import 4 important libraries
1. "net/http" to access the core go http functionality
2. "fmt" for formatting text
3. "html/template" is a library that allows us to interact with our html file.
4. "time" for working with date and time
*/


import(
    "net/http"
    "fmt"
    "time"
    "html/template"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct{
    Name string
    Time string
}

//Go application entrypoint
func main(){
    // Instantiate a Welcome struct object and pass in some random information
    // We shall get the name of the user as a query parameter from the URL

    welcome := Welcome{"Kim-hanjoo", time.Now().Format(time.Stamp)}

    /* We tell Go exactly where we can find our html file. We ask go to parse the html file 
    (Notice the relative path). We wrap it in a call to template.Must()
    wich handles any errors and halts if there are fatal errors*/

    templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

    /* Our html comes with CSS that go needs to provide when we run the app. Here we tell go to create
    a handle that looks in the static directory, go then uses the "/static/" as a url that out 
    html can refer to when looking for our css and other files*/

    http.Handle("/static/", //final url can be anything
    http.StripPrefix("/static/",
    http.FileServer(http.Dir("static"))))

    /*Go looks in the relative "static" directory first using http.FileServer(),
    then matches it to a url of our choice as shown in http.Handle("/static/").
    This url is what we need when referencing our css files once the server begins.
    HTML code would therefore be <link rel = "stylesheet" href = "/static/stylesheets/...">
    It is important to note the url in http.Handle can be whatever we like, so long as we are consistent.*/

    // This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        // Takes the name from the URL query e.g ?name=Martin, will set welcome.Name = Martin
        if name := r.FormValue("neme"); name != "" {
            welcome.Name = name ;
        }
        // if errors show an internal server error message
        // I also pass the welcome struct to the welcome-template.html file.

        if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    //Start the web server, set the port to listen to 8080. Without a path it assumes localhost
    fmt.Println("Listening");
    fmt.Println(http.ListenAndServe(":8080", nil));
}

