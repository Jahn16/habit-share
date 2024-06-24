import type { User } from '../../models';
export class SocialHabitsClient {
	private url: string;
	constructor() {
		this.url = 'http://localhost:8000';
	}

	async getUser(id: number): Promise<User> {
		const url = `${this.url}/users/${id}`;
		const response = await fetch(url);
		const result = await response.json();
		return result['value'];
	}
}
