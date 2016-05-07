export default function authenticationReducer(state, action) {
	action |= {};
	switch (action.type) {
		case 'SIGNUP':
			Object.assign({}, state, {
				user: {
					username: action.username,
					password: action.password
				}
			});
			break;
		case 'SIGNIN':
			break;
		default:
			return state;
	}
};
