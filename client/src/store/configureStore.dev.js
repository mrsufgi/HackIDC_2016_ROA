import {
	createStore,
	applyMiddleware
} from 'redux';
import thunk from 'redux-thunk';
import createLogger from 'redux-logger';

import rootReducer from '../reducers';
import authenticationReducer from '../reducers/authentication';

export default function configureStore(initialState) {
	const store = createStore(
		rootReducer,
		initialState,
		authenticationReducer,
		applyMiddleware(thunk, createLogger())
	);
	if (module.hot) {
		// Enable Webpack hot module replacement for reducers
		module.hot.accept('../reducers', () => {
			const nextRootReducer = require('../reducers').default;
			store.replaceReducer(nextRootReducer);
		});
	}

	return store;
};
