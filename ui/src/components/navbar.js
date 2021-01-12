import React, { useEffect } from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import classNames from "classnames";
import { fetchTopics } from "../api/topics";
import {setSelectedTopic as setSelectedTopicAction} from "../redux/actions/topicActions"
import {
	getTopics, getSelectedTopic
} from "../redux/reducers/topicReducer";
import {getMessages} from "../redux/reducers/messagesReducer"

const setSelectedTopic = topicName => dispatch => dispatch(setSelectedTopicAction(topicName))

const NavBar = ({ fetchTopics, setSelectedTopic, topics, messages, selectedTopic }) => {
	useEffect(() => {
		fetchTopics();
	}, [fetchTopics]);

	return (
		<>
            <nav className="navbar navbar-expand-lg fixed-top navbar-dark bg-dark">
				<div className="container-fluid pl-lg-0">
					<a className="navbar-brand" href="#">Kafka UI</a>
					<button className="navbar-toggler p-0 border-0" type="button">
						<span className="navbar-toggler-icon"></span>
					</button>
					<div className="collapse navbar-collapse">
						<ul className="navbar-nav me-auto mb-2 mb-lg-0">
							<li className="nav-item">
							<a className="nav-link active" aria-current="page" href="#">Home</a>
							</li>
						</ul>
					</div>
				</div>
			</nav>

			<div className="nav-scroller bg-white shadow-sm" id="topic-nav">
				<nav className="nav nav-underline">
                    {topics.map((topic, key) => {
                        const messageCount = messages.filter(x => x.topic === topic.name).length

                        const handleClick = e => {
                            e.preventDefault()

                            setSelectedTopic(topic.name)
                        }

                        const classes = classNames("nav-link", {"active": selectedTopic === topic.name})

                        return (
                            <a className={classes} href="#" onClick={handleClick} key={key}>{topic.displayName || topic.name}
                                {messageCount > 0 && (
                                    <span className="badge bg-light text-dark rounded-pill align-text-bottom">{messageCount}</span>
                                )}
                            </a>
                        )
                    })}
				</nav>
			</div>

        </>
	);
};

NavBar.propTypes = {
    topics: PropTypes.array.isRequired,
    messages: PropTypes.array.isRequired,
    selectedTopic: PropTypes.string,
    fetchTopics: PropTypes.func.isRequired,
    setSelectedTopic: PropTypes.func.isRequired
};

NavBar.defaultProps = {
    topics: [],
    messages: [],
    selectedTopic: null
};

const mapStateToProps = (state) => ({
    topics: getTopics(state),
    messages: getMessages(state),
    selectedTopic: getSelectedTopic(state)
});

const mapDispatchToProps = (dispatch) =>
	bindActionCreators(
		{
            fetchTopics: fetchTopics,
            setSelectedTopic: setSelectedTopic
		},
		dispatch
	);

export default connect(mapStateToProps, mapDispatchToProps)(NavBar);
