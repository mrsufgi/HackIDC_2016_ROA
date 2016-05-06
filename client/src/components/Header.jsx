var React = require('react');
var Header = React.createClass({
	style: {
		list: {
			color: 'blue'
		}
	},
	propTypes: {
		children: React.PropTypes.element,
		style: React.PropTypes.string
	},
	render() {
		return (
			<div>
				<ul style={this.style}>
					<li>
						1
					</li>
					<li>
						2
					</li>
				</ul>
			</div>
		);
	}
});

module.exports = Header;
