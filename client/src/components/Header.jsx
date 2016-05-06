// import { Button } from 'react-bootstrap';
var React = require('react');
var Header = React.createClass({
	style: {
		list: {
			color: 'blue'
		}
	},
	propTypes: {
		children: React.PropTypes.element,
		style: React.PropTypes.string
	},
	render() {
		return (
			<nav className='navbar navbar-default'>
				<div className='container-fluid'>
					<div className='navbar-header'>
						<a className='navbar-brand' href='#'>
							RoastMe!
						</a>
					</div>
				</div>
			</nav>
		);
	}
});

module.exports = Header;
