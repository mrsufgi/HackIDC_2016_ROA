var React = require('react');
import {Badge} from 'react-bootstrap';

var Comment = React.createClass({
	propTypes: {
		comment: React.PropTypes.object
	},
	getInitialState: function() {
		return ({
			voted: false,
			like: this.props.comment.like
		});
	},
	vote: function() {
		if (!this.state.voted) {
			this.setState({
				like: this.props.comment.like + 1,
				voted: true
			});
		}
	},
	render() {
		return (
			<table>
				<tr>
					<td style={{width: '100%', textAlign: 'left'}}>
						<p style={{lineHeight: '2'}}>
							{this.props.comment.content}
						</p>
					</td>
					<td>
						<Badge onClick={this.vote} className={'cursor-pointer roast-likes ' + (this.state.voted ? 'roast-likes-done' : '')}>
							{this.state.like}
						</Badge>
					</td>
				</tr>
			</table>
		);
	}
});

module.exports = Comment;
