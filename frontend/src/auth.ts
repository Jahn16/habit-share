import { SvelteKitAuth } from '@auth/sveltekit';
import Google from '@auth/sveltekit/providers/google';
// import { DefaultSession } from '@auth/core/types';

declare module '@auth/sveltekit' {
	interface Session {
		accessToken: string;
	}
}

export const { handle } = SvelteKitAuth({
	providers: [Google],
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
