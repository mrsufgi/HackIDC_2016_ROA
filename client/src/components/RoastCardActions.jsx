var React = require('react');
import {Button} from 'react-bootstrap';
var FontAwesome = require('react-fontawesome');

var RoastCardActions = React.createClass({
	propTypes: {
		id: React.PropTypes.string
	},
	render() {
		return (
			<table style={{width: '100%'}}>
				<tr>
					<td style={{width: '50%'}}>
						<Button style={{width: '90%', margin: 'auto'}} bsStyle='success'>
							<FontAwesome name='thumbs-o-up' />
						</Button>
					</td>
					<td style={{width: '50%'}}>
						<Button style={{width: '90%', margin: 'auto'}} bsStyle='danger'>
							<FontAwesome name='thumbs-o-down'/>
						</Button>
					</td>
				</tr>
			</table>
		);
	}
});

module.exports = RoastCardActions;
