import { SocialHabitsClient } from '$lib/server/socialhabits';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	const client = new SocialHabitsClient();
	const user = await client.getUser(1);
	return {
		user: user
	};
};
