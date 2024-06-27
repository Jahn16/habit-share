export interface HabitRecord {
	habitID: number;
	date: string;
}

export interface Habit {
	ID?: number;
	name: string;
	goal: number;
	records: HabitRecord[];
}

export interface User {
	name: string;
	picture: string;
	habits: Habit[];
}
