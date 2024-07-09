<script lang="ts">
	import { Button, Col, Container, Form, Icon, Input, Table } from '@sveltestrap/sveltestrap';
	let addingHabit = false;
</script>

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
