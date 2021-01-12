import * as types from "./messageActionTypes";

export const addMessage = (topic, payload) => ({
	type: types.ADD_MESSAGE,
	topic: topic,
	payload: payload,
});

export const setSelectedMessage = (message) => ({
	type: types.SET_SELECTED_MESSAGE,
	message: message,
});
