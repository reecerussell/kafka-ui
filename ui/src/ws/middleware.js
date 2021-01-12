import * as messageActions from "../redux/actions/messageActions";
import { Listener } from "./listener";

const middleware = (store) => {
	const ws = new Listener();
	ws.Start();

	ws.Handle("Message", (payload) => {
		console.log("Handler", payload);
		store.dispatch(
			messageActions.addMessage(payload.topic, payload.message)
		);
	});

	return (next) => (action) => next(action);
};

export default middleware;
