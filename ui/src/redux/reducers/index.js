import { combineReducers } from "redux";
import topicReducer from "./topicReducer";
import messagesReducer from "./messagesReducer";

const appReducer = combineReducers({
	topics: topicReducer,
	messages: messagesReducer,
});

const rootReducer = (state, action) => appReducer(state, action);

export default rootReducer;
