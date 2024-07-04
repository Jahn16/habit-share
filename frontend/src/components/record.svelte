<script lang="ts">
	import type { HabitRecord } from '../models';
	export let habitId: number;
	export let date: string;
	export let record: HabitRecord | undefined = undefined;

	let form: HTMLFormElement;
	const submitForm = () => {
		if (form) {
			form.submit();
		}
	};
</script>

{#if !record}
	<form method="post" action="?/record" bind:this={form}>
		<input type="hidden" value={habitId} name="habit-id" />
		<input type="hidden" value={date} name="date" />
		<input type="checkbox" on:change={submitForm} />
	</form>
{:else}
	<form method="post" action="?/undo" bind:this={form}>
		<input type="hidden" value={record.ID} name="record-id" />
		<input type="checkbox" on:change={submitForm} checked />
	</form>
{/if}
