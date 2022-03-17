package MeetingScheduler

var events []EventObject = []EventObject{
	{
		"1",
		"Daily Stand-up",
		[]string{"1", "2", "3", "4"},
		2,
		3,
		"Discuss daily tasks",
		INVITE_ONLY,
		"2",
		"ACTIVE",
	},
	{
		"2",
		"Analytics meet",
		[]string{},
		4,
		5,
		"Seminar on data analytics",
		OPEN_FOR_ALL,
		"ADMIN",
		"ACTIVE",
	},
}

var users []User = []User{
	{
		"1",
		"Vinay",
		"SE",
	}, {
		"2",
		"Saurabh",
		"Manager",
	}, {
		"3",
		"Naman",
		"SE",
	}, {
		"4",
		"Subharaj",
		"SE",
	},
}
