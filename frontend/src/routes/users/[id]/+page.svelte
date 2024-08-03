<script lang="ts">
	import { Container, Table, Image, Row, Col, Icon, Button } from '@sveltestrap/sveltestrap';

	import type { Habit, HabitRecord, User } from '../../../models';
	import { getColors } from '$lib/utils/colors';
	import { goto } from '$app/navigation';

	export let data: { user: User; loggedUser: User | null };

	const daysInMonth = (): number => {
		var now = new Date();
		return new Date(now.getFullYear(), now.getMonth() + 1, 0).getDate();
	};
	let dayNumbers = Array.from({ length: daysInMonth() }, (_, i) => i + 1);
	let dates = dayNumbers.map((day) => `2024-06-${day > 9 ? day : '0' + day}T00:00:00Z`);

	const getRecord = (habit: Habit, day: string): HabitRecord | undefined => {
		return habit.records.find((r) => r.date.startsWith(day));
	};

	let colors: string[] = getColors(data.user.colorPalette, data.user.habits.length);
	const checkIsFriend = (): boolean => {
		if (!data.loggedUser) {
			return false;
		}
		const friendIDs = data.loggedUser.friends.map(({ id }) => id);
		return friendIDs.includes(data.user.id);
	};
	const isFriend = checkIsFriend();
	let friendForm: HTMLFormElement;
	const formAction = isFriend ? '?/remove' : '?/add';
</script>

<form action={formAction} method="POST" bind:this={friendForm}>
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
				{#if !isFriend}
					<Button
						color="light"
						on:click={() => {
							friendForm.submit();
						}}
					>
						<Icon name="person-fill-add" />
					</Button>
				{:else}
					<Button
						color="light"
						on:click={() => {
							friendForm.submit();
						}}
					>
						<Icon name="person-fill-slash" />
					</Button>
				{/if}
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
