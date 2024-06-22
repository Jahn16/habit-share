export interface HabitRecord {
	date: string;
}

export interface Habit {
	name: string;
	goal: number;
	records: HabitRecord[];
}

export interface User {
	name: string;
	picture: string;
	habits: Habit[];
}
