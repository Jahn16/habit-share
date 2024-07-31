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
	ID?: string;
	name: string;
	picture: string;
	colorPalette?: string;
	habits: Habit[];
}

export interface Quote {
	quote: string;
	source: string;
}
