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
	getInitialState: function() {
		return (
		{
			user: this.props.user || null
		}
		);
	},
	render() {
		console.log(this.state.user, 'state');
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
							<NavItem href='#'>Sign-in</NavItem>
							<NavItem href='#'>Sign-up</NavItem>
						</Nav>
					)}
				</Navbar.Collapse>
			</Navbar>
		);
	}
});

module.exports = NavBar;
