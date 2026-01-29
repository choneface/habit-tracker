package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

var habits = []Habit {
	{ 
		"Chinese Lessons", 
		"Need passable chinese before having a kid",
		[]byte{1,0,1,1,0,1,0,1,0,1,1,1,0,1,0},
	}, {
		"Spanish Lessons", 
		"Need passable spanish before having a kid",
		[]byte{0,1,0,1,0,1,1,1,0,1,1,0,0,0,1},
	},
}

func Init() Storage {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		db.Close()
		log.Fatal("Failed to open DB, closing")
	}

	createHabitTable := `
	create table if not exists habit (
		id integer not null primary key,
		title text not null,
		description text,
		color text not null
	);
	`
	_, err = db.Exec(createHabitTable)
	if err != nil {
		db.Close()
		log.Fatal("Failed to create habit table, closing")
	}

	createHistoryTable := `
	create table if not exists history (
		id integer not null primary key,
		habit_id integer not null references habit(id),
		date text not null,
		value integer not null
	);
	`
	_, err = db.Exec(createHistoryTable)
	if err != nil {
		db.Close()
		log.Fatal("Failed to create history table, closing")
	}

	return Storage{
		db: db,
	}
}

func (s Storage) GetHabits() []Habit {
	ret := []Habit{}

	cutoffTime := time.Now().AddDate(0, -1, 0)
	rows := s.fetchHabitsWithHistory(cutoffTime)
	count := 0
	for rows.Next() {
		count = count + 1
		var id int
		var title string 
		var description string
		var color string 
		var date sql.NullString
		var value sql.NullByte 
		err := rows.Scan(&id, &title, &description, &color, &date, &value)
		if err != nil {
			log.Fatalf("Failed to scan row \n %s", err.Error())
		}
		tmp := Habit {
			Title: title,
			Description: description,
			History: generateHistory(),
		}
		ret = append(ret, tmp)
	}
	return ret 
}

func (s Storage) SaveNewHabit(t string, d string) {
	s.insertHabit(t, d, "#6ceb8e")
}

func (s Storage) fetchHabitsWithHistory(cutoff time.Time) *sql.Rows {
	raw := `
	select 
		habit.id, 
		habit.title, 
		habit.description, 
		habit.color,
		history.date,
		history.value
	from habit
	left join history on habit.id = history.habit_id
	where (history.date is null or history.date != ?)
	order by date desc;
	`
	statement, err := s.db.Prepare(raw)
	if err != nil {
		log.Fatalf("Failed to prepare habits query \n%s", err.Error()) 
	}
	cutoffString := cutoff.Format("2006-01-02")
	rows, err := statement.Query(cutoffString)
	if err != nil {
		log.Fatalf("Failed to fetch habits from DB\n%s", err.Error()) 
	}
	return rows 
}

func (s Storage) insertHabit(t string, d string, c string) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatalf("Failed to start transaction on the DB\n%s", err.Error()) 
	}
	stmnt, err := tx.Prepare(`
	insert into habit (title, description, color) 
	values (?, ?, ?);
	`)
	if err != nil {
		log.Fatalf("Failed to prepare insert statement\n%s", err.Error()) 
	}
	defer stmnt.Close()
	_, err = stmnt.Exec(t, d, c)
	if err != nil {
		log.Fatalf("Failed to execute insert statement\n%s", err.Error()) 
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalf("Failed to commit transaction\n%s", err.Error()) 
	}
}

func generateHistory() []byte {
	return make([]byte, 15)
}
