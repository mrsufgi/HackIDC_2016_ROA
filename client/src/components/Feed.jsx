var React = require('react');
var RoastCard = require('../components/RoastCard.jsx');
import {Grid, Row, Col} from 'react-bootstrap';

var Feed = React.createClass({
	propTypes: {
		children: React.PropTypes.element,
		dummy: React.PropTypes.object
	},
	dummyRoastCard: [1, 2, 3, 4, 5].map((x) => {
		return (
		{
			title: 'Roast me mofos!',
			name: 'Aviadhahami',
			roastPicSrc: 'https://fbcdn-profile-a.akamaihd.net/hprofile-ak-xaf1/v/t1.0-1/p160x160/11898648_10207488666114110_1556072271457749853_n.jpg?oh=cf32c5a9d4cf727242c7de8bbf1e7e03&oe=57B4AED5&__gda__=1474549326_40ee7ad9336e48a8d66cece017b1e6c0',
			comments: [
				{
					content: 'Haha you\'re fat!',
					user: '',
					score: {
						up: 5,
						down: 9
					}
				}
			]
		}
		);
	}),
	render() {
		var data = this.dummyRoastCard.map((card, i) => {
			return (
				<Col xs={12} md={4}>
					<RoastCard cardData={card} key={i}/>
				</Col>
			);
		});
		return (
			<Grid>
				<Row className='show-grid'>
					{data}
				</Row>
			</Grid>
		);
	}
});

module.exports = Feed;
