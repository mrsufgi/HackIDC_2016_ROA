var React = require('react');

var CardContent = React.createClass({
	propTypes: {
		name: React.PropTypes.string,
		pic: React.PropTypes.string,
		title: React.PropTypes.string
	},
	render() {
		return (
			<div className='roast-card-header'>
				<img className='roast-card-img' src={this.props.pic} alt='pic'/>
			</div>
		);
	}
});
module.exports = CardContent;
