import type { Habit } from '../models';

export function load(): { habits: Habit[] } {
	return {
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
	};
}
