<script lang="ts">
	import { Button, Container, Form, FormGroup, Input } from '@sveltestrap/sveltestrap';

	import EmojiPicker from './emojiPicker.svelte';

	let emojiPicker: EmojiPicker;
	let icon = '';
	const onEmojiSelect = (emoji: string) => {
		icon = emoji;
	};
</script>

<EmojiPicker onSelect={onEmojiSelect} bind:this={emojiPicker} />
<Container>
	<h4>
		Hi, it seems like it's your first time here.<br /> Let's start by creating your first habit
	</h4>
	<br />
	<Form method="post" action="?/setup">
		<div class="mb-4">
			<FormGroup floating label="Habit Icon">
				<Input
					name="habit-icon"
					value={icon}
					on:click={emojiPicker.show()}
					required
					readonly
				/></FormGroup
			>
		</div>
		<div class="mb-4">
			<FormGroup>
				<FormGroup floating label="Habit Name">
					<Input name="habit-name" required />
				</FormGroup>
			</FormGroup>
		</div>
		<div class="mb-4">
			<FormGroup floating label="Weekly Goal">
				<Input type="select" name="habit-goal">
					{#each [1, 2, 3, 4, 5, 6, 7] as option}
						<option>{option}</option>
					{/each}
				</Input>
			</FormGroup>
		</div>
		<Button color="primary">Submit</Button>
	</Form>
</Container>
