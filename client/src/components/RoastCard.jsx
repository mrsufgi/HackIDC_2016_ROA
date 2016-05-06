var React = require('react');
var CardHeader = require('../components/CardHeader.jsx');
var CommentBox = require('../components/CommentBox.jsx');

var RoastCard = React.createClass({
	propTypes: {
		cardData: React.PropTypes.any
	},
	render() {
		return (
			<div className='roast-card'>
				<CardHeader name={this.props.cardData.name}
							pic={this.props.cardData.roastPicSrc}
							title={this.props.cardData.title}>

				</CardHeader>
				<CommentBox comments={this.props.cardData.comments}></CommentBox>
			</div>
		);
	}
});
module.exports = RoastCard;
