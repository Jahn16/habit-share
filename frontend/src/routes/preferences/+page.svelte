<script lang="ts">
	import {
		Button,
		Col,
		Container,
		Form,
		FormGroup,
		Input,
		Label,
		Row
	} from '@sveltestrap/sveltestrap';
	import palette from 'google-palette';
	import palettes from '$lib/data/palettes.json';

	const paletteKinds = Object.keys(palettes);
	let paletteKind = 'sequential';
	let colors: string[];
	let paletteName = 'cb-Blues';

	const updateColors = () => {
		colors = palette(paletteName, 9, 0) || [];
		colors.reverse();
	};
	updateColors();
</script>

<Container class="w-50">
	<Form>
		<legend>User</legend>
		<FormGroup>
			<Label>Name</Label>
			<Input name="username" />
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
			<Input type="select" id="color-pallete" bind:value={paletteName} on:change={updateColors}>
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
		<Button type="submit" color="primary" class="mt-3">Submit</Button>
	</Form>
</Container>

<style>
	.color-square {
		height: 1em;
		width: 1em;
	}
</style>
