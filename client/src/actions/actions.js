export default actions = {
	signup(username,pass){
		return {
			type: 'SIGNUP',
			username: username,
			password: pass
		}
	}
}