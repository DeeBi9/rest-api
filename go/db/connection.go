package db

import (
	"database/sql"
	"log"

	"github.com/Deepanshuisjod/rest-api/models"
	_ "github.com/lib/pq"
)

func connect() (*sql.DB, error) {
	conn := "user=postgres password=123456 dbname=taskapi sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to the database")
	return db, err
}

func GetEmpInformation() []*models.Information {
	var InformationList []*models.Information

	db, err := connect()
	if err != nil {
		log.Fatal("[Error] while connecting to the database:", err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, name, designation FROM employees;`)
	if err != nil {
		log.Println("[Error] running query:", err)
		return InformationList
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var designation string

		err = rows.Scan(&id, &name, &designation)
		if err != nil {
			log.Println("[Error] scanning row:", err)
			continue
		}

		InformationList = append(InformationList, &models.Information{
			Post: designation,
			Name: name,
			Id:   id,
		})
	}

	if err = rows.Err(); err != nil {
		log.Println("[Error] iterating rows:", err)
	}

	return InformationList
}

func SetProjectInformation(emp *models.Assignees) error {
	db, err := connect()
	if err != nil {
		log.Fatal("Error while connecting to the database:", err)
	}
	defer db.Close()

	setQuery := `INSERT INTO employees (id,name,designation,projectname,projectid)
	VALUES ($1,$2,$3,$4,$5)`

	rows, err := db.Query(setQuery,
		emp.Information.Id,
		emp.Information.Name,
		emp.Information.Post,
		emp.Project.ProjectName,
		emp.Project.ProjectId)

	if err != nil {
		log.Println("Error running query:", err)
		return err
	}

	defer rows.Close()

	return nil
}

func CheckID(id int) bool {
	db, err := connect()
	if err != nil {
		log.Println("[Error] while connecting to the database:", err)
		return false
	}
	defer db.Close()

	// Query to check if the ID exists
	var exists bool
	err = db.QueryRow(`SELECT EXISTS(SELECT 1 FROM employees WHERE id = $1)`, id).Scan(&exists)
	if err != nil {
		log.Println("[Error] running query", err)
		return false
	}

	return exists
}
