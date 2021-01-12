import React, { useEffect } from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import { fetchTopics } from "../api/topics";
import {
	getTopics,
	getTopicsLoading,
	getTopicsError,
} from "../redux/reducers/topicReducer";
import Messages from "./messages";

const Topics = ({ error, loading, fetchTopics, topics }) => {
	useEffect(() => {
		fetchTopics();
	}, [fetchTopics]);

	if (loading) {
		<p>Loading...</p>;
	}

	if (error) {
		return <p>Error: {error}</p>;
	}

	return (
		<div className="row">
			<div class="col-lg-4">
				<ul className="list-group list-group-flush">
					{topics.map((topic, key) => (
						<li className="list-group-item" key={key}>
							{topic.displayName || topic.name}
							<br />
							<Messages topic={topic.name} />
						</li>
					))}
				</ul>
			</div>
		</div>
	);
};

Topics.propTypes = {
	error: PropTypes.string,
	loading: PropTypes.bool.isRequired,
	topics: PropTypes.array.isRequired,
	fetchTopics: PropTypes.func.isRequired,
};

Topics.defaultProps = {
	topics: [],
};

const mapStateToProps = (state) => ({
	error: getTopicsError(state),
	topics: getTopics(state),
	loading: getTopicsLoading(state),
});

const mapDispatchToProps = (dispatch) =>
	bindActionCreators(
		{
			fetchTopics: fetchTopics,
		},
		dispatch
	);

export default connect(mapStateToProps, mapDispatchToProps)(Topics);
