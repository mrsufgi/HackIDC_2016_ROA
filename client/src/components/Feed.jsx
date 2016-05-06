var React = require('react');
var RoastCard = require('../components/RoastCard.jsx');
var Feed = React.createClass({
	propTypes: {
		children: React.PropTypes.element,
		dummy: React.PropTypes.object
	},
	dummyRoastCard: [
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

		},
		{
			title: 'What do you think about me?2',
			name: 'nadavg',
			roastPicSrc: 'https://fbcdn-profile-a.akamaihd.net/hprofile-ak-xaf1/v/t1.0-1/c0.0.160.160/p160x160/1535443_10203361054681037_1835679698273652054_n.jpg?oh=a422a89c222fe4f4f6a6d017c31c46ba&oe=57997593&__gda__=1470518677_71e3f57d5434c6abd3ec9d55afd6736e',
			comments: [
				{
					content: '',
					user: '',
					score: {
						up: 5,
						down: 9
					}
				}
			]

		}
	],
	render() {
		var data = this.dummyRoastCard.map((card) => {
			return <RoastCard cardData={card}/>;
		});
		return (
			<div>
				{data}
			</div>
		);
	}
});

module.exports = Feed;
