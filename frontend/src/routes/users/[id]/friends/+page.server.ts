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
			error(404, { message: e.message });
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
export const actions: Actions = {
	remove: async ({ request, locals }) => {
		const session = await locals.auth();
		if (!session) {
			return;
		}
		const data = await request.formData();
		const client = new SocialHabitsClient();

		const friendID = data.get('friend-id') as string;
		console.log(`FriendID: ${friendID}`);
		try {
			await client.removeFriend(friendID, session.accessToken);
		} catch {}
	}
};
