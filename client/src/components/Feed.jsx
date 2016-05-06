var React = require('react');
var Feed = React.createClass({
	propTypes: {
		children: React.PropTypes.element
	},
	render() {
		return (
			<div>
				<h1>This is feed</h1>
			</div>
		);
	}
});

module.exports = Feed;
