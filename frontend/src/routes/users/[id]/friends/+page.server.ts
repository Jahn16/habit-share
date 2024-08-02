import { error, type Actions } from '@sveltejs/kit';
import { SocialHabitsClient } from '$lib/server/socialhabits';
import type { PageServerLoad } from './$types';
import type { User } from '@auth/sveltekit';
import { UserNotFoundError } from '$lib/errors/UserNotFoundError';

export const load: PageServerLoad = async ({ params }) => {
	const client = new SocialHabitsClient();
	let user: User;
	try {
		user = await client.getUser(params.id);
	} catch (e: unknown) {
		if (e instanceof UserNotFoundError) {
			return { user: null };
		}
		if (e instanceof Error) {
			error(500, { message: e.message });
		}
		error(500);
	}

	return {
		user: user
	};
};
