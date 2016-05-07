import React from 'react';
import NavBar from '../components/NavBar.jsx';

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

export default App;
