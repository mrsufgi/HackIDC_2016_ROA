import { Navbar, Nav, NavDropdown, NavItem, MenuItem } from 'react-bootstrap';
var React = require('react');
var NavBar = React.createClass({
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
						<a href='/'>RoastMe!</a>
					</Navbar.Brand>
					<Navbar.Toggle />
				</Navbar.Header>
				<Navbar.Collapse>
					<Nav>
						<NavItem href='#'>New</NavItem>
						<NavItem href='#'>Trending</NavItem>
						<NavItem href='#'>Favorites</NavItem>
					</Nav>
					<Nav pullRight>
						<NavDropdown title='Profile' id='basic-nav-dropdown'>
							<MenuItem>My Profile</MenuItem>
							<MenuItem>Edit Profile</MenuItem>
							<MenuItem></MenuItem>
							<MenuItem divider />
							<MenuItem>Sign Out</MenuItem>
						</NavDropdown>
					</Nav>
				</Navbar.Collapse>
			</Navbar>
		);
	}
});

module.exports = NavBar;
