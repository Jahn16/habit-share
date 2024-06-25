<script lang="ts">
	import { signIn } from '@auth/sveltekit/client';
	import { page } from '$app/stores';
	import { Container, Table } from '@sveltestrap/sveltestrap';
	import type { Habit, HabitRecord, User } from '../../../models';
	import Record from '../../../components/record.svelte';
	export let data: { user: User };
	let dayNumbers = Array.from({ length: 30 }, (_, i) => i + 1);
	let dates = dayNumbers.map((day) => (day > 9 ? `2024-06-${day}` : `2024-06-0${day}`));
	const habitRecorded = (habit: Habit, day: string): boolean => {
		const habitRecordedInThatDay = (record: HabitRecord): boolean => {
			return record.date.startsWith(day);
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

<Container>
	<Table>
		<thead>
			<tr>
				<th scope="col">#</th>
				{#each dates as _, i}
					<th scope="col">{i + 1}</th>
				{/each}
			</tr>
		</thead>
		<tbody>
			{#each data.user.habits as habit}
				<tr>
					<th scope="row">{habit.name}</th>
					{#each dates as date}
						<td>
							<Record done={habitRecorded(habit, date)} />
						</td>
					{/each}
				</tr>
			{/each}
		</tbody>
	</Table>
</Container>
