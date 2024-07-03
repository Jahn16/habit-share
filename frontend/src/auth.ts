import { SvelteKitAuth } from '@auth/sveltekit';
import Google from '@auth/sveltekit/providers/google';
import Auth0 from '@auth/sveltekit/providers/auth0';
// import { DefaultSession } from '@auth/core/types';

declare module '@auth/sveltekit' {
	interface Session {
		accessToken: string;
	}
}

export const { handle } = SvelteKitAuth({
	providers: [Auth0],
	callbacks: {
		jwt({ token, account }) {
			if (account?.id_token) {
				return { ...token, accessToken: account.id_token };
			}
			return token;
		},
		async session({ session, token }) {
			if (typeof token.accessToken === 'string') {
				session.accessToken = token.accessToken;
			}
			return session;
		}
	}
});
