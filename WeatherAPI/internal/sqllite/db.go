package sqllite

import (
	"WeatherAPI/internal/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

//Package to handle caching of prev 2 mins data
type DataAccessObject struct {
	db *sql.DB
}

type Dbops interface {
	CreateSchema() error
	InsertData(string, string, model.Response) error
	DeleteData() error
	GetData(string, string) ([]byte, error)
	CloseConnection()
}

var Instance *DataAccessObject

func Initdatabase() (err error, dbfunc Dbops) {
	if Instance == nil {

		database, err := sql.Open("sqlite3", "./wether.db")
		if err != nil {
			return err, nil
		}
		if database == nil {
			return errors.New("failed to establish db conection"), nil
		}
		//DAO.db = database
		Instance = &DataAccessObject{
			db: database,
		}
	}
	return nil, Instance

}

func (DAO *DataAccessObject) CloseConnection() {
	DAO.db.Close()
	time.Sleep(35 * time.Second)
}
func (DAO *DataAccessObject) CreateSchema() error {
	statement, err := DAO.db.Prepare("CREATE TABLE IF NOT EXISTS cache_data  (city varchar(255), country varchar(255),  response blob, created_time timestamp default current_timestamp, Primary Key(city, country) )")
	if err != nil {
		fmt.Println(err)
		return err
	}
	result, err := statement.Exec()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result)
	return nil

}

func (DAO *DataAccessObject) InsertData(city, country string, response model.Response) error {
	statement, err := DAO.db.Prepare("INSERT INTO cache_data(city, country, response ) VALUES ( ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	out, _ := json.Marshal(response)
	_, err = statement.Exec(city, country, out)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (DAO *DataAccessObject) DeleteData() error {
	fmt.Println("DeleteData", time.Now())
	statement, err := DAO.db.Prepare("DELETE FROM cache_data WHERE created_time <= datetime('now','-2 minutes')")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (DAO *DataAccessObject) GetData(city, country string) (Response []byte, err error) {
	row := DAO.db.QueryRow("SELECT response  from cache_data where city=? and country=?", city, country)
	err1 := row.Scan(&Response)
	switch {
	case err1 == sql.ErrNoRows:
		return []byte{}, errors.Wrapf(err1, "no data with this city as %s and country as %s\n", city, country)
	case err1 != nil:
		return []byte{}, errors.Wrap(err1, "query error: ")
	}
	return Response, nil
}
