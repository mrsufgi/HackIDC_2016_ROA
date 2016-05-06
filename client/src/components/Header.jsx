import { Navbar, Nav, NavDropdown, NavItem, MenuItem } from 'react-bootstrap';
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
			<Navbar inverse>
				<Navbar.Header>
					<Navbar.Brand>
						<a href='#'>RoastMe!</a>
					</Navbar.Brand>
					<Navbar.Toggle />
				</Navbar.Header>
				<Navbar.Collapse>
					<Nav>
						<NavItem eventKey={1} href='#'>Feed</NavItem>
						<NavItem eventKey={2} href='#'>Link</NavItem>
					</Nav>
					<Nav pullRight>
						<NavDropdown eventKey={3} title='Profile' id='basic-nav-dropdown'>
							<MenuItem eventKey={3.1}>Me</MenuItem>
							<MenuItem eventKey={3.2}>Edit Profile</MenuItem>
							<MenuItem eventKey={3.3}></MenuItem>
							<MenuItem divider />
							<MenuItem eventKey={3.3}>Sign Out</MenuItem>
						</NavDropdown>
					</Nav>
				</Navbar.Collapse>
			</Navbar>
		);
	}
});

module.exports = Header;
