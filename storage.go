package main

var habits = []habit {
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

func getHabits() []habit {
	return habits 
}

func saveHabit(h habit) {
	habits = append(habits, h)
}

func saveNewHabit(t string, d string) {
	h := habit {
		Title: t,
		Description: d,
		History: _generateHistory(),
	}
	habits = append(habits, h)
}

func _generateHistory() []byte {
	return make([]byte, 15)
}
