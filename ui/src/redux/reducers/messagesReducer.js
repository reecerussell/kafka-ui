import * as types from "../actions/messageActionTypes";
import initialState from "../initialState";

const reducer = (state = initialState.messages, action) => {
	switch (action.type) {
		case types.ADD_MESSAGE:
			return {
				...state,
				messages: state.messages.concat({
					...action.payload,
					topic: action.topic,
				}),
			};
		case types.SET_SELECTED_MESSAGE:
			return {
				...state,
				selectedMessages: state.selectedMessages
					.filter((x) => x.topic !== action.message.topic)
					.concat(action.message),
			};
		default:
			return state;
	}
};

export const getMessages = (state) => state.messages.messages;
export const getSelectedMessages = (state) => state.messages.selectedMessages;

export default reducer;
