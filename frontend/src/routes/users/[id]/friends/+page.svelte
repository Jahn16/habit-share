<script lang="ts">
	import { Container, Icon } from '@sveltestrap/sveltestrap';
	import type { User } from '../../../../models';

	export let data: { user: User };
</script>

<Container class="w-50">
	<h3 class="text-center mb-4">Friends <Icon name="people" /></h3>
	{#each data.user.friends as friend}
		<div class="card mb-3">
			<div class="row g-0">
				<div class="col-md-auto">
					<img src={friend.picture} class="img-fluid rounded-start" alt="..." />
				</div>
				<div class="col-md-auto">
					<div class="card-body">
						<h5 class="card-title">
							<a href="/users/{friend.slug}" class="friend-url">
								{friend.name}
								<Icon name="box-arrow-up-right" class="text-primary" />
							</a>
						</h5>
						<p class="card-text">
							<small class="text-body-secondary"><Icon name="people-fill" /> 3</small>
						</p>
					</div>
				</div>
			</div>
		</div>
	{/each}
</Container>

<style>
	.friend-url {
		color: inherit;
		text-decoration: none;
	}
</style>
