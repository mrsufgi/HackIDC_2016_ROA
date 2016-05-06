var React = require('react');

var CardHeader = React.createClass({
	propTypes: {
		name: React.PropTypes.string,
		pic: React.PropTypes.string,
		title: React.PropTypes.string
	},
	render() {
		return (
			<div className='roast-card-header'>
				<img src={this.props.pic} alt='pic'/>
				<h3>{this.props.title}</h3>
				<p>by {this.props.name}</p>
			</div>
		);
	}
});
module.exports = CardHeader;
