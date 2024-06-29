import { UserNotFoundError } from '$lib/errors/UserNotFoundError';
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
			throw new UserNotFoundError(id);
		}
		const result = await response.json();
		return result['value'];
	}

	async getMe(accessToken: string): Promise<User> {
		const url = `${this.url}/users/me`;
		const response = await fetch(url, { headers: { Authorization: `Bearer ${accessToken}` } });
		if (!response.ok) {
			console.log(response.status);
			if (response.status == 404) {
				throw new UserNotFoundError();
			}
			throw Error('Could not get User');
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

	async deleteRecord(ID: number, accessToken: string) {
		const url = `${this.url}/records/${ID}`;
		const response = await fetch(url, {
			method: 'DELETE',
			headers: { Authorization: `Bearer ${accessToken}` }
		});
		if (!response.ok) {
			throw Error(`Could not delete record ${ID}`);
		}
		const result = await response.json();
		return result['value'];
	}
}
