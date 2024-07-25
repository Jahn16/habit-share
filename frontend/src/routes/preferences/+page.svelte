<script lang="ts">
	import {
		Button,
		Col,
		Container,
		Form,
		FormGroup,
		Icon,
		Input,
		Label,
		Row
	} from '@sveltestrap/sveltestrap';
	import palette from 'google-palette';
	import rawPalettes from '$lib/data/palettes.json';
	import type { User } from '../../models';
	import { enhance } from '$app/forms';

	const palettes: { [key: string]: string[] } = rawPalettes;
	const paletteKinds = Object.keys(palettes);
	let colors: string[];

	export let data: { user: User };

	const getPaletteKind = (paletteName: string): string => {
		for (const pk of paletteKinds) {
			if (palettes[pk].includes(paletteName)) {
				return pk;
			}
		}
		return 'N/A';
	};
	let paletteName = data.user.colorPalette || 'cb-Blues';
	let paletteKind = getPaletteKind(paletteName);

	const updateColors = () => {
		colors = palette(paletteName, 9, 0) || [];
		colors.reverse();
	};
	updateColors();
</script>

<Container class="w-50">
	<form
		action="?/update"
		method="POST"
		use:enhance={() => {
			return async ({ update }) => {
				update({ reset: false });
			};
		}}
	>
		<legend>User</legend>
		<FormGroup>
			<Label>Name</Label>
			<Input name="username" value={data.user.name} />
		</FormGroup>
		<legend>Palette</legend>
		<FormGroup>
			<Label>Type</Label>
			<Input type="select" bind:value={paletteKind}>
				{#each paletteKinds as pk}
					<option>{pk}</option>
				{/each}
			</Input>
		</FormGroup>
		<FormGroup>
			<Label for="color-pallete">Palette</Label>
			<Input type="select" name="color-palette" bind:value={paletteName} on:change={updateColors}>
				{#each palettes[paletteKind] as p}
					<option>{p}</option>
				{/each}
			</Input>
		</FormGroup>
		<Row style="border: thin solid;">
			{#each colors as color}
				<Col class="color-square" style="width: 1em;height:1em;background-color: #{color};"></Col>
			{/each}
		</Row>
		<Button type="submit" color="primary" class="mt-3 float-end"
			><Icon name="check-square" /></Button
		>
	</form>
</Container>

<style>
	.color-square {
		height: 1em;
		width: 1em;
	}
</style>
