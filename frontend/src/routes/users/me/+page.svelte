<script lang="ts">
	import { signIn } from '@auth/sveltekit/client';
	import { Button, Col, Container, Form, Input, Row, Table } from '@sveltestrap/sveltestrap';
	import type { Habit, HabitRecord, User } from '../../../models';
	import Record from '../../../components/record.svelte';
	export let data: { user: User };
	let dayNumbers = Array.from({ length: 30 }, (_, i) => i + 1);
	let dates = dayNumbers.map((day) => `2024-06-${day > 9 ? day : '0' + day}T00:00:00Z`);
	const getRecord = (habit: Habit, day: string): HabitRecord | undefined => {
		return habit.records.find((r) => r.date.startsWith(day));
	};
</script>

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
							<Record record={getRecord(habit, date)} habitId={habit.ID ? habit.ID : -1} {date} />
						</td>
					{/each}
				</tr>
			{/each}
		</tbody>
	</Table>
	<Container>
		<Form method="post" action="?/add">
			<Row>
				<Col>
					<Input type="text" placeholder="Habit Name" name="name" />
				</Col>
				<Col>
					<Input type="select">
						{#each [1, 2, 3, 4, 5, 6, 7] as option}
							<option>{option}</option>
						{/each}
					</Input>
				</Col>
				<Col>
					<Button color="primary" children="Add" />
				</Col>
			</Row>
		</Form>
	</Container>
</Container>
