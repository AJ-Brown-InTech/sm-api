// postgres database
// todo: migrations and hosting(possible contanerization)
//todo : add a config yaml file
package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
	"github.com/jmoiron/sqlx"
	
)

func DatabaseInstance(hostname, port, user, password, dbname string)(*sqlx.DB, error){


connectionStr := fmt.Sprintf(`port=%s user=%s password=%s dbname=%s host=%s sslmode=disable`, port, user, password, dbname, hostname)

	db, err := sqlx.Open("postgres", connectionStr)
	if err != nil{
		logrus.Errorf("Failed to establish database connection: %v", err)
	 	return nil,err
	}
	
	db.SetMaxOpenConns(100) //not sure how many but from what i read 100 is max before performance becomes an issue
	db.SetMaxIdleConns(30) //idle just added a few incase some connections are hung up
	db.SetConnMaxLifetime(30 * time.Second)
	db.SetConnMaxIdleTime(30 * time.Second)


	 if err := db.Ping(); err != nil{
		logrus.Errorf("Failed to ping db: %v", err)
	 	return nil,err
	 }

	logrus.Info("Postgres database connection established") 
	return db, nil
}