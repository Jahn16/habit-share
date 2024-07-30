<script lang="ts">
	import { Container } from '@sveltestrap/sveltestrap';
	import quotes from '$lib/data/quotes.json';

	const quote = quotes[Math.floor(Math.random() * quotes.length)];

	export let username: string;
</script>

<Container>
	<h3>Welcome back, {username}</h3>
	<br />
	<blockquote cite="https://www.huxley.net/bnw/four.html">
		<p>{quote.quote}</p>
		<footer>— <cite>{quote.source}</cite></footer>
	</blockquote>
</Container>
