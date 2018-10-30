package main

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "os"
        "time"
        "golang.org/x/net/context"
        "golang.org/x/oauth2"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/classroom/v1"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "github.com/gorilla/mux"
)

type Course struct {
  ID bson.ObjectId `bson:"_id" json:"id"`
  CoverImage string `bson:"cover_image" json:"coverImage"`
  Name            string    `bson:"name" json:"name"`
	Description      string    `bson:"description" json:"description"`
	CreatedAt        time.Time `json:"createdAt" bson:"created_at"`
  CourseId             string    `json:"courseId" bson:"course_id"`
  TeacherName     string      `json:"teacherName" bson:"teacher_name"`
  SchoolName      string      `json:"schoolName" bson:"school_name"`
}

type Rate struct {
  ID bson.ObjectId `bson:"_id" json:"id"`
  Title            string    `bson:"title" json:"title"`
	Description      string    `bson:"description" json:"description"`
	CreatedAt        time.Time `json:"createdAt" bson:"created_at"`
  Rate     int      `json:"rate" bson:"rate"`
  CourseId bson.ObjectId `bson:"_id" json:"id"`
  //agregar el te resultó útil, responder, comentar un abuso
}

type PayloadData struct {
	Text      string    `json:"text" bson:"text"`
}

type Post struct {
	Text      string    `json:"text" bson:"text"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
}

type Token struct {
	source      string    `json:"text" bson:"text"`
	createdAt time.Time `json:"createdAt" bson:"created_at"`
  access_token string `json:"text" bson:"text"`
  token_type string `json:"text" bson:"text"`
  refresh_token string `json:"text" bson:"text"`
  expiry time.Time `json:"createdAt" bson:"created_at"`
}

var posts *mgo.Collection
var courses *mgo.Collection
var rates *mgo.Collection

func FindCourses(w http.ResponseWriter, r *http.Request) {
  result := []Course{}
	if err := courses.Find(nil).Sort("-created_at").All(&result); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
	} else {
		responseJSON(w, result)
	}
}

func FindCourse(w http.ResponseWriter, r *http.Request) {
 vars := mux.Vars(r)
 id := vars["id"] // param id
 var course Course
 if !bson.IsObjectIdHex(id) {
	responseError(w, "", http.StatusNotFound)
	return
 }
 // Grab id
 oid := bson.ObjectIdHex(id)
 if err := courses.Find(bson.M{"_id": oid}).One(&course); err != nil {
	 responseError(w, err.Error(), http.StatusInternalServerError)
	 return
 }
 responseJSON(w, course)
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	course := &Course{}
	err = json.Unmarshal(data, course)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	course.CreatedAt = time.Now().UTC()
  course.ID = bson.NewObjectId()

	if err := courses.Insert(course); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responseJSON(w, course)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
	var course Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
    responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
  if err := courses.UpdateId(course.ID, &course); err != nil {
    responseError(w, err.Error(), http.StatusInternalServerError)
		return
  }
	responseJSON(w, course)
}

//Return something
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"] // param id
		// Verify id is ObjectId, otherwise bail
	 if !bson.IsObjectIdHex(id) {
		responseError(w, "", http.StatusNotFound)
	  return
	 }
	 // Grab id
	 oid := bson.ObjectIdHex(id)
	 // Remove
	 if err := courses.RemoveId(oid); err != nil {
	  responseError(w, err.Error(), http.StatusInternalServerError)
	  return
	 }
	// Write status
	 w.WriteHeader(http.StatusOK)
	 return
}



func main() {
    // Connect to mongo
  	session, err := mgo.Dial("mongo:27017")
  	if err != nil {
  		log.Fatalln(err)
  		log.Fatalln("mongo err")
  		os.Exit(1)
  	}
  	defer session.Close()
  	session.SetMode(mgo.Monotonic, true)

    // Get posts collection
  	posts = session.DB("app").C("posts")
    courses = session.DB("app").C("courses")

    r := mux.NewRouter()
    r.HandleFunc("/courses", FindCourses).Methods("GET")
    r.HandleFunc("/courses", CreateCourse).Methods("POST")
    r.HandleFunc("/courses", UpdateCourse).Methods("PUT")
    r.HandleFunc("/courses/{id}", DeleteCourse).Methods("DELETE")
    r.HandleFunc("/courses/{id}", FindCourse).Methods("GET")

    r.HandleFunc("/posts", createPost).
  		Methods("POST")

  	r.HandleFunc("/posts", readPosts).
  		Methods("GET")

    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}

//Posts
func createPost(w http.ResponseWriter, r *http.Request) {
	// Read body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Read post
	post := &Post{}
	err = json.Unmarshal(data, post)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.CreatedAt = time.Now().UTC()

	// Insert new post
	if err := posts.Insert(post); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON(w, post)
}

func readPosts(w http.ResponseWriter, r *http.Request) {
	result := []Post{}
	if err := posts.Find(nil).Sort("created_at").All(&result); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
	} else {
		responseJSON(w, result)
	}
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
        tokFile := "token.json"
        tok, err := tokenFromFile(tokFile)
        if err != nil {
                tok = getTokenFromWeb(config)
                saveToken(tokFile, tok)
        }
        return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
        authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
        fmt.Printf("Go to the following link in your browser then type the "+
                "authorization code: \n%v\n", authURL)

        var authCode string
        if _, err := fmt.Scan(&authCode); err != nil {
                log.Fatalf("Unable to read authorization code: %v", err)
        }

        tok, err := config.Exchange(oauth2.NoContext, authCode)
        if err != nil {
                log.Fatalf("Unable to retrieve token from web: %v", err)
        }
        return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
        f, err := os.Open(file)
        defer f.Close()
        if err != nil {
                return nil, err
        }
        tok := &oauth2.Token{}
        err = json.NewDecoder(f).Decode(tok)
        return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
        fmt.Printf("Saving credential file to: %s\n", path)
        f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
        defer f.Close()
        if err != nil {
                log.Fatalf("Unable to cache oauth token: %v", err)
        }
        json.NewEncoder(f).Encode(token)
}

func classroomConnection() {
        b, err := ioutil.ReadFile("credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, classroom.ClassroomCoursesReadonlyScope)
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        client := getClient(config)

        srv, err := classroom.New(client)
        if err != nil {
                log.Fatalf("Unable to create classroom Client %v", err)
        }

        r, err := srv.Courses.List().PageSize(10).Do()
        if err != nil {
                log.Fatalf("Unable to retrieve courses. %v", err)
        }
        if len(r.Courses) > 0 {
                fmt.Print("Courses:\n")
                for _, c := range r.Courses {
                        fmt.Printf("%s (%s)\n", c.Name, c.Id)
                }
        } else {
                fmt.Print("No courses found.")
        }
}

func responseError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
