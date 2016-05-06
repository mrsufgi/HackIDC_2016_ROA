import React, { Component } from 'react';
var Header = require('../components/Header.jsx');
var Feed = require('../components/Feed.jsx');

class Home extends Component {
	render() {
		return (
			<div>
				<Header/>
				<Feed />
			</div>
		);
	}
}

export default Home;
