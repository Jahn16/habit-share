<script lang="ts">
	import {
		Button,
		Col,
		Container,
		Icon,
		Input,
		Modal,
		ModalFooter,
		Row
	} from '@sveltestrap/sveltestrap';

	import type { Habit } from '../models';
	import EmojiPicker from './emojiPicker.svelte';
	let open = false;
	const toggle = () => (open = !open);
	export function editHabit(selectedHabit: Habit): void {
		habit = selectedHabit;
		open = true;
	}

	let habit: Habit;
	let updateForm: HTMLFormElement;
	let deleteForm: HTMLFormElement;
	let emojiPicker: EmojiPicker;
	const onEmojiSelect = (emoji: string) => {
		habit.icon = emoji;
	};
</script>

{#if habit}
	<EmojiPicker onSelect={onEmojiSelect} bind:this={emojiPicker} />
	<Modal body header="Editing habit: {habit.name}" isOpen={open} {toggle}>
		<form method="POST" action="?/update" bind:this={updateForm}>
			<input type="hidden" name="habit-id" value={habit.ID} />
			<div class="mb-3">
				<label for="edit-habit-icon" class="form-label">Habit Icon</label>
				<Input
					type="text"
					class="form-control"
					id="edit-habit-icon"
					name="habit-icon"
					value={habit.icon}
					on:click={() => emojiPicker.show()}
					readonly
				/>
			</div>

			<div class="mb-3">
				<label for="edit-habit-name" class="form-label">Habit Name</label>
				<Input
					type="text"
					class="form-control"
					id="edit-habit-name"
					name="habit-name"
					value={habit.name}
				/>
			</div>
			<div class="mb-3">
				<label for="edit-habit-goal" class="form-label">Weekly Goal</label>
				<Input type="select" id="edit-habit-goal" name="habit-goal" class="form-select">
					{#each [1, 2, 3, 4, 5, 6, 7] as goal}
						{#if goal === habit.goal}
							<option selected>{goal}</option>
						{:else}
							<option>{goal}</option>
						{/if}
					{/each}
				</Input>
			</div>
		</form>
		<form method="POST" action="?/delete" bind:this={deleteForm}>
			<input type="hidden" name="habit-id" value={habit.ID} />
		</form>
		<ModalFooter>
			<Container>
				<Row class="justify-content-between">
					<Col sm="auto">
						<Button
							color="danger"
							on:click={() => {
								toggle();
								deleteForm.requestSubmit();
							}}><Icon name="trash" /></Button
						>
					</Col>
					<Col sm="auto">
						<Button
							color="primary"
							on:click={() => {
								toggle();
								updateForm.requestSubmit();
							}}><Icon name="check-square" /></Button
						>
					</Col>
				</Row>
			</Container>
		</ModalFooter>
	</Modal>
{/if}
