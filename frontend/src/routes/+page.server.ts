import type { User } from '../models';

export function load(): { user: User } {
	return {
		user: {
			name: 'Jahn',
			picture:
				'https://lh3.googleusercontent.com/a/ACg8ocLtJIkiH-j_IKsapo82ne7mJUax1sifro8gTCrz38jW4OKCVV93=s96-c',
			habits: [
				{
					name: 'Beber aqua',
					goal: 7,
					records: [
						{ date: '2024-06-17T00:00:00Z' },
						{ date: '2024-06-15T00:00:00Z' },
						{ date: '2024-06-14T00:00:00Z' }
					]
				},
				{
					name: 'Malhar',
					goal: 4,
					records: [
						{ date: '2024-06-15T00:00:00Z' },
						{ date: '2024-06-14T00:00:00Z' },
						{ date: '2024-06-13T00:00:00Z' }
					]
				}
			]
		}
	};
}
