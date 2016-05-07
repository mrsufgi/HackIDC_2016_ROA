var React = require('react');
var RoastCardPic = require('../components/RoastCardPic.jsx');
var CommentBox = require('../components/CommentBox.jsx');
var CustomPanelHeader = require('../components/CustomPanelHeader.jsx');
var RoastCardActions = require('../components/RoastCardActions.jsx');
var SeeMore = require('../components/SeeMore.jsx');
import {Link} from 'react-router';
import {Panel} from 'react-bootstrap';

var RoastCard = React.createClass({
	propTypes: {
		cardData: React.PropTypes.object
	},
	render() {
		var location = '/' + this.props.cardData.id;
		return (
			<div className='roast-card shadow'>
				<Panel header={<CustomPanelHeader title={this.props.cardData.title} user={this.props.cardData.name}/>} bsStyle='primary' footer={<RoastCardActions id={this.props.cardData.id} />}>
					<div>
						<RoastCardPic pic={this.props.cardData.roastPicSrc}></RoastCardPic>
						<CommentBox comments={this.props.cardData.comments}></CommentBox>
						<Link to={location}>
							<SeeMore />
						</Link>
					</div>
				</Panel>

			</div>
		);
	}
});
module.exports = RoastCard;
