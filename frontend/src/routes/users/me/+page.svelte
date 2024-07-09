<script lang="ts">
	import {
		Button,
		Col,
		Container,
		Form,
		FormGroup,
		Icon,
		Input,
		Table
	} from '@sveltestrap/sveltestrap';

	import type { Habit, HabitRecord, User, Quote } from '../../../models';
	import Record from '../../../components/record.svelte';

	import { Sound } from 'svelte-sound';
	import notificationSoundSrc from '$lib/assets/ding.mp3';

	const notificationSound = new Sound(notificationSoundSrc);

	export let data: { user?: User; quote: Quote };
	let addingHabit = false;
	let dayNumbers = Array.from({ length: 30 }, (_, i) => i + 1);
	let dates = dayNumbers.map((day) => `2024-06-${day > 9 ? day : '0' + day}T00:00:00Z`);
	const getRecord = (habit: Habit, day: string): HabitRecord | undefined => {
		return habit.records.find((r) => r.date.startsWith(day));
	};
</script>

{#if data.user}
	<Container>
		<Container>
			<h3>Welcome back, {data.user.name}</h3>
			<br />
			<blockquote cite="https://www.huxley.net/bnw/four.html">
				<p>
					{data.quote.quote}
				</p>
				<footer>— <cite>{data.quote.source}</cite></footer>
			</blockquote>
		</Container>

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
								<Record
									record={getRecord(habit, date)}
									habitId={habit.ID ? habit.ID : -1}
									{date}
									{notificationSound}
								/>
							</td>
						{/each}
					</tr>
				{/each}
			</tbody>
		</Table>
		<Container>
			{#if addingHabit}
				<Form method="post" action="?/add" class="row align-items-center">
					<Col xs="auto"
						><Button
							type="button"
							size="md"
							on:click={() => {
								addingHabit = false;
							}}><Icon name="x-square" /></Button
						></Col
					>
					<Col>
						<Input name="name" placeholder="Habit Name" required />
					</Col>
					<Col>
						<Input type="select">
							<option disabled selected>Weekly Goal</option>
							{#each [1, 2, 3, 4, 5, 6, 7] as option}
								<option>{option}</option>
							{/each}
						</Input>
					</Col>
					<Col>
						<Button color="primary" size="md"><Icon name="arrow-right-square" /></Button>
					</Col>
				</Form>
			{:else}
				<Button color="primary" size="md" on:click={() => (addingHabit = true)}
					><Icon name="plus-square" /></Button
				>
			{/if}
		</Container>
	</Container>
{:else}
	<Container>
		<h4>
			Hi, it seems like it's your first time here.<br /> Let's start by creating your first habit
		</h4>
		<br />
		<Form method="post" action="?/setup">
			<FormGroup>
				<FormGroup floating label="Habit Name">
					<Input name="habit-name" required />
				</FormGroup>
			</FormGroup>
			<FormGroup floating label="Weekly Goal">
				<Input type="select" name="habit-goal">
					{#each [1, 2, 3, 4, 5, 6, 7] as option}
						<option>{option}</option>
					{/each}
				</Input>
			</FormGroup>
			<Button children="Submit" color="primary" />
		</Form>
	</Container>
{/if}
