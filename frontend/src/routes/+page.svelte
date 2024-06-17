<script lang="ts">
	import type { Habit, HabitRecord } from '../models.ts';
	export let data: { habits: Habit[] };
	let dayNumbers = Array.from({ length: 30 }, (_, i) => i + 1);
	let dates = dayNumbers.map((day) => `2024-06-${day}T00:00:00Z`);
	const habitRecorded = (habit: Habit, day: string): boolean => {
		const habitRecordedInThatDay = (record: HabitRecord): boolean => {
			return record.date === day;
		};
		return habit.records.some(habitRecordedInThatDay);
	};
</script>

<ul>
	{#each dates as date}
		{#each data.habits as habit}
			{#if habitRecorded(habit, date)}
				<li>{date}: {habit.name}</li>
			{/if}
		{/each}
	{/each}
</ul>
