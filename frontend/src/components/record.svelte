<script lang="ts">
	import { Form } from '@sveltestrap/sveltestrap';

	import type { HabitRecord } from '../models';
	export let habitId: number;
	export let date: string;
	export let record: HabitRecord | undefined = undefined;
</script>

{#if !record}
	<Form method="post" action="?/record">
		<input type="hidden" value={habitId} name="habit-id" />
		<input type="hidden" value={date} name="date" />
		<input type="checkbox" onChange="this.form.submit()" />
	</Form>
{:else}
	<Form method="post" action="?/undo">
		<input type="hidden" value={record.ID} name="record-id" />
		<input type="checkbox" onChange="this.form.submit()" checked />
	</Form>
{/if}
