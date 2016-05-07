import React from 'react';
import NavBar from '../components/NavBar.jsx';
import { connect } from 'react-redux';

require('../../css/style.css');

var App = React.createClass({
	propTypes: {
		children: React.PropTypes.any,
		user: React.PropTypes.object
	},
	render() {
		return (
			<div>
				<NavBar user={this.props.user}/>
				{this.props.children}
			</div>
		);
	}
});

const mapStateToProps = function(state) {
	console.log(state);
	return Object.assign({}, state.app, state.routing);
};
export default connect(mapStateToProps)(App);
