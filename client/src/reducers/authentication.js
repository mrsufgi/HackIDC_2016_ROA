export default function authenticationReducer(state, action) {
	switch (action.type) {
		case 'SIGNUP':
			console.log('signup vi reducer');
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
