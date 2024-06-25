import { error } from '@sveltejs/kit';
import { SocialHabitsClient } from '$lib/server/socialhabits';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	const client = new SocialHabitsClient();
	let user;
	try {
		user = await client.getUser(parseInt(params.id));
	} catch (e: unknown) {
		if (e instanceof Error) {
			error(404, { message: e.message });
		}
		error(500);
	}
	return {
		user: user
	};
};
