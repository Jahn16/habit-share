<script lang="ts">
	import type { HabitRecord } from '../models';
	import { enhance } from '$app/forms';
	import { Howl } from 'howler';

	export let color: string = 'red';
	export let habitId: number;
	export let date: string;
	export let record: HabitRecord | undefined = undefined;
	export let notificationSound: Howl;

	let form: HTMLFormElement;
	const submitForm = () => {
		if (form) {
			form.requestSubmit();
		}
	};
</script>

<div style="--color: {color};">
	{#if !record}
		<form
			method="post"
			action="?/record"
			bind:this={form}
			use:enhance={() => {
				notificationSound.play();
				return async ({ update }) => {
					update({ reset: false });
				};
			}}
		>
			<input type="hidden" value={habitId} name="habit-id" />
			<input type="hidden" value={date} name="date" />
			<input type="checkbox" on:change={submitForm} />
		</form>
	{:else}
		<form
			method="post"
			action="?/undo"
			bind:this={form}
			use:enhance={() => {
				return async ({ update }) => {
					update({ reset: false });
				};
			}}
		>
			<input type="hidden" value={record.ID} name="record-id" />
			<input type="checkbox" on:change={submitForm} checked />
		</form>
	{/if}
</div>

<style>
	input[type='checkbox'] {
		accent-color: var(--color);
	}
</style>
