<script lang="ts">
	import { Col, Container, Icon, Row, Table } from '@sveltestrap/sveltestrap';
	import { Howl } from 'howler';

	import type { Habit, HabitRecord, User, Quote } from '../../../models';
	import Record from '../../../components/record.svelte';
	import AddHabit from '../../../components/addHabit.svelte';
	import Header from '../../../components/headers.svelte';

	import notificationSoundSrc from '$lib/assets/ding.mp3';
	import EditHabitModal from '../../../components/editHabitModal.svelte';
	import FirstHabit from '../../../components/firstHabit.svelte';
	import Share from '../../../components/share.svelte';
	import { getColors } from '$lib/utils/colors';

	const notificationSound = new Howl({ src: [notificationSoundSrc] });

	export let data: { user?: User };

	const daysInMonth = (): number => {
		var now = new Date();
		return new Date(now.getFullYear(), now.getMonth() + 1, 0).getDate();
	};
	let dayNumbers = Array.from({ length: daysInMonth() }, (_, i) => i + 1);
	let dates = dayNumbers.map((day) => `2024-06-${day > 9 ? day : '0' + day}T00:00:00Z`);

	const getRecord = (habit: Habit, day: string): HabitRecord | undefined => {
		return habit.records.find((r) => r.date.startsWith(day));
	};
	let editHabitModal: EditHabitModal;

	let colors: string[] = [];
	if (data.user) {
		colors = getColors(data.user.colorPalette, data.user.habits.length);
	}
</script>

{#if data.user}
	<Container>
		<Header username={data.user.name} />
		<Table hover={true}>
			<thead>
				<tr>
					<th scope="col"></th>
					<th scope="col"></th>
					<th scope="col">#</th>
					{#each dates as _, i}
						<th scope="col">{i + 1}</th>
					{/each}
				</tr>
			</thead>
			<tbody>
				{#each data.user.habits as habit, i}
					<tr>
						<th scope="row"
							><button
								class="edit-habit text-secondary"
								on:click={() => {
									editHabitModal.editHabit(habit);
								}}><Icon name="pencil-square" /></button
							></th
						>
						<th scope="row">{habit.icon}</th>
						<th scope="row">{habit.name}</th>
						{#each dates as date}
							<td>
								<Record
									record={getRecord(habit, date)}
									habitId={habit.ID ? habit.ID : -1}
									{date}
									{notificationSound}
									color={'#' + colors[i]}
								/>
							</td>
						{/each}
					</tr>
				{/each}
			</tbody>
		</Table>
		<Row class="justify-content-between">
			<Col xs="auto">
				<AddHabit />
			</Col>
			<Col xs="auto">
				<Share userID={data.user.slug} />
			</Col>
		</Row>
		<EditHabitModal bind:this={editHabitModal} />
	</Container>
{:else}<FirstHabit />{/if}

<style>
	.edit-habit {
		all: unset;
	}
	.edit-habit:hover {
		cursor: pointer;
	}
</style>
