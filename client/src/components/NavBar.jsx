import { Navbar, Nav, NavDropdown, NavItem, MenuItem } from 'react-bootstrap';
import React from 'react';
import { Link } from 'react-router';

var NavBar = React.createClass({
	propTypes: {
		children: React.PropTypes.element,
		style: React.PropTypes.string
	},
	render() {
		return (
			<Navbar inverse>
				<Navbar.Header>
					<Navbar.Brand>
						<Link to='/'>
							RoastMe!
						</Link>
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
							<MenuItem>
								<Link to='/profile'>
									My Profile
								</Link>
							</MenuItem>
							<MenuItem>
								<Link to='/editProfile'>
									Edit Profile
								</Link>
							</MenuItem>
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
