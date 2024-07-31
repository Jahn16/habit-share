import { UserNotFoundError } from '$lib/errors/UserNotFoundError';
import type { User, Habit, HabitRecord } from '../../models';
import { env } from '$env/dynamic/private';
import { logger } from '$lib/logger';

export class SocialHabitsClient {
	private url: string;
	constructor() {
		if (!env.BACKEND_URL) {
			throw Error('You must set BACKEND_URL env variable');
		}

		this.url = 'http://localhost:8000';
	}

	async getUser(id: string): Promise<User> {
		const url = `${this.url}/users/${id}`;
		const response = await fetch(url);
		if (!response.ok) {
			logger.log({ level: 'info', message: 'User not found', id: id });
			throw new UserNotFoundError(id);
		}
		logger.log({ level: 'info', message: 'Retrieved user data', id: id });
		const result = await response.json();
		return result['value'];
	}

	async getMe(accessToken: string): Promise<User> {
		const url = `${this.url}/users/me`;
		const response = await fetch(url, { headers: { Authorization: `Bearer ${accessToken}` } });
		if (!response.ok) {
			if (response.status == 404) {
				logger.log({ level: 'info', message: 'Current user not found' });
				throw new UserNotFoundError();
			}
			throw Error('Could not get User');
		}
		logger.log({ level: 'info', message: 'Retrieved current user data' });
		const result = await response.json();
		return result['value'];
	}

	async createUser(accessToken: string): Promise<User> {
		const url = `${this.url}/users`;
		const response = await fetch(url, {
			method: 'POST',
			headers: { Authorization: `Bearer ${accessToken}` }
		});
		if (!response.ok) {
			logger.log({ level: 'error', message: 'Could not create user' });
			throw Error('Could create user');
		}
		logger.log({ level: 'info', message: 'Created user' });
		const result = await response.json();
		return result['value'];
	}

	async updateUser(user: { name: string; colorPalette: string }, accessToken: string) {
		const url = `${this.url}/users/me`;
		const response = await fetch(url, {
			method: 'PATCH',
			headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${accessToken}` },
			body: JSON.stringify(user)
		});
		if (!response.ok) {
			logger.error({ level: 'error', message: 'Could not update user' });
			throw Error(`Could not update user ${user.name} ${await response.text()}`);
		}
		logger.log({ level: 'info', message: 'Updated user' });
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
			logger.error({ level: 'error', message: 'Could not create habit' });
			throw Error(`Could not create habit ${habit.name}`);
		}
		logger.info({ level: 'info', message: 'Created habit' });
		const result = await response.json();
		return result['value'];
	}

	async updateHabit(habit: Habit, accessToken: string) {
		const url = `${this.url}/habits/${habit.ID}`;
		const response = await fetch(url, {
			method: 'PATCH',
			headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${accessToken}` },
			body: JSON.stringify(habit)
		});
		if (!response.ok) {
			logger.error({ level: 'error', message: 'Could not update habit', habitID: habit.ID });
			throw Error(`Could not update habit ${habit.name}`);
		}
		logger.info({ level: 'info', message: 'Updated habit', habitID: habit.ID });
		const result = await response.json();
		return result['value'];
	}

	async deleteHabit(ID: number, accessToken: string) {
		const url = `${this.url}/habits/${ID}`;
		const response = await fetch(url, {
			method: 'DELETE',
			headers: { Authorization: `Bearer ${accessToken}` }
		});
		if (!response.ok) {
			logger.error({ level: 'error', message: 'Could not delete habit', habitID: ID });
			throw Error(`Could not delete record ${ID}`);
		}
		logger.info({ level: 'info', message: 'Deleted habit', habitID: ID });
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
			logger.error({
				level: 'error',
				message: 'Could not add habit register',
				habitID: record.habitID
			});
			throw Error('Could not record habit');
		}
		logger.info({ level: 'info', message: 'Could not add habit record', habitID: record.habitID });
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
			logger.error({ level: 'error', message: 'Could not delete record', recordID: ID });
			throw Error(`Could not delete record ${ID}`);
		}

		logger.error({ level: 'info', message: 'Deleted record', recordID: ID });
		const result = await response.json();
		return result['value'];
	}
}
