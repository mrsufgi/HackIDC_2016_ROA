var React = require('react');
import {Link} from 'react-router';

var CustomPanelHeader = React.createClass({
	propTypes: {
		title: React.PropTypes.string,
		user: React.PropTypes.string
	},
	render() {
		var link = '/user/' + this.props.user;
		return (
			<div className='show-grid'>
				{this.props.title} (
				<Link to={link} className='no-style'>
					{this.props.user}
				</Link>
				)
			</div>
		);
	}
});

module.exports = CustomPanelHeader;
