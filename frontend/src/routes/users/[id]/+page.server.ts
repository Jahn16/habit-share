import { UserNotFoundError } from '$lib/errors/UserNotFoundError';
import { SocialHabitsClient } from '$lib/server/socialhabits';
import type { User } from '../../../models';
import type { PageServerLoad } from './$types';
import { error, redirect, type Actions } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ params, locals }) => {
	const client = new SocialHabitsClient();
	let user: User;
	try {
		user = await client.getUser(params.id);
	} catch (e: unknown) {
		if (e instanceof UserNotFoundError) {
			error(404, { message: e.message });
		}
		if (e instanceof Error) {
			error(500, { message: e.message });
		}
		error(500);
	}
	const session = await locals.auth();
	if (!session) {
		return {
			user: user,
			loggedUser: null
		};
	}
	let loggedUser: User | null;
	try {
		loggedUser = await client.getMe(session.accessToken);
	} catch {
		loggedUser = null;
	}

	if (user.id === loggedUser?.id) {
		redirect(307, '/users/me');
	}
	return {
		user: user,
		loggedUser: loggedUser
	};
};

export const actions: Actions = {
	add: async ({ request, locals }) => {
		const session = await locals.auth();
		if (!session) {
			return;
		}
		const data = await request.formData();
		const client = new SocialHabitsClient();

		const friendID = data.get('friend-id') as string;
		await client.addFriend(friendID, session.accessToken);
	},
	remove: async ({ request, locals }) => {
		const session = await locals.auth();
		if (!session) {
			return;
		}
		const data = await request.formData();
		const client = new SocialHabitsClient();

		const friendID = data.get('friend-id') as string;
		console.log(`FriendID: ${friendID}`);
		await client.removeFriend(friendID, session.accessToken);
	}
};
