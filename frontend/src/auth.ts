import { SvelteKitAuth } from '@auth/sveltekit';
import Auth0 from '@auth/sveltekit/providers/auth0';
import { env } from '$env/dynamic/private';

declare module '@auth/sveltekit' {
	interface Session {
		accessToken: string;
	}
}

export const { handle, signIn } = SvelteKitAuth({
	providers: [
		Auth0({ authorization: { params: { scope: 'openid profile email offline_access' } } })
	],
	callbacks: {
		async jwt({ token, account }) {
			if (account) {
				return {
					...token,
					accessToken: account.id_token,
					refreshToken: account.refresh_token,
					expiresAt: account.expires_at ? account.expires_at * 1000 : 0
				};
			} else if (typeof token.expiresAt === 'number' && token.expiresAt > Date.now()) {
				return token;
			} else {
				if (!token.refreshToken) {
					throw Error('Missing refresh token');
				}
				const url = new URL('/oauth/token', env.AUTH_AUTH0_ISSUER);
				const response = await fetch(url.toString(), {
					method: 'POST',
					headers: {
						'Content-Type': 'application/x-www-form-urlencoded'
					},
					body: new URLSearchParams({
						grant_type: 'refresh_token',
						client_id: env.AUTH_AUTH0_ID,
						client_secret: env.AUTH_AUTH0_SECRET,
						refresh_token: token.refreshToken as string
					})
				});
				if (!response.ok) {
					throw Error('Could not refresh token');
				}
				const result = await response.json();
				return {
					...token,
					accessToken: result['id_token'],
					expiresAt: Date.now() + result['expires_in'] * 1000
				};
			}
		},
		async session({ session, token }) {
			if (typeof token.accessToken === 'string') {
				session.accessToken = token.accessToken;
			}
			return session;
		}
	},
	trustHost: true,
	secret: env.AUTH_AUTH0_SECRET
});
