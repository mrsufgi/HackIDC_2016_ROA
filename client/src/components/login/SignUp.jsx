var React = require('react');
import {FormGroup, ControlLabel, FormControl, Jumbotron, Grid, Row, Col} from 'react-bootstrap';

var SignOut = React.createClass({
	handleSubmit(event) {
		var data = {
			username: event.target.username.value,
			password: event.target.password.value
		};
		event.preventDefault();
		console.log(data);
	},
	render() {
		return (
			<Grid>
				<Row className='show-grid'>
					<Col xs={12} md={12}>
						<div>
							<Jumbotron>
								<h1 style={{textAlign: 'center'}}>Disclaimer!</h1>
								<blockquote>
									<p>
										<b>This app is not for the weak!</b>
										<br/>
										By signing up you will get <i>ROASTED</i> immediately by our community of pure assholes.
										<br/>
										After signing up and uploading your first photo, you will be able to <i>ROAST</i> others and
										have the opportunity to be the true asshole that you are!
										<br/>
										<br/>
										Isn't it lovely?
									</p>
									<small> ♥ Ramsay Bolton</small>
								</blockquote>
							</Jumbotron>
						</div>
						<div>
							<form action='' onSubmit={this.handleSubmit} style={{width: '80%', margin: 'auto'}}>
								<FormGroup controlId='username'>
									<ControlLabel>Nickname</ControlLabel>
									<FormControl type='text' placeholder='Nickname' />
								</FormGroup>
								<FormGroup controlId='password'>
									<ControlLabel>Password</ControlLabel>
									<FormControl type='password' />
								</FormGroup>
								<FormControl type='submit' bsStyle='success' style={{float: 'right', width: '30%'}} value='Submit' />
							</form>
						</div>
					</Col>
				</Row>
			</Grid>
		);
	}
});
module.exports = SignOut;

