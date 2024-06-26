import { error, type Actions } from '@sveltejs/kit';
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
	}
};
