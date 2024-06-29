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
	add: async ({ request, locals }) => {
		const session = await locals.auth();
		if (!session) {
			return;
		}
		const data = await request.formData();
		const client = new SocialHabitsClient();
		const habitName = data.get('name') as string;
		const habitGoal = data.get('goal') as string;
		client.addHabit(
			{ name: habitName, goal: parseInt(habitGoal), records: [] },
			session.accessToken
		);
	},
	record: async ({ request, locals }) => {
		const session = await locals.auth();
		if (!session) {
			return;
		}

		const data = await request.formData();
		const habitId = data.get('habit-id') as string;
		const date = data.get('date') as string;
		const client = new SocialHabitsClient();
		client.recordHabit({ habitID: parseInt(habitId), date: date }, session.accessToken);
	},
	undo: async ({ request, locals }) => {
		const session = await locals.auth();
		if (!session) {
			return;
		}

		const data = await request.formData();
		const recordID = data.get('record-id') as string;
		const client = new SocialHabitsClient();
		client.deleteRecord(parseInt(recordID), session.accessToken);
	}
};
