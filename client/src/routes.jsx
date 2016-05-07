import React from 'react';
import Route from 'react-router/lib/Route';

import App from './routes/App';
import Feed from './components/Feed.jsx';
import Profile from './components/Profile.jsx';
import SignIn from './components/login/SignIn.jsx';
import SignUp from './components/login/SignUp.jsx';

export default (
	<Route component={App}>
		<Route path='/' component={Feed} />
		<Route path='/profile' component={Profile} />
		<Route path='/signin' component={SignIn} />
		<Route path='/signup' component={SignUp} />
	</Route>
);
