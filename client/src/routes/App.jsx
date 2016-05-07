import React from 'react';
import NavBar from '../components/NavBar.jsx';
import { connect } from 'react-redux';
require('../../css/style.css');

var App = React.createClass({
	propTypes: {
		children: React.PropTypes.element,
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

function mapStateToProps(state) {
	return state;
}
export default connect(mapStateToProps)(App);
