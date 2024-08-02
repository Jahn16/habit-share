<script lang="ts">
	import { page } from '$app/stores';
	import { Container, Icon } from '@sveltestrap/sveltestrap';
	import type { User } from '../../../../models';
	import FriendCard from '../../../../components/friendCard.svelte';

	export let data: { user: User };
	const isCurrentUser = $page.params.id === 'me';
	let form: HTMLFormElement;
	let friendInput: HTMLInputElement;
	let onRemove: ((friend: User) => void) | undefined;
	if (isCurrentUser) {
		onRemove = (f) => {
			friendInput.value = f.id || '';
			form.submit();
		};
	}
</script>

<form action="?/remove" method="POST" bind:this={form}>
	<input type="hidden" name="friend-id" bind:this={friendInput} />
</form>

<Container style="max-width: 40%;">
	<h3 class="text-center mb-4">Friends <Icon name="people" /></h3>
	{#each data.user.friends as friend}
		<FriendCard {friend} {onRemove} />
	{/each}
</Container>

<style>
	.friend-url {
		color: inherit;
		text-decoration: none;
	}
</style>
