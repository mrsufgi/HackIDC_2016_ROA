var React = require('react');
import {Button} from 'react-bootstrap';
var FontAwesome = require('react-fontawesome');

var RoastCardActions = React.createClass({
	propTypes: {
		id: React.PropTypes.string
	},
	like: function() {
		if (this.state.votedD) {
			this.setState({
				votedD: false,
				votedL: true,
				dislike: this.state.dislike - 1,
				like: this.state.like + 1
			});
		} else if (!this.state.votedD && !this.state.votedL) {
			this.setState({
				votedD: false,
				votedL: true,
				like: this.state.like + 1
			});
		}
		console.log('like');
	},
	dislike: function() {
		if (this.state.votedL) {
			this.setState({
				votedD: true,
				votedL: false,
				dislike: this.state.dislike + 1,
				like: this.state.like - 1
			});
		} else if (!this.state.votedD && !this.state.votedL) {
			this.setState({
				votedD: true,
				votedL: false,
				dislike: this.state.like + 1
			});
		}
		console.log('dislike');
	},
	getInitialState: function() {
		return ({dislike: 0, like: 0, votedL: false, votedD: false});
	},
	render() {
		return (
			<table style={{width: '100%'}}>
				<tr>
					<td style={{width: '50%'}}>
						<Button style={{width: '90%', margin: 'auto'}} bsStyle='success' onClick={this.like}>
							<div>
								{this.state.like} <FontAwesome name='thumbs-o-up' />
							</div>
						</Button>
					</td>
					<td style={{width: '50%'}}>
						<Button style={{width: '100%', margin: 'auto'}} bsStyle='danger' onClick={this.dislike}>
							<div>
								{this.state.dislike} <FontAwesome name='thumbs-o-down'/>
							</div>
						</Button>
					</td>
				</tr>
			</table>
		);
	}
});

module.exports = RoastCardActions;
