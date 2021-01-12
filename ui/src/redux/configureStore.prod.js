import { createStore, applyMiddleware } from "redux";
import reduxImmutableStateInvariant from "redux-immutable-state-invariant";
import thunk from "redux-thunk";
import rootReducer from "./reducers";
import wsMiddleware from "../ws/middleware";

const configureStore = (initialState) => {
	return createStore(
		rootReducer,
		initialState,
		applyMiddleware(thunk, reduxImmutableStateInvariant(), wsMiddleware)
	);
};

export default configureStore;
