var React = require('react');
import {FormGroup, ControlLabel, FormControl, Jumbotron, Grid, Row, Col, Button} from 'react-bootstrap';

var SignOut = React.createClass({
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
							<form action='' style={{width: '80%', margin: 'auto'}}>
								<FormGroup controlId='formControlsText'>
									<ControlLabel>Nickname</ControlLabel>
									<FormControl type='text' placeholder='Nickname' />
								</FormGroup>
								<FormGroup controlId='formControlsPassword'>
									<ControlLabel>Password</ControlLabel>
									<FormControl type='password' />
								</FormGroup>
								<Button type='submit' bsStyle='success' style={{float: 'right', width: '30%'}}>Submit</Button>
							</form>
						</div>
					</Col>
				</Row>
			</Grid>
		);
	}
});
module.exports = SignOut;
