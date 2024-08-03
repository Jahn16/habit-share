<script lang="ts">
	import { page } from '$app/stores';
	import { Container, Icon } from '@sveltestrap/sveltestrap';
	import type { User } from '../../../../models';
	import FriendCard from '../../../../components/friendCard.svelte';
	import Share from '../../../../components/share.svelte';

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
	const userPageURL = `${$page.url.origin}/users/${data.user.slug}`;
</script>

<form action="?/remove" method="POST" bind:this={form}>
	<input type="hidden" name="friend-id" bind:this={friendInput} />
</form>

<Container style="max-width: 40%;">
	<h3 class="text-center mb-4">Friends <Icon name="people" /></h3>
	{#each data.user.friends as friend}
		<FriendCard {friend} {onRemove} />
	{/each}
	{#if data.user.friends.length == 0}
		<div class="text-center">
			<p>There is nothing here 😔</p>
			{#if isCurrentUser}
				<p>Share your page link to get more friends!</p>
				<Share userID={data.user.slug} />
			{/if}
		</div>
	{/if}
</Container>
