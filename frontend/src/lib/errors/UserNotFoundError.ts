export class UserNotFoundError extends Error {
	constructor(public userID?: number) {
		super(userID ? `User ${userID} not found` : 'User not found');
		this.name = 'UserNotFoundError';
	}
}
