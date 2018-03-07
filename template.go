package main

import (  
    "time"
    "html/template"
    "os"
    "fmt"
)

type Account struct {  
    FirstName string
    LastName  string
}



type Statement struct {  
    FromDate  time.Time
        Account   Account

}

func main() {  
    fmap := template.FuncMap{
        "formatAsDate": formatAsDate,
    }
    t := template.Must(template.New("email.tmpl").Funcs(fmap).ParseFiles("email.tmpl"))
    err := t.Execute(os.Stdout, createMockStatement())
    if err != nil {
        panic(err)
    }
}

func formatAsDate(t time.Time) string {  
    year, month, day := t.Date()
    return fmt.Sprintf("%d/%d/%d", day, month, year)
}

func createMockStatement() Statement {  
    return Statement{
        FromDate: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
        Account: Account{
            FirstName: "Client1",
            LastName: "surname1",
        },
        
    }
}