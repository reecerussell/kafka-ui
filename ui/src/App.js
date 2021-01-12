import React from "react";
import Messages from "./components/messages";
import NavBar from "./components/navbar";
import "./scss/styles.scss"

import {getSelectedTopic} from "./redux/reducers/topicReducer" 
import { connect } from "react-redux";

const App = ({selectedTopic}) => {
	return (
		<>
			<NavBar />

			<main className="container">
			{selectedTopic && <Messages topic={selectedTopic} />}
			</main>
		</>
	);
}

const mapStateToProps = (state) => ({
	selectedTopic: getSelectedTopic(state)
})

const mapDispatchToProps = (state) => ({})

export default connect(mapStateToProps, mapDispatchToProps)(App);
