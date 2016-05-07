import { Navbar, Nav, NavDropdown, NavItem, MenuItem } from 'react-bootstrap';
import React from 'react';
import { Link } from 'react-router';
import { LinkContainer } from 'react-router-bootstrap';

var NavBar = React.createClass({
	propTypes: {
		children: React.PropTypes.element,
		user: React.PropTypes.object,
		style: React.PropTypes.string
	},
	getInitialState() {
		return ({
			user: null
		});
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
					{this.state.user ? (
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
						</Nav>) : (
						<Nav pullRight>
							<LinkContainer to='/signin'>
								<NavItem href='#'>Sign-in</NavItem>
							</LinkContainer>
							<LinkContainer to='/signup'>
								<NavItem href='#'>Sign-up</NavItem>
							</LinkContainer>
						</Nav>
					)}
				</Navbar.Collapse>
			</Navbar>
		);
	}
});

module.exports = NavBar;
