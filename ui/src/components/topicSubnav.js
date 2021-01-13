import React, { useEffect } from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import classNames from "classnames";
import { fetchTopics } from "../api/topics";
import { setSelectedTopic as setSelectedTopicAction } from "../redux/actions/topicActions";
import { getTopics, getSelectedTopic } from "../redux/reducers/topicReducer";
import { getMessages } from "../redux/reducers/messagesReducer";

const setSelectedTopic = (topicName) => (dispatch) =>
	dispatch(setSelectedTopicAction(topicName));

const TopicSubnav = ({
	fetchTopics,
	setSelectedTopic,
	topics,
	messages,
	selectedTopic,
}) => {
	useEffect(() => {
		fetchTopics();
	}, [fetchTopics]);

	return (
		<div className="nav-scroller bg-white shadow-sm" id="topic-nav">
			<nav className="nav nav-underline">
				{topics.map((topic, key) => {
					const messageCount = messages.filter(
						(x) => x.topic === topic.name
					).length;

					const handleClick = (e) => {
						e.preventDefault();

						setSelectedTopic(topic.name);
					};

					const classes = classNames("nav-link", {
						active: selectedTopic === topic.name,
					});

					return (
						<a
							className={classes}
							href="#"
							onClick={handleClick}
							key={key}
						>
							{topic.displayName || topic.name}
							{messageCount > 0 && (
								<span className="badge bg-light text-dark rounded-pill align-text-bottom">
									{messageCount}
								</span>
							)}
						</a>
					);
				})}
			</nav>
		</div>
	);
};

TopicSubnav.propTypes = {
	topics: PropTypes.array.isRequired,
	messages: PropTypes.array.isRequired,
	selectedTopic: PropTypes.string,
	fetchTopics: PropTypes.func.isRequired,
	setSelectedTopic: PropTypes.func.isRequired,
};

TopicSubnav.defaultProps = {
	topics: [],
	messages: [],
	selectedTopic: null,
};

const mapStateToProps = (state) => ({
	topics: getTopics(state),
	messages: getMessages(state),
	selectedTopic: getSelectedTopic(state),
});

const mapDispatchToProps = (dispatch) =>
	bindActionCreators(
		{
			fetchTopics: fetchTopics,
			setSelectedTopic: setSelectedTopic,
		},
		dispatch
	);

export default connect(mapStateToProps, mapDispatchToProps)(TopicSubnav);
