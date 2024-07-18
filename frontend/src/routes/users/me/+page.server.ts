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
	setup: async ({ request, locals }) => {
		const session = await locals.auth();
		if (!session) {
			error(403);
		}
		const data = await request.formData();
		const habitName = data.get('habit-name') as string;
		const habitGoal = data.get('habit-goal') as string;

		const client = new SocialHabitsClient();
		await client
			.getMe(session.accessToken)
			.catch((e) => {
				if (e instanceof UserNotFoundError) {
					return client.createUser(session.accessToken);
				}
			})
			.then(() => {
				return client.addHabit(
					{ name: habitName, goal: parseInt(habitGoal), records: [] },
					session.accessToken
				);
			});
	},
	add: async ({ request, locals }) => {
		const session = await locals.auth();
		if (!session) {
			return;
		}
		const data = await request.formData();
		const client = new SocialHabitsClient();
		const habitName = data.get('name') as string;
		const habitGoal = data.get('goal') as string;
		await client.addHabit(
			{ name: habitName, goal: parseInt(habitGoal), records: [] },
			session.accessToken
		);
	},
	update: async ({ request, locals }) => {
		const session = await locals.auth();
		if (!session) {
			return;
		}
		const data = await request.formData();
		const client = new SocialHabitsClient();
		const habitID = data.get('habit-id') as string;
		const habitName = data.get('habit-name') as string;
		const habitGoal = data.get('habit-goal') as string;
		await client.updateHabit(
			{ ID: parseInt(habitID), name: habitName, goal: parseInt(habitGoal), records: [] },
			session.accessToken
		);
	},

	delete: async ({ request, locals }) => {
		const session = await locals.auth();
		if (!session) {
			return;
		}

		const data = await request.formData();
		const habitID = data.get('habit-id') as string;
		const client = new SocialHabitsClient();
		await client.deleteHabit(parseInt(habitID), session.accessToken);
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
		await client.recordHabit({ habitID: parseInt(habitId), date: date }, session.accessToken);
	},
	undo: async ({ request, locals }) => {
		const session = await locals.auth();
		if (!session) {
			return;
		}

		const data = await request.formData();
		const recordID = data.get('record-id') as string;
		const client = new SocialHabitsClient();
		await client.deleteRecord(parseInt(recordID), session.accessToken);
	}
};
