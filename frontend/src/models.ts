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
	id?: string;
	name: string;
	picture: string;
	colorPalette?: string;
	habits: Habit[];
	friends: User[];
}

export interface Quote {
	quote: string;
	source: string;
}
