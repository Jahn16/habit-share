import { error, type Actions } from '@sveltejs/kit';
import { SocialHabitsClient } from '$lib/server/socialhabits';
import type { PageServerLoad } from './$types';
import type { User } from '@auth/sveltekit';
import { UserNotFoundError } from '$lib/errors/UserNotFoundError';

export const load: PageServerLoad = async ({ locals }) => {
	const session = await locals.auth();
	if (!session) {
		error(403, { message: 'You must be authenticated to access this page' });
	}

	const client = new SocialHabitsClient();
	let user: User;
	try {
		user = await client.getMe(session.accessToken);
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

export const actions: Actions = {
	update: async ({ request, locals }) => {
		const session = await locals.auth();
		if (!session) {
			return;
		}
		const data = await request.formData();
		const client = new SocialHabitsClient();
		const username = data.get('username') as string;
		const colorPalette = data.get('color-palette') as string;
		await client.updateUser(
			{
				name: username,
				colorPalette: colorPalette
			},
			session.accessToken
		);
	}
};
