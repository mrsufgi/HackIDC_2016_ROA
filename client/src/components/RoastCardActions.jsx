var React = require('react');
import {Button} from 'react-bootstrap';

var RoastCardActions = React.createClass({
	render() {
		return (
			<table style={{width: '100%'}}>
				<tr>
					<td style={{width: '50%;'}}>
						<Button style={{width: '100%'}}>Test</Button>
					</td>
					<td style={{width: '50%;'}}>
						<Button style={{width: '100%'}}>Test</Button>
					</td>
				</tr>
			</table>
		);
	}
});

module.exports = RoastCardActions;
