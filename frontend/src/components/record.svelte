<script lang="ts">
	import { Form } from '@sveltestrap/sveltestrap';

	// import type { HabitRecord } from '../models';
	export let done = false;
	export let habitId: number;
	export let date: string;
</script>

<Form method="post" action="?/record">
	<input type="hidden" value={habitId} name="habit-id" />
	<input type="hidden" value={date} name="date" />
	<input type="checkbox" checked={done} onChange="this.form.submit()" />
</Form>
