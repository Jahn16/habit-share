<script lang="ts">
	import { signIn } from '@auth/sveltekit/client';
	import { page } from '$app/stores';
	import type { Habit, HabitRecord } from '../models.ts';
	import Record from '../components/record.svelte';
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

<div>
	<button on:click={signIn} />
</div>
<nav>
	<img
		src={$page.data.session?.user?.image ?? 'https://source.boringavatars.com/marble/120/'}
		alt="User Avatar"
	/>
</nav>
<ul>
	{#each data.habits as habit}
		<span>{habit.name}:</span>
		{#each dates as date}
			<Record done={habitRecorded(habit, date)} />
		{/each}
		<br />
	{/each}
</ul>
