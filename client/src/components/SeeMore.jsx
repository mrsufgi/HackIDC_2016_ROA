var React = require('react');
import {Button} from 'react-bootstrap';

var SeeMore = React.createClass({
	render() {
		return (
			<Button style={{width: '100%', backgroundColor: 'rgba(0,0,0,0.05)'}}>Read More</Button>
		);
	}
});

module.exports = SeeMore;
