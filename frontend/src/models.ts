export interface HabitRecord {
	ID?: number;
	habitID: number;
	date: string;
}

export interface Habit {
	ID?: number;
	icon: string;
	name: string;
	goal: number;
	records: HabitRecord[];
}

export interface User {
	name: string;
	picture: string;
	habits: Habit[];
}

export interface Quote {
	quote: string;
	source: string;
}
