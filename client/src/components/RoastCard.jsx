var React = require('react');
var RoastCardPic = require('../components/RoastCardPic.jsx');
var CommentBox = require('../components/CommentBox.jsx');
var CustomPanelHeader = require('../components/CustomPanelHeader.jsx');
import {Panel} from 'react-bootstrap';

var RoastCard = React.createClass({
	propTypes: {
		cardData: React.PropTypes.any
	},
	render() {
		return (
			<div className='roast-card'>
				<Panel header={<CustomPanelHeader title={this.props.cardData.title} user={this.props.cardData.name}/>} bsStyle='primary'>
					<RoastCardPic pic={this.props.cardData.roastPicSrc} />
					<CommentBox comments={this.props.cardData.comments}></CommentBox>
				</Panel>

			</div>
		);
	}
});
module.exports = RoastCard;
