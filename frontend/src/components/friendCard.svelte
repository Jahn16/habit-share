<script lang="ts">
	import { Icon } from '@sveltestrap/sveltestrap';
	import type { User } from '../models';

	export let friend: User;
	export let onRemove: ((user: User) => void) | undefined;
</script>

<div class="card mb-3">
	<div class="card-header">
		<ul class="nav card-header-tabs justify-content-end">
			<li class="nav-item">
				<a class="nav-link" href="/users/{friend.slug}"
					>Profile <Icon name="box-arrow-up-right" /></a
				>
			</li>
			{#if onRemove}
				<li class="nav-item">
					<button
						class="nav-link text-danger"
						on:click={() => {
							onRemove(friend);
						}}>Remove <Icon name="person-dash" /></button
					>
				</li>
			{/if}
		</ul>
	</div>
	<div class="row g-0">
		<div class="col-md-auto">
			<img src={friend.picture} class="img-fluid rounded-start" alt="..." />
		</div>
		<div class="col-md-auto">
			<div class="card-body">
				<h5 class="card-title">
					<a href="/users/{friend.slug}" class="friend-url">
						{friend.name}
					</a>
				</h5>
				<p class="card-text">
					<small class="text-body-secondary"><Icon name="people-fill" /> </small>
				</p>
			</div>
		</div>
	</div>
</div>
