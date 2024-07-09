<script lang="ts">
	import {
		Button,
		Nav,
		NavItem,
		NavLink,
		Navbar,
		NavbarBrand,
		Image
	} from '@sveltestrap/sveltestrap';
	import { page } from '$app/stores';
	import { signIn } from '@auth/sveltekit/client';
</script>

<Navbar color="dark" theme="dark" class="mb-5">
	<NavbarBrand href="/">Social Habits</NavbarBrand>
	<Nav>
		{#if !$page.data.session}
			<NavItem>
				<Button
					children="Login"
					color="primary"
					on:click={() => {
						signIn('auth0', { callbackUrl: '/users/me' });
					}}
				/>
			</NavItem>
		{:else}
			<NavItem>
				<NavLink href="/users/me">
					<Image
						thumbnail
						src={$page.data.session?.user?.image ?? 'https://source.boringavatars.com/marble/120/'}
						alt="User Avatar"
						class="img-circle"
						style="max-width: 48px;max-height: 48px;"
					/>
				</NavLink>
			</NavItem>
		{/if}
	</Nav>
</Navbar>

<slot></slot>
