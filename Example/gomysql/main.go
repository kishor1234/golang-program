package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	// mysql pkg
	_ "github.com/go-sql-driver/mysql"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/rs/cors"
)

var (
	dbhostsip  = "127.0.0.1:3306" //Mysql Server Ip address
	dbusername = "root"           //Database user name
	dbpassword = "root@123"       //Data Base Password
	dbname     = "pest"           //Databasename
	dbcharset  = "utf8"
)

func getConnString() string {
	return dbusername + ":" + dbpassword + "@tcp(" + dbhostsip + ")/" + dbname + "?charset=" + dbcharset
}

func getJSON(sqlString string, taskID string) (string, error) {
	sqlConnString := getConnString()
	db, err := sql.Open("mysql", sqlConnString)
	if err != nil {
		return "", err
	}

	defer db.Close()
	rows, err := db.Query(sqlString, taskID)
	if err != nil {
		return "", err
	}

	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
			str := fmt.Sprintf("%v", v)
			if col == "title" {
				res1 := strings.Replace(str, " ", "-", -1)

				res2 := strings.ToLower(res1)
				entry["urldata"] = res2
			}
			if col == "description" {
				stripped := strip.StripTags(str)

				entry["description"] = stripped[0:100]
			}
			//fmt.Println(col)
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func getJSONSingleBlog(sqlString string, taskID string) (string, error) {
	sqlConnString := getConnString()
	db, err := sql.Open("mysql", sqlConnString)
	if err != nil {
		return "", err
	}

	defer db.Close()
	rows, err := db.Query(sqlString, taskID)
	if err != nil {
		return "", err
	}

	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v

		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

/*func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wel come Go Lang form aasksoft!")

}
func login_handler(w http.ResponseWriter, r *http.Request) {

}
func blogSnap_handler(w http.ResponseWriter, r *http.Request) {
	var taskID = "Blog"
	sql := "select * from post where category = ?"
	rows, err := getJSON(sql, taskID)
	fmt.Fprintf(w, rows)
	fmt.Println(err)
}*/
type titlestruct struct {
	Title string
}

func main() {
	//http.HandleFunc("/", index_handler)
	//http.HandleFunc("/login", login_handler)
	//http.HandleFunc("/blog", blogSnap_handler)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		var taskID = "Blog"
		sql := "select * from post where category = ?"
		rows, err := getJSON(sql, taskID)
		//fmt.Fprintf(w, rows)
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(rows))
	})
	//for single blog request
	mux.HandleFunc("/singleBlog", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var t titlestruct
		err = json.Unmarshal(body, &t)
		if err != nil {
			panic(err)
		}
		title := strings.Replace(t.Title, "-", " ", -1)
		sql := "select * from post where title = ?"
		rows, err := getJSONSingleBlog(sql, title)
		//fmt.Fprintf(w, rows)
		fmt.Println(err)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(rows))

	})
	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)

}
