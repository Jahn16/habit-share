import type { Habit } from '../models';

export function load(): { habits: Habit[] } {
	return {
		habits: [
			{
				name: 'Beber aqua',
				goal: 7
			},
			{
				name: 'Malhar',
				goal: 4
			}
		]
	};
}
