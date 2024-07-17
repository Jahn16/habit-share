<script lang="ts">
	import { Button, Col, Container, Icon, Modal, ModalFooter, Row } from '@sveltestrap/sveltestrap';

	import type { Habit } from '../models';
	let open = false;
	const toggle = () => (open = !open);
	export function editHabit(selectedHabit: Habit): void {
		habit = selectedHabit;
		open = true;
	}

	let habit: Habit;
	let deleteForm: HTMLFormElement;
</script>

{#if habit}
	<Modal body header="Editing habit: {habit.name}" isOpen={open} {toggle}>
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
						<Button color="secondary" on:click={toggle}>Cancel</Button>
					</Col>
				</Row>
			</Container>
		</ModalFooter>
	</Modal>
{/if}
