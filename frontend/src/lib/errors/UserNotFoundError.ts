export class UserNotFoundError extends Error {
	constructor(public userID?: string) {
		super(userID ? `User ${userID} not found` : 'User not found');
		this.name = 'UserNotFoundError';
	}
}
