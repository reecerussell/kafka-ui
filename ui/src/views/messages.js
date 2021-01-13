import React from "react";
import { getSelectedTopic } from "../redux/reducers/topicReducer";
import { connect } from "react-redux";
import Messages from "../components/messages";

const MessagesView = ({ selectedTopic }) =>
	selectedTopic && <Messages topic={selectedTopic} />;

const mapStateToProps = (state) => ({
	selectedTopic: getSelectedTopic(state),
});

const mapDispatchToProps = (dispatch) => ({});

export default connect(mapStateToProps, mapDispatchToProps)(MessagesView);
