<script lang="ts">
	import { onMount } from 'svelte';

	import { Modal } from '@sveltestrap/sveltestrap';

	let open = false;
	const toggle = () => (open = !open);
	export function show(): void {
		open = !open;
	}
	export let onSelect: (emoji: string) => void = () => {};
	onMount(() => import('emoji-picker-element'));
</script>

<Modal body isOpen={open} {toggle} id="emoji-picker-modal">
	<emoji-picker
		on:emoji-click={(e) => {
			toggle();
			onSelect(e.detail.unicode);
		}}
	/>
</Modal>

<style>
	:global(#emoji-picker-modal .modal-dialog) {
		width: fit-content !important;
	}
</style>
