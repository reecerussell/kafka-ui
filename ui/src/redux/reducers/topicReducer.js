import * as types from "../actions/topicActionTypes";
import initialState from "../initialState";

const topicReducer = (state = initialState.topics, action) => {
	switch (action.type) {
		case types.GET_TOPICS:
			return {
				...state,
				loading: true,
			};
		case types.GET_TOPICS_SUCCESS:
			return {
				...state,
				loading: false,
				topics: action.topics,
				selectedTopic:
					state.selectedTopic ||
					(action.topics.length > 0 ? action.topics[0].name : null),
			};
		case types.GET_TOPICS_ERROR:
			return {
				...state,
				loading: false,
				error: action.error,
			};

		case types.SET_SELECTED_TOPIC:
			return {
				...state,
				selectedTopic: action.topic,
			};
		default:
			return state;
	}
};

export const getTopics = (state) => state.topics.topics;
export const getTopicsLoading = (state) => state.topics.loading;
export const getTopicsError = (state) => state.topics.error;
export const getSelectedTopic = (state) => state.topics.selectedTopic;

export default topicReducer;
