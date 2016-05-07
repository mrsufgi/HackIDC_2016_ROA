var React = require('react');
import {Badge} from 'react-bootstrap';

var Comment = React.createClass({
	propTypes: {
		comment: React.PropTypes.object
	},
	render() {
		return (
				<table>
					<tr>
						<td className='comments-table-cell' style={{width: '100%', textAlign: 'left'}}>
							<p style={{lineHeight: '2'}}>
								{this.props.comment.content}
							</p>
						</td>
						<td className='comments-table-cell'>
							<Badge>
								{this.props.comment.like}
							</Badge>
						</td>
					</tr>
				</table>
		);
	}
});

module.exports = Comment;
