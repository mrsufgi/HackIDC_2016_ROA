import { Navbar, Nav, NavDropdown, NavItem, MenuItem } from 'react-bootstrap';
import React from 'react';
import { Link } from 'react-router';
import { LinkContainer } from 'react-router-bootstrap';

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
							RoasteMe!
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
							<LinkContainer to='/profile'>
								<MenuItem>
									My Profile
								</MenuItem>
							</LinkContainer>
							<LinkContainer to='/editProfile'>
								<MenuItem>
									Edit Profile
								</MenuItem>
							</LinkContainer>
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
