import type { User, Habit, HabitRecord } from '../../models';
export class SocialHabitsClient {
	private url: string;
	constructor() {
		this.url = 'http://localhost:8000';
	}

	async getUser(id: number): Promise<User> {
		const url = `${this.url}/users/${id}`;
		const response = await fetch(url);
		if (!response.ok) {
			throw Error(`User ${id} not found`);
		}
		const result = await response.json();
		return result['value'];
	}

	async addHabit(habit: Habit, accessToken: string) {
		const url = `${this.url}/habits`;
		const response = await fetch(url, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${accessToken}` },
			body: JSON.stringify(habit)
		});
		if (!response.ok) {
			throw Error(`Could not create habit ${habit.name}`);
		}
		const result = await response.json();
		return result['value'];
	}

	async recordHabit(record: HabitRecord, accessToken: string) {
		const url = `${this.url}/habits/${record.habitID}/record`;
		const data = { date: record.date };
		const response = await fetch(url, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${accessToken}` },
			body: JSON.stringify(data)
		});
		if (!response.ok) {
			throw Error('Could not record habit');
		}
		const result = await response.json();
		return result['value'];
	}
}
