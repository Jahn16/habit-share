import { UserNotFoundError } from '$lib/errors/UserNotFoundError';
import { SocialHabitsClient } from '$lib/server/socialhabits';
import type { User } from '../../../models';
import type { PageServerLoad } from './$types';
import { error } from '@sveltejs/kit';

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
