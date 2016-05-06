var React = require('react');
var Header = React.createClass({
	propTypes: {
		children: React.PropTypes.element
	},
	render() {
		return (
			<div>
				{this.props.children}
			</div>
		);
	}
});

module.exports = Header;
