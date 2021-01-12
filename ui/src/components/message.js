import React from "react";
import { connect } from "react-redux";
import PropTypes from "prop-types";
import { getSelectedMessages } from "../redux/reducers/messagesReducer";
import Json from "./json";

const Message = ({ selectedMessages, topic }) => {
	const message = selectedMessages.find((x) => x.topic === topic);
	if (!message) {
		return null;
	}

	let isJson = false;
	let messageValue = message.value;

	try {
		messageValue = JSON.parse(messageValue);
		isJson = true;
	} catch {
		isJson = false;
	}

	return (
		<div className="card kui-message">
			<div className="card-body">
				<p>
					<b>Key:</b>{" "}
					{message.key || (
						<span className="text-muted">(no key)</span>
					)}
					<br />
					<b>Timestamp:</b>{" "}
					{new Date(message.timestamp * 1000).toISOString()}
					<br />
					{isJson ? <b>Value:</b> : <b>Value: {messageValue}</b>}
				</p>
				{isJson && <Json data={messageValue} />}
			</div>
		</div>
	);
};

Message.propTypes = {
	selectedMessages: PropTypes.array.isRequired,
	topic: PropTypes.string.isRequired,
};

Message.defaultProps = {
	selectedMessages: [],
};

const mapStateToProps = (state) => ({
	selectedMessages: getSelectedMessages(state),
});

const mapDispatchToProps = () => ({});

export default connect(mapStateToProps, mapDispatchToProps)(Message);
