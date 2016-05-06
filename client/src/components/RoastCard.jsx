var React = require('react');
var CardHeader = require('../components/CardHeader.jsx');
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
					<CardHeader name={this.props.cardData.name}
								pic={this.props.cardData.roastPicSrc}>
					</CardHeader>
					<CommentBox comments={this.props.cardData.comments}></CommentBox>
				</Panel>

			</div>
		);
	}
});
module.exports = RoastCard;
