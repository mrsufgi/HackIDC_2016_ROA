import { combineReducers } from 'redux';
import { routerReducer as routing } from 'react-router-redux';
import { authenticationReducer } from './authentication';
import app from './App';

const rootReducer = combineReducers({
	app,
	routing,
	authenticationReducer
});

export default rootReducer;
