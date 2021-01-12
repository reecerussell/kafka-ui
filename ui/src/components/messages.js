import React from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import { getMessages } from "../redux/reducers/messagesReducer";
import { setSelectedMessage as setSelectedMessageAction } from "../redux/actions/messageActions";
import Message from "./message";

const setSelectedMessage = (message) => (dispatch) =>
	dispatch(setSelectedMessageAction(message));

const Messages = ({ messages, topic, setSelectedMessage }) => {
	const renderHeader = (message) => {
		if (message.key) {
			return message.key;
		}

		return (
			<>
				Message
				<span>no key</span>
			</>
		);
	};

	return (
		<div className="row kui-message-container">
			<div className="col-md-4">
				<ul className="list-group list-group-flush list-group-hover kui-messages">
					{messages
						.filter((m) => m.topic === topic)
						.map((message, key) => (
							<li className="list-group-item" key={key}>
								<h4 onClick={() => setSelectedMessage(message)}>
									{renderHeader(message)}
								</h4>
								<small className="text-muted">
									{new Date(
										message.timestamp * 1000
									).toISOString()}
								</small>
							</li>
						))}
				</ul>
			</div>
			<div className="col-md-8">
				<Message topic={topic} />
			</div>
		</div>
	);
};

Messages.propTypes = {
	messages: PropTypes.array.isRequired,
	topic: PropTypes.string.isRequired,
	setSelectedMessage: PropTypes.func.isRequired,
};

Messages.defaultProps = {
	messages: [],
};

const mapStateToProps = (state) => ({
	messages: getMessages(state),
});

const mapDispatchToProps = (dispatch) =>
	bindActionCreators(
		{
			setSelectedMessage: setSelectedMessage,
		},
		dispatch
	);

export default connect(mapStateToProps, mapDispatchToProps)(Messages);
