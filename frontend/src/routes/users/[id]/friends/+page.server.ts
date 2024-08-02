import { error, type Actions } from '@sveltejs/kit';
import { SocialHabitsClient } from '$lib/server/socialhabits';
import type { PageServerLoad } from './$types';
import type { User } from '@auth/sveltekit';
import { UserNotFoundError } from '$lib/errors/UserNotFoundError';

export const load: PageServerLoad = async ({ params, locals }) => {
	const session = await locals.auth();
	if (params.id === 'me' && !session) {
		error(403, { message: 'You must be authenticated to access this page' });
	}
	const client = new SocialHabitsClient();
	let user: User;
	try {
		if (params.id === 'me' && session) {
			user = await client.getMe(session.accessToken);
		} else {
			user = await client.getUser(params.id);
		}
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
