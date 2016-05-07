var React = require('react');
import {FormGroup, ControlLabel, FormControl, Jumbotron, Grid, Row, Col, Button} from 'react-bootstrap';

var SignIn = React.createClass({
	render() {
		return (
			<Grid>
				<Row className='show-grid'>
					<Col xs={12} md={12}>
						<div>
							<Jumbotron>
								<h1 style={{textAlign: 'center'}}>Brave, aren't you?</h1>
								<blockquote>
									<p>
										Do not forget that the top community assholes will enter a raffle for the <b>reek trip</b>!
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
module.exports = SignIn;

