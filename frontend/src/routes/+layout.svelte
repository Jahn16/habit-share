<script lang="ts">
	import {
		Button,
		Nav,
		NavItem,
		Navbar,
		NavbarBrand,
		Image,
		Dropdown,
		DropdownToggle,
		DropdownMenu,
		DropdownItem,
		Icon
	} from '@sveltestrap/sveltestrap';
	import { page } from '$app/stores';
	import { signIn, signOut } from '@auth/sveltekit/client';
	import { goto } from '$app/navigation';
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
				<Dropdown>
					<DropdownToggle style="all: unset;">
						<Image
							thumbnail
							src={$page.data.session?.user?.image ??
								'https://source.boringavatars.com/marble/120/'}
							alt="User Avatar"
							class="img-circle"
							style="max-width: 48px;max-height: 48px;"
						/>
					</DropdownToggle>
					<DropdownMenu>
						<DropdownItem
							on:click={async () => {
								await goto('/users/me');
							}}>Home <Icon name="house" /></DropdownItem
						>
						<DropdownItem
							on:click={async () => {
								await goto('/users/me/friends');
							}}>Friends <Icon name="people" /></DropdownItem
						>

						<DropdownItem
							on:click={async () => {
								await goto('/preferences');
							}}>Preferences <Icon name="gear" /></DropdownItem
						>
						<DropdownItem divider />
						<DropdownItem
							class="text-danger"
							on:click={() => {
								signOut();
							}}>Signout <Icon name="box-arrow-right" /></DropdownItem
						>
					</DropdownMenu>
				</Dropdown>
			</NavItem>
		{/if}
	</Nav>
</Navbar>

<slot></slot>
