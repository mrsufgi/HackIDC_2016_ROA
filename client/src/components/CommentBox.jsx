var React = require('react');
var Comment = require('../components/Comment.jsx');

var CommentBox = React.createClass({
	propTypes: {
		comments: React.PropTypes.object
	},
	render() {
		var postProcessComments = this.props.comments.map((comment) => {
			return (
				<Comment className='line-after' comment={comment}></Comment>
			);
		});
		return (
			<div>
				{postProcessComments}
			</div>
		);
	}
});
module.exports = CommentBox;
