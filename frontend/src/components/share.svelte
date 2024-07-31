<script lang="ts">
	import { Icon } from '@sveltestrap/sveltestrap';
	import { copy } from 'svelte-copy';
	import { page } from '$app/stores';

	export let userID = '';
	const userPageURL = `${$page.url.origin}/users/${userID}`;
</script>

<button class="share-button btn btn-light" use:copy={userPageURL}><Icon name="share-fill" /></button
>

<style>
	.share-button:hover {
		cursor: pointer;
	}
</style>
