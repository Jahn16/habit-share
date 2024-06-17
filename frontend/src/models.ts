export interface HabitRecord {
	date: string;
}

export interface Habit {
	name: string;
	goal: number;
	records: HabitRecord[];
}
