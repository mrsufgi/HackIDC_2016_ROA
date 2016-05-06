import React from 'react';
import NavBar from '../components/NavBar.jsx';
var App = React.createClass({
	propTypes: {
		children: React.PropTypes.element
	},
	render() {
		return (
			<div>
				<NavBar />
				{this.props.children}
			</div>
		);
	}
});

export default App;
