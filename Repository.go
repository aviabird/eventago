package eventago

import (
	"database/sql"
	"fmt"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "123456"
	dbName := "golang"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// createOrderQueryModel Insert event related information to database.
func createOrderQueryModel(eid string, event string, eventDetails string, version int) bool {
	db := dbConn()

	insForm, err := db.Prepare("INSERT INTO eventdetails(eid,event,version,details) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("INSERT INTO eventdetails(eid,event,version,details) VALUES('2f0d0cfe-91b1-471a-a979-6de4bd39ab4c', 'amount credited.','100 rs amout created by  2f0d0cfe-91b1-471a-a979-6de4bd39ab4c',	'0')")
	insForm.Exec(eid, event, version, eventDetails)
	fmt.Println(insForm)
	defer db.Close()
	return true
}

// FindEntryQueryModel find event related information to database.
func FindEntryQueryModel(eid string) *EventStream {
	db := dbConn()

	selDB, err := db.Query("SELECT eid,event,version,details  FROM eventdetails WHERE eid=?", eid)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	emp := EventStream{}
	for selDB.Next() {
		var eid, eventDetails string
		var event []string
		var version int
		err = selDB.Scan(&eid, &event, &eventDetails, &version)
		if err != nil {
			panic(err.Error())
		}
		emp.eid = eid
		emp.event = event
		emp.eventDetails = eventDetails
		emp.version = version
	}
	return &emp
}
