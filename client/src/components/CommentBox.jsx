var React = require('react');

var CommentBox = React.createClass({
	propTypes: {
		comments: React.PropTypes.array
	},
	render() {
		var postProcessComments = this.props.comments.map((comment) => {
			return (
				<li>
					<p>{comment.content}</p>
					<p>{comment.user}</p>
				</li>
			);
		});
		return (
			<ul>
				{postProcessComments}
			</ul>
		);
	}
});
module.exports = CommentBox;
