export default function authenticationReducer(state = {name: 123, password: 'test'}, action) {
	switch (action.type) {
		case 'SIGNUP':
			console.log('signup vi reducer');
			Object.assign({}, state,
				{
					username: action.username,
					password: action.password
				}
			);
			break;
		case 'SIGNIN':
			break;
		default:
			return {username: '', password: ''};
	}
};
