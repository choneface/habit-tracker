package main

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

func GetHabits() []Habit {
	return habits 
}

func SaveHabit(h Habit) {
	habits = append(habits, h)
}

func SaveNewHabit(t string, d string) {
	h := Habit {
		Title: t,
		Description: d,
		History: generateHistory(),
	}
	habits = append(habits, h)
}

func generateHistory() []byte {
	return make([]byte, 15)
}
