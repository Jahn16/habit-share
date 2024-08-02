<script lang="ts">
	import { Container, Table, Image, Row, Col, Icon, Button } from '@sveltestrap/sveltestrap';
	import palette from 'google-palette';

	import type { Habit, HabitRecord, User } from '../../../models';
	import { redirect } from '@sveltejs/kit';
	import { goto } from '$app/navigation';

	export let data: { user: User };

	const daysInMonth = (): number => {
		var now = new Date();
		return new Date(now.getFullYear(), now.getMonth() + 1, 0).getDate();
	};
	let dayNumbers = Array.from({ length: daysInMonth() }, (_, i) => i + 1);
	let dates = dayNumbers.map((day) => `2024-06-${day > 9 ? day : '0' + day}T00:00:00Z`);

	const getRecord = (habit: Habit, day: string): HabitRecord | undefined => {
		return habit.records.find((r) => r.date.startsWith(day));
	};

	let colors: string[] = [];
	if (data.user) {
		const habitQty = data.user.habits.length;
		const colorPalette = habitQty <= 9 ? data.user.colorPalette || 'cb-Blues' : 'tol';
		const colorQty = habitQty < 9 ? 9 : habitQty;
		colors = palette(colorPalette, colorQty);
		colors.reverse();
	}
	let addFriendForm: HTMLFormElement;
</script>

<form action="?/add" method="POST" bind:this={addFriendForm}>
	<input type="hidden" name="friend-id" value={data.user.id} />
</form>
<Container>
	<Container class="mb-3"
		><Row class="justify-content-center mb-2">
			<Col xs="auto">
				<Image src={data.user.picture} />
			</Col>
		</Row>
		<Row class="justify-content-center">
			<Col xs="auto">
				<h3 class="mt-2">{data.user.name}</h3>
			</Col>
		</Row>
		<Row class="justify-content-center">
			<Col xs="auto" class="px-1">
				<Button
					color="light"
					on:click={() => {
						addFriendForm.submit();
					}}
					><Icon name="person-fill-add" />
				</Button>
			</Col>
			<Col xs="auto" class="px-1">
				<Button
					color="light"
					class="position-relative"
					on:click={() => {
						goto(`${data.user.slug}/friends`);
					}}
				>
					<Icon name="people-fill" />
					<span
						class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-info"
					>
						{data.user.friends.length}
						<span class="visually-hidden">unread messages</span>
					</span>
				</Button>
			</Col>
		</Row>
	</Container>
	<Table hover={true}>
		<thead>
			<tr>
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
					<th scope="row">{habit.icon}</th>
					<th scope="row">{habit.name}</th>
					{#each dates as date}
						<td>
							<input
								type="checkbox"
								style="accent-color: #{colors[i]};"
								checked={!!getRecord(habit, date)}
							/>
						</td>
					{/each}
				</tr>
			{/each}
		</tbody>
	</Table>
</Container>

<style>
	input[type='checkbox'] {
		pointer-events: none;
	}
</style>
