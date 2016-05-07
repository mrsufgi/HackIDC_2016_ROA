var React = require('react');
import {Button} from 'react-bootstrap';

var SeeMore = React.createClass({
	render() {
		return (
			<Button className='see-more' bsStyle='info'>see more</Button>
		);
	}
});

module.exports = SeeMore;
