var React = require('react');
var Comment = require('../components/Comment.jsx');

var CommentBox = React.createClass({
	propTypes: {
		comments: React.PropTypes.object,
		children: React.PropTypes.components
	},
	render() {
		var postProcessComments = this.props.comments.map((comment) => {
			return (
				<Comment comment={comment}></Comment>
			);
		});
		return (
			<div>
				<div className='comments-divider'>
					{postProcessComments}
				</div>
				<div>
					{this.props.children}
				</div>
			</div>
		);
	}
});
module.exports = CommentBox;
