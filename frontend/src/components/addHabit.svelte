<script lang="ts">
	import { Button, Col, Container, Form, Icon, Input, Table } from '@sveltestrap/sveltestrap';
	import EmojiPicker from './emojiPicker.svelte';

	let addingHabit = false;
	let emojiPicker: EmojiPicker;
	let icon = '';
</script>

<EmojiPicker
	bind:this={emojiPicker}
	onSelect={(emoji) => {
		icon = emoji;
	}}
/>
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

			<Col xs="1"
				><Input
					name="icon"
					placeholder="☺"
					id="icon"
					on:click={() => {
						emojiPicker.show();
					}}
					value={icon}
					required
				/></Col
			>
			<Col>
				<Input name="name" placeholder="Habit Name" required />
			</Col>
			<Col>
				<Input type="select" name="goal">
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

<style>
	:global(#icon) {
		caret-color: transparent;
		text-align: center;
	}
</style>
