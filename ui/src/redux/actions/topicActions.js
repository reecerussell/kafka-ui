import * as types from "./topicActionTypes";

export const getTopics = () => ({
	type: types.GET_TOPICS,
});

export const getTopicsSuccess = (topics) => ({
	type: types.GET_TOPICS_SUCCESS,
	topics: topics,
});

export const getTopicsError = (error) => ({
	type: types.GET_TOPICS_ERROR,
	error: error,
});

export const setSelectedTopic = (topicName) => ({
	type: types.SET_SELECTED_TOPIC,
	topic: topicName
})